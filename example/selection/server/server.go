package main

import (
	"log"
	"net/http"

	"git.sr.ht/~sircmpwn/gqlgen/example/selection"
	"git.sr.ht/~sircmpwn/gqlgen/graphql/handler"
	"git.sr.ht/~sircmpwn/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Selection Demo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(selection.NewExecutableSchema(selection.Config{Resolvers: &selection.Resolver{}})))
	log.Fatal(http.ListenAndServe(":8086", nil))
}
