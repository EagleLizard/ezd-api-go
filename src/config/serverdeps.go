package config

import (
	"context"

	"cloud.google.com/go/datastore"
)

type ServerDeps struct {
	context            context.Context
	GcpDataStoreClient *datastore.Client
}
