package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"git.sr.ht/~sircmpwn/gqlgen/example/todo"
	"git.sr.ht/~sircmpwn/gqlgen/graphql/handler"
	"git.sr.ht/~sircmpwn/gqlgen/graphql/playground"
)

func main() {
	srv := handler.NewDefaultServer(todo.NewExecutableSchema(todo.New()))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		// send this panic somewhere
		log.Print(err)
		debug.PrintStack()
		return errors.New("user message on panic")
	})

	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
