package queries

import (
	"io"
)

func TermQueryAggregationQuery(label string, term string) (io.Reader, error) {

	body := `{
    "size": 0,
    "aggs" : {
        "{{ .Label }}" : {
	    "terms": {
		"field" : "{{ .Term }}",
		"size": 10000
	    }
        }
    }
}`

	vars := struct {
		Label string
		Term  string
	}{
		Label: label,
		Term:  term,
	}

	return RenderQuery(body, vars)
}
