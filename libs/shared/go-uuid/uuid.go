package gouuid

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// ID is a unique identifier generated based on the provided properties.
type ID = string

// GetID generates a unique ID based on the provided properties.
// It serializes the properties to JSON, hashes the resulting byte slice, and generates a UUID from the hash.
// Returns the generated ID and any error encountered during the process.
func GetID(properties map[string]interface{}) (string, error) {
	serializedConfigProperties, err := json.Marshal(properties)
	if err != nil {
		return "", fmt.Errorf("error marshaling config properties: %w", err)
	}

	configHash := hashConfig(serializedConfigProperties)
	configID := generateUUIDFromHash(configHash)

	return ID(configID), nil
}

// hashConfig computes the SHA-256 hash of the given data.
// Returns the hash as a byte slice.
func hashConfig(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// generateUUIDFromHash generates a UUID based on the provided hash.
// It computes a second SHA-256 hash of the input hash and then generates a UUID using the SHA-1 algorithm.
// Returns the generated UUID as a string.
func generateUUIDFromHash(hash []byte) string {
	combinedHash := sha256.Sum256(hash)
	return uuid.NewSHA1(uuid.Nil, combinedHash[:]).String()
}
