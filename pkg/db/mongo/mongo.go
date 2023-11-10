package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewClient(username, password string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(fmt.Sprintf(
		"mongodb+srv://%s:%s@cluster0.jetskns.mongodb.net/?retryWrites=true&w=majority",
		username, password),
	)
	if username != "" && password != "" {
		opts.SetAuth(options.Credential{
			Username: username,
			Password: password,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CloseClient(client *mongo.Client) error {
	return client.Disconnect(context.Background())
}
