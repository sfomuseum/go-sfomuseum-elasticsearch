package index

import (
	"bytes"
	"context"
	"encoding/json"
	es "github.com/elastic/go-elasticsearch"
	esapi "github.com/elastic/go-elasticsearch/esapi"
)

func IndexDocument(es_client *es.Client, es_index string, doc_id string, doc interface{}) error {

	b, err := json.Marshal(doc)

	if err != nil {
		return err
	}

	fh := bytes.NewReader(b)

	req := esapi.IndexRequest{
		Index:      es_index,
		DocumentID: doc_id,
		Body:       fh,
		Refresh:    "true",
	}

	_, err = req.Do(context.Background(), es_client)

	if err != nil {
		return err
	}

	return nil
}
