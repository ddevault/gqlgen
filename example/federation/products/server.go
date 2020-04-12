//go:generate go run ../../../testdata/gqlgen.go
package main

import (
	"log"
	"net/http"
	"os"

	"git.sr.ht/~sircmpwn/gqlgen/example/federation/products/graph"
	"git.sr.ht/~sircmpwn/gqlgen/example/federation/products/graph/generated"
	"git.sr.ht/~sircmpwn/gqlgen/graphql/handler"
	"git.sr.ht/~sircmpwn/gqlgen/graphql/handler/debug"
	"git.sr.ht/~sircmpwn/gqlgen/graphql/playground"
)

const defaultPort = "4002"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.Use(&debug.Tracer{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
