package database

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

// Post
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

type Client struct {
	Path string
}

func NewClient(path string) Client {
	return Client{
		Path: path,
	}
}

// EnsureDB creates the database file if it doesn't exist
func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.Path)
	if errors.Is(err, os.ErrNotExist) {
		return c.createDB()
	}
	return err
}

func (c Client) updateDB(db databaseSchema) error {
	data, err := json.Marshal(db)
	if err != nil {
		return err
	}
	err = os.WriteFile(c.Path, data, 0600)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) readDB() (databaseSchema, error) {
	data, err := os.ReadFile(c.Path)
	if err != nil {
		return databaseSchema{}, err
	}
	db := databaseSchema{}
	err = json.Unmarshal(data, &db)
	//fmt.Println(db, err)
	return db, err
}

func (c Client) createDB() error {
	data, err := json.Marshal(databaseSchema{
		Users: make(map[string]User),
		Posts: make(map[string]Post),
	})
	if err != nil {
		return err
	}
	err = os.WriteFile(c.Path, data, 0600)
	if err != nil {
		return err
	}
	return nil
}
