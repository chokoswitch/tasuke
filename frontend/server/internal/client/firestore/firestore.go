package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

// Client is an interface wrapping the Firebase Admin SDK for our usage of firestore.
// It does not use a fluent API to make mocking easier.
type Client interface {
	// SetDocument creates a document in the given collection at the given path with the given data.
	SetDocument(ctx context.Context, collection string, path string, data any) error
}

// NewClient returns a new Client wrapping the given firestore.Client.
func NewClient(store *firestore.Client) Client {
	return &client{
		store: store,
	}
}

type client struct {
	store *firestore.Client
}

// SetDocument implements Client.
func (c *client) SetDocument(ctx context.Context, collection string, path string, data any) error {
	if _, err := c.store.Collection(collection).Doc(path).Set(ctx, data); err != nil {
		return fmt.Errorf("firestore: failed to create document: %w", err)
	}

	return nil
}
