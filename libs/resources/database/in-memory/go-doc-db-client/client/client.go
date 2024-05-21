package client

import (
	"errors"
	"fmt"
	"libs/resources/database/in-memory/go-doc-db/database"
)

// Client provides an interface to interact with the in-memory document database.
type Client struct {
	db *database.InMemoryDocBD
}

// NewClient creates and returns a new Client instance.
func NewClient(db *database.InMemoryDocBD) *Client {
	return &Client{
		db: db,
	}
}

// getCollection retrieves a collection by its name. Returns an error if the collection does not exist.
func (c *Client) getCollection(collectionName string) (*database.Collection, error) {
	collection, err := c.db.GetCollection(collectionName)
	if err != nil {
		return nil, fmt.Errorf("collection %s does not exist", collectionName)
	}
	return collection, nil
}

// CreateCollection creates a new collection with the given name. Returns an error if the collection already exists.
func (c *Client) CreateCollection(collectionName string) error {
	err := c.db.CreateCollection(collectionName)
	if err != nil {
		return err
	}
	return nil
}

// DropCollection drops a collection by its name. Returns an error if the collection does not exist.
func (c *Client) DropCollection(collectionName string) error {
	if _, exists := c.db.Collections[collectionName]; !exists {
		return fmt.Errorf("collection %s does not exist", collectionName)
	}
	err := c.db.DropCollection(collectionName)
	if err != nil {
		return err
	}
	return nil
}

// ListCollections lists the names of all collections in the database.
func (c *Client) ListCollections() []string {
	return c.db.ListCollections()
}

// ConvertToDocument converts a map to a Document type. Returns an error if the document is nil or the _id field is missing.
func (c *Client) ConvertToDocument(document map[string]interface{}) (database.Document, error) {
	if document == nil {
		return nil, errors.New("document is nil")
	}
	if _, ok := document["_id"]; !ok {
		return nil, errors.New("_id field is required")
	}
	return database.Document(document), nil
}

// InsertOne inserts a single document into the specified collection. Returns an error if the collection does not exist or the document is invalid.
func (c *Client) InsertOne(collectionName string, document map[string]interface{}) error {
	collection, err := c.getCollection(collectionName)
	if err != nil {
		return err
	}
	doc, err := c.ConvertToDocument(document)
	if err != nil {
		return err
	}
	return collection.InsertOne(doc)
}

// FindOne finds and returns a single document by its ID from the specified collection. Returns an error if the collection or document does not exist.
func (c *Client) FindOne(collectionName string, id string) (map[string]interface{}, error) {
	collection, err := c.getCollection(collectionName)
	if err != nil {
		return nil, err
	}
	doc, err := collection.FindOne(id)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}(doc), nil
}

// FindAll returns all documents from the specified collection. Returns an error if the collection does not exist.
func (c *Client) FindAll(collectionName string) ([]map[string]interface{}, error) {
	collection, err := c.getCollection(collectionName)
	if err != nil {
		return nil, err
	}
	docs := collection.FindAll()
	documents := make([]map[string]interface{}, 0, len(docs))
	for _, doc := range docs {
		documents = append(documents, map[string]interface{}(doc))
	}
	return documents, nil
}

// Find returns documents matching the given query from the specified collection. Returns an error if the collection does not exist.
func (c *Client) Find(collectionName string, filter map[string]interface{}) ([]map[string]interface{}, error) {
	collection, err := c.getCollection(collectionName)
	if err != nil {
		return nil, err
	}
	docs := collection.Find(filter)
	documents := make([]map[string]interface{}, 0, len(docs))
	for _, doc := range docs {
		documents = append(documents, map[string]interface{}(doc))
	}
	return documents, nil
}

// UpdateOne updates a single document by its ID with the given update in the specified collection. Returns an error if the collection or document does not exist.
func (c *Client) UpdateOne(collectionName string, id string, update map[string]interface{}) error {
	collection, err := c.getCollection(collectionName)
	if err != nil {
		return err
	}
	if len(update) == 0 {
		return errors.New("update is empty")
	}
	return collection.UpdateOne(id, update)
}

// DeleteOne deletes a single document by its ID from the specified collection. Returns an error if the collection or document does not exist.
func (c *Client) DeleteOne(collectionName string, id string) error {
	collection, err := c.getCollection(collectionName)
	if err != nil {
		return err
	}
	return collection.DeleteOne(id)
}

// DeleteAll deletes all documents from the specified collection. Returns an error if the collection does not exist.
func (c *Client) DeleteAll(collectionName string) error {
	collection, err := c.getCollection(collectionName)
	if err != nil {
		return err
	}
	return collection.DeleteAll()
}
