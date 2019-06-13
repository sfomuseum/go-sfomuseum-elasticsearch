package queries

import (
	"bufio"
	"bytes"
	"io"
	_ "log"
	"text/template"
)

func RenderQuery(body string, vars interface{}) (io.Reader, error) {

	t := template.New("query")
	t, err := t.Parse(body)

	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	err = t.Execute(wr, vars)

	if err != nil {
		return nil, err
	}

	wr.Flush()

	// log.Println(buf.String())
	return bytes.NewReader(buf.Bytes()), nil
}
