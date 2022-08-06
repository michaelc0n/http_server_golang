package database

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

func (c Client) CreatePost(userEmail, text string) (Post, error) {
	db, err := c.readDB()
	if err != nil {
		return Post{}, err
	}
	if _, ok := db.Users[userEmail]; !ok {
		return Post{}, errors.New("user doesn't exist")
	}
	id := uuid.New().String()
	newPost := Post{
		CreatedAt: time.Now().UTC(),
		ID:        id,
		Text:      text,
		UserEmail: userEmail,
	}
	db.Posts[id] = newPost
	err = c.updateDB(db)
	if err != nil {
		return Post{}, err
	}

	return newPost, err

}
