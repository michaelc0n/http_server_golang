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

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	db, err := c.readDB()
	if err != nil {
		return nil, err
	}
	posts := []Post{}
	for _, post := range db.Posts {
		if post.UserEmail == userEmail {
			posts = append(posts, post)
		}
	}
	return posts, nil
}

func (c Client) DeletePost(id string) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}
	delete(db.Posts, id)
	err = c.updateDB(db)
	if err != nil {
		return err
	}
	return nil
}
