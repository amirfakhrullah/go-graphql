package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/amirfakhrullah/go-graphql/database"
	"github.com/amirfakhrullah/go-graphql/env"
	"github.com/amirfakhrullah/go-graphql/graph/generated"
	resolvers "github.com/amirfakhrullah/go-graphql/graph/resolvers"
)

func main() {
	env.InitEnv()
	db, err := database.InitDB()
	if err != nil {
		panic(err.Error())
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	customCtx := &database.CustomContext{
		Database: db,
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", database.CreateContext(customCtx, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", env.PORT)
	log.Fatal(http.ListenAndServe(":"+env.PORT, nil))
}
