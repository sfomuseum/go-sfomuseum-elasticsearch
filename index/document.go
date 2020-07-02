package index

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	es "github.com/elastic/go-elasticsearch/v7"
	esapi "github.com/elastic/go-elasticsearch/v7/esapi"
	"io"
	"io/ioutil"
)

func IndexDocument(ctx context.Context, es_client *es.Client, es_index string, doc_id string, doc interface{}) error {

	b, err := json.Marshal(doc)

	if err != nil {
		return err
	}

	fh := bytes.NewReader(b)

	return IndexDocumentWithReader(ctx, es_client, es_index, doc_id, fh)
}

func IndexDocumentWithReader(ctx context.Context, es_client *es.Client, es_index string, doc_id string, fh io.Reader) error {

	select {
	case <-ctx.Done():
		return nil
	default:
		// pass
	}

	req := esapi.IndexRequest{
		Index:      es_index,
		DocumentID: doc_id,
		Body:       fh,
		Refresh:    "true",
	}

	rsp, err := req.Do(ctx, es_client)

	if err != nil {
		return err
	}

	defer rsp.Body.Close()

	switch rsp.StatusCode {
	case 200, 201:
	// pass
	default:
		body, _ := ioutil.ReadAll(rsp.Body)
		return errors.New(string(body))
	}

	return nil
}
