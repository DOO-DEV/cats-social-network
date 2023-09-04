package main

import (
	"log"
	"meower/db"
	"meower/schema"
	"meower/search"
	"meower/util"
	"net/http"
	"strconv"
)

func searchMeowsHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	query := req.FormValue("query")
	if len(query) == 0 {
		util.ResponseError(w, http.StatusBadRequest, "missing query parameter")
		return
	}

	skip, take, err := bindParameter(w, req)
	if err != nil {
		return
	}
	meows, err := search.SearchMeows(ctx, query, skip, take)
	if err != nil {
		log.Println(err)
		util.ResponseOk(w, []schema.Meow{})
		return
	}

	util.ResponseOk(w, meows)
}

func listMeowsHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	skip, take, err := bindParameter(w, req)
	if err != nil {
		return
	}

	meows, err := db.ListMeows(ctx, skip, take)
	if err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "could not fetch meows")
		return
	}

	util.ResponseOk(w, meows)
}

func bindParameter(w http.ResponseWriter, req *http.Request) (uint64, uint64, error) {
	skip := uint64(0)
	skipStr := req.FormValue("skip")
	take := uint64(100)
	takeStr := req.FormValue("take")

	skip, err := strconv.ParseUint(skipStr, 10, 64)
	if err != nil {
		util.ResponseError(w, http.StatusBadRequest, "invalid skip parameter")
		return 0, 0, err
	}
	take, err = strconv.ParseUint(takeStr, 10, 64)
	if err != nil {
		util.ResponseError(w, http.StatusBadRequest, "invalid take parameter")
		return 0, 0, err
	}

	return skip, take, nil
}
