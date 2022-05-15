package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"whitesource_home_assingment/internal/movie"
	"whitesource_home_assingment/internal/rest"
	"whitesource_home_assingment/internal/storage/mongo"
	"whitesource_home_assingment/internal/storage/postgres"
)

const (
	dbMongo    = "mongo"
	dbPostgres = "postgres"
)

func main() {
	var db string
	flag.StringVar(&db, "database", dbMongo, "select application database")
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var r movie.Repository
	switch db {
	case dbMongo:
		r = mongo.NewStorage(ctx, "mongodb://mongodb:27017")
	case dbPostgres:
		r = postgres.NewStorage(ctx, "postgresql://user@postgres:5432/movies?sslmode=disable")
	default:
		panic("unexpected value")
	}
	server := rest.Server{
		MovieService: movie.Service{
			MovieRepository: r,
		},
	}
	server.Run(
		gin.Default(),
		":443",
		"/etc/tls/cert.pem",
		"/etc/tls/key.pem",
	)
}
