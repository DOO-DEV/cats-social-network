package search

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"log"
	"meower/schema"
)

type ElasticRepository struct {
	client *elastic.Client
}

func New(url string) (*ElasticRepository, error) {
	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(true))
	if err != nil {
		return nil, err
	}

	return &ElasticRepository{client: client}, nil
}

func (e *ElasticRepository) Close() {
	e.Close()
}

func (e *ElasticRepository) InsertMeow(ctx context.Context, meow schema.Meow) error {
	_, err := e.client.Index().
		Index("meows").
		Type("meow-service").
		Id(meow.ID).
		BodyJson(meow).
		Refresh("wait_for").
		Do(ctx)

	return err
}

func (e *ElasticRepository) SearchMeows(ctx context.Context, query string, skip, take uint64) ([]schema.Meow, error) {
	res, err := e.client.Search().
		Index("meows").
		Query(
			elastic.NewMultiMatchQuery(query, "body").
				Fuzziness("3").PrefixLength(1).CutoffFrequency(0.0001),
		).From(int(skip)).Size(int(take)).Do(ctx)
	if err != nil {
		return nil, err
	}

	meows := make([]schema.Meow, 0)
	for _, hit := range res.Hits.Hits {
		var meow schema.Meow
		if err := json.Unmarshal(*hit.Source, &meow); err != nil {
			log.Println(err)
		}

		meows = append(meows, meow)
	}

	return meows, nil
}
