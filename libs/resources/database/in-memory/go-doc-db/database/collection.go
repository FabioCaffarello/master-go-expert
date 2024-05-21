package database

import (
	"errors"
	"sync"
)

// DocumentID represents the unique identifier for a document.
type DocumentID string

// Document represents a document as a map with string keys and interface{} values.
type Document map[string]interface{}

// Collection represents a collection of documents with thread-safe operations.
type Collection struct {
	data map[string]Document
	mu   sync.RWMutex
}

// NewCollection creates and returns a new Collection instance.
func NewCollection() *Collection {
	return &Collection{
		data: make(map[string]Document),
	}
}

// InsertOne inserts a single document into the collection.
func (c *Collection) InsertOne(document Document) error {
	c.mu.Lock() // Lock for writing
	defer c.mu.Unlock()

	documentID, ok := document["_id"]
	if !ok {
		return errors.New("_id field is required")
	}
	documentIDStr, ok := documentID.(string)
	if !ok {
		return errors.New("_id field must be a string")
	}

	if _, ok := c.data[documentIDStr]; ok {
		return errors.New("document already exists")
	}
	c.data[documentIDStr] = document
	return nil
}

// FindOne finds and returns a single document by its ID.
func (c *Collection) FindOne(id string) (Document, error) {
	c.mu.RLock() // Lock for reading
	defer c.mu.RUnlock()

	document, ok := c.data[id]
	if !ok {
		return nil, errors.New("document not found")
	}
	return document, nil
}

// FindAll returns all documents in the collection.
func (c *Collection) FindAll() []Document {
	c.mu.RLock() // Lock for reading
	defer c.mu.RUnlock()
	documents := make([]Document, 0, len(c.data))
	for _, document := range c.data {
		documents = append(documents, document)
	}
	return documents
}

// matchesQuery checks if a document matches the query criteria.
func matchesQuery(document, query map[string]interface{}) bool {
	for key, value := range query {
		docValue, exists := document[key]
		if !exists {
			return false
		}

		// If the value is a map, recurse into it.
		if queryMap, ok := value.(map[string]interface{}); ok {
			docMap, ok := docValue.(map[string]interface{})
			if !ok || !matchesQuery(docMap, queryMap) {
				return false
			}
		} else {
			// Direct comparison for non-map values
			if docValue != value {
				return false
			}
		}
	}
	return true
}

// Find searches documents matching a given query.
func (c *Collection) Find(query map[string]interface{}) []Document {
	c.mu.RLock() // Lock for reading
	defer c.mu.RUnlock()
	documents := make([]Document, 0, len(c.data))
	for _, document := range c.data {
		if matchesQuery(document, query) {
			documents = append(documents, document)
		}
	}
	return documents
}

// DeleteOne deletes a single document by its ID.
func (c *Collection) DeleteOne(id string) error {
	c.mu.Lock() // Lock for writing
	defer c.mu.Unlock()
	_, ok := c.data[id]
	if !ok {
		return errors.New("document not found")
	}
	delete(c.data, id)
	return nil
}

// UpdateOne updates a single document by its ID with the given update.
func (c *Collection) UpdateOne(id string, update Document) error {
	c.mu.Lock() // Lock for writing
	defer c.mu.Unlock()
	_, ok := c.data[id]
	if !ok {
		return errors.New("document not found")
	}
	for key, value := range update {
		c.data[id][key] = value
	}
	return nil
}

// DeleteAll deletes all documents in the collection.
func (c *Collection) DeleteAll() error {
	c.mu.Lock() // Lock for writing
	defer c.mu.Unlock()
	c.data = make(map[string]Document)
	return nil
}
