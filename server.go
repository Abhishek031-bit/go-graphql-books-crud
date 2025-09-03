package main

import (
	"go-graphql/database"
	"go-graphql/graph"
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const defaultPort = "8080"

func main() {
	database.ConnectDB()
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{
		Max:        10,               // max 10 requests
		Expiration: 30 * time.Second, // per 30 seconds
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests, slow down!",
			})
		},
	}))

	app.Get("/hi", func(c *fiber.Ctx) error {
		return c.SendString("hi")
	})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	app.All("/query", adaptor.HTTPHandler(srv))
	app.Get("/", adaptor.HTTPHandler(playground.Handler("GraphQL playground", "/query")))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(app.Listen(":" + port))

	// srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// srv.AddTransport(transport.Options{})
	// srv.AddTransport(transport.GET{})
	// srv.AddTransport(transport.POST{})

	// srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	// srv.Use(extension.Introspection{})
	// srv.Use(extension.AutomaticPersistedQuery{
	// 	Cache: lru.New[string](100),
	// })

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
