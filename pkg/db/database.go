package database

import "fmt"

// Database is a key/value store
type Database struct {
	data map[string]string
}

type KeyNotFound struct {
	key string
}

func (m *KeyNotFound) Error() string {
	return fmt.Sprintf("key %q not found in database", m.key)
}

// New creates and initalizes an instance of Database
// initialSize allows pre-allocating memory
func New(initialSize int) *Database {
	db := Database{}
	db.data = make(map[string]string, initialSize)

	db.data["debug"] = "working"

	return &db
}

// Write accepts a key/value pair and inserts it into the database db
func (db *Database) Write(key string, value string) {
	db.data[key] = value
}

// Read accepts a key and return the key if it exists, otherwise KeyNotFoundError
func (db *Database) Read(key string) (string, error) {
	value, ok := db.data[key]
	if !ok {
		return "", &KeyNotFound{key}
	}

	return value, nil
}

// Read accepts a key and remove the key/value pair from the database db, otherwise KeyNotFoundError
func (db *Database) Delete(key string) error {
	if _, ok := db.data[key]; !ok {
		return &KeyNotFound{key}
	}

	delete(db.data, key)

	return nil
}

// Keys returns a slice of keys in database
func (db *Database) Keys() []string {
	keys := make([]string, len(db.data))
	for k := range db.data {
		keys = append(keys, k)
	}

	return keys
}
