package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/EagleLizard/ezd-api-go/config"
	gcpdatastore "github.com/EagleLizard/ezd-api-go/datastore"
	"github.com/EagleLizard/ezd-api-go/handlers"
)

func main() {
	cfg := config.LoadConfig()

	ctx := initCtx()
	sd := initServerDeps(ctx, cfg)

	gcpDsClient := gcpdatastore.CreateClient(ctx, cfg.GcpProjectId)
	query := datastore.NewQuery("JcdProjectKeyV3").KeysOnly()
	keys, err := gcpDsClient.GetAll(ctx, query, nil)
	if err != nil {
		log.Fatalf("client.GetAll: %v", err)
	}
	log.Printf("%v", keys)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "EZD API, route: %s\n", r.URL.Path)
	})

	http.HandleFunc("/v1/health", handlers.HealthCheckHandler)

	http.HandleFunc("/v1/projects", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetProjects(sd, w, r)
	})

	fmt.Printf("Server is starting on port %s...\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}

func initCtx() context.Context {
	ctx := context.Background()
	return ctx
}

func initServerDeps(ctx context.Context, cfg config.Config) config.ServerDeps {
	sd := config.ServerDeps{
		GcpDataStoreClient: gcpdatastore.CreateClient(ctx, cfg.GcpProjectId),
	}
	return sd
}
