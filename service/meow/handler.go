package main

import (
	"github.com/segmentio/ksuid"
	"html/template"
	"log"
	"meower/db"
	"meower/event"
	"meower/schema"
	"meower/util"
	"net/http"
	"time"
)

func createMeowHandler(w http.ResponseWriter, req *http.Request) {
	type res struct {
		ID string `json:"id"`
	}

	ctx := req.Context()
	body := template.HTMLEscapeString(req.FormValue("body"))
	if len(body) < 1 || len(body) > 140 {
		util.ResponseError(w, http.StatusBadRequest, "invalid body")
		return
	}

	createdAt := time.Now().UTC()
	id, err := ksuid.NewRandomWithTime(createdAt)
	if err != nil {
		util.ResponseError(w, http.StatusInternalServerError, "failed to create meow")
		return
	}
	meow := schema.Meow{
		ID:        id.String(),
		Body:      body,
		CreatedAt: createdAt,
	}

	if err := db.InsertMeow(ctx, meow); err != nil {
		util.ResponseError(w, http.StatusInternalServerError, "failed to create meow")
		return
	}

	if err := event.PublishMeowCreated(meow); err != nil {
		log.Println(err)
	}

	util.ResponseOk(w, res{ID: meow.ID})
}
