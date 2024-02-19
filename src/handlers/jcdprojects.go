package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/EagleLizard/ezd-api-go/config"
)

func GetProjects(sd config.ServerDeps, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	q := datastore.NewQuery("JcdProjectKeyV3").KeysOnly()
	keys, err := sd.GcpDataStoreClient.GetAll(ctx, q, nil)
	if err != nil {
		log.Fatalf("client.GetAll: %v", err)
	}
	log.Printf("%v", keys)
	projectKeyNames := make([]string, 0)
	for _, projectKey := range keys {
		log.Printf("Name: %v", projectKey.Name)
		projectKeyNames = append(projectKeyNames, projectKey.Name)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projectKeyNames)
	// fmt.Fprint(w, "get projects")
}
