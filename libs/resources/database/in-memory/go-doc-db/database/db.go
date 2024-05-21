package database

import "errors"

// InMemoryDocBD represents an in-memory document database containing multiple collections.
type InMemoryDocBD struct {
	Name        string
	Collections map[string]*Collection
}

// NewInMemoryDocBD creates and returns a new InMemoryDocBD instance with the given name.
func NewInMemoryDocBD(name string) *InMemoryDocBD {
	return &InMemoryDocBD{
		Name:        name,
		Collections: make(map[string]*Collection),
	}
}

// GetCollection retrieves a collection by its name.
func (d *InMemoryDocBD) GetCollection(collectionName string) (*Collection, error) {
	collection, ok := d.Collections[collectionName]
	if !ok {
		return nil, errors.New("collection not found")
	}
	return collection, nil
}

// CreateCollection creates a new collection with the given name.
func (d *InMemoryDocBD) CreateCollection(collectionName string) error {
	if _, ok := d.Collections[collectionName]; ok {
		return errors.New("collection already exists")
	}
	d.Collections[collectionName] = NewCollection()
	return nil
}

// DropCollection drops a collection by its name.
func (d *InMemoryDocBD) DropCollection(collectionName string) error {
	if _, ok := d.Collections[collectionName]; !ok {
		return errors.New("collection not found")
	}
	delete(d.Collections, collectionName)
	return nil
}

// ListCollections lists the names of all collections in the database.
func (d *InMemoryDocBD) ListCollections() []string {
	collectionNames := make([]string, 0, len(d.Collections))
	for collectionName := range d.Collections {
		collectionNames = append(collectionNames, collectionName)
	}
	return collectionNames
}
