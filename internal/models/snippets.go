package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define a Snippet type to hold the data for an individual snippet.
type Snippet struct {
	ID      string    `bson:"_id,omitempty"`
	UserID  string    `bson:"user_id"`
	Title   string    `bson:"title"`
	Content string    `bson:"content"`
	Created time.Time `bson:"created"`
	Expires time.Time `bson:"expires"`
	Public  bool      `bson:"public"`
}

// Define a SnippetModel type which wraps a MongoDB client.
type SnippetModel struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

// This will insert a new snippet into the database.
func (m *SnippetModel) Insert(user_id string, title string, content string, expires int) (string, error) {
	snippet := &Snippet{
		UserID:  user_id,
		Title:   title,
		Content: content,
		Created: time.Now(),
		Expires: time.Now().Add(time.Duration(expires) * time.Second),
		Public:  false,
	}

	result, err := m.Collection.InsertOne(context.TODO(), snippet)
	if err != nil {

		return "", err
	}

	insertedID := result.InsertedID.(primitive.ObjectID)
	return insertedID.String(), nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id string) (*Snippet, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	snippet := &Snippet{}
	err = m.Collection.FindOne(context.TODO(), filter).Decode(&snippet)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
		return nil, err
	}

	return snippet, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"created": -1})
	findOptions.SetLimit(10)

	filter := bson.M{}

	cur, err := m.Collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var snippets []*Snippet
	for cur.Next(context.TODO()) {
		var snippet Snippet
		err := cur.Decode(&snippet)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, &snippet)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}

func (m *SnippetModel) Delete(id string) error {
	filter := bson.M{"_id": id}

	_, err := m.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (m *SnippetModel) Share(id string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"public": true}}

	_, err := m.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
