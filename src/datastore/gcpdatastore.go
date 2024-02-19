package gcpdatastore

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

// import (
// 	"context"
// 	""
// )

type Task struct {
	Description string
	Done        bool
}

func CreateClient(ctx context.Context, projectId string) *datastore.Client {
	client, err := datastore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create GCP Datastore client: %v", err)
	}
	return client
}
