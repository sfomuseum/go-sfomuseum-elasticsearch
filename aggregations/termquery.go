package aggregations

import (
	"context"
	"errors"
	"fmt"
	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/sfomuseum/go-sfomuseum-elasticsearch/queries"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

func TermQueryAggregation(es_client *es.Client, es_index string, label string, term string) ([]*Bucket, error) {

	q, err := queries.TermQueryAggregationQuery(label, term)

	if err != nil {
		return nil, err
	}

	rsp, err := es_client.Search(
		es_client.Search.WithContext(context.Background()),
		es_client.Search.WithIndex(es_index),
		es_client.Search.WithBody(q),
		es_client.Search.WithPretty(),
	)

	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("aggregations.%s.buckets", label)

	buckets_rsp := gjson.GetBytes(body, path)

	if !buckets_rsp.Exists() {
		return nil, errors.New("Unable to find buckets")
	}

	buckets := make([]*Bucket, 0)

	for _, b := range buckets_rsp.Array() {

		k_rsp := b.Get("key")
		c_rsp := b.Get("doc_count")

		b := &Bucket{
			Key:      k_rsp.String(),
			DocCount: c_rsp.Int(),
		}

		buckets = append(buckets, b)
	}

	return buckets, nil
}

func TermQueryAggregationKeys(es_client *es.Client, es_index string, label string, term string) ([]string, error) {

	buckets, err := TermQueryAggregation(es_client, es_index, label, term)

	if err != nil {
		return nil, err
	}

	keys := make([]string, 0)

	for _, b := range buckets {
		keys = append(keys, b.Key)
	}

	return keys, nil
}
