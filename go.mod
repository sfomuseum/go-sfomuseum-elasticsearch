module github.com/sfomuseum/go-sfomuseum-elasticsearch

// Note that elastic/go-elasticsearch/v7 v7.13.0 is the last version known to work with AWS
// Elasticsearch instances. v7.14.0 and higher will fail with this error message:
// "the client noticed that the server is not a supported distribution of Elasticsearch"
// Good times...

go 1.17

require (
	github.com/elastic/go-elasticsearch/v7 v7.13.0
	github.com/tidwall/gjson v1.9.1
)

require (
	github.com/tidwall/match v1.0.3 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
)
