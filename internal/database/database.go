package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

// User -
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Another   string    `json:"another"`
}

// Post -
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

func (c Client) CreateUser(email, password, name string, age int, another string) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err
	}
	if _, ok := db.Users[email]; ok {
		fmt.Printf("creating user: %v", ok)
		return User{}, errors.New("user already exists")
	}
	user := User{
		CreatedAt: time.Now().UTC(),
		Email:     email,
		Password:  password,
		Name:      name,
		Age:       age,
		Another:   another,
	}
	db.Users[email] = user

	err = c.updateDB(db)
	if err != nil {
		return User{}, err
	}
	fmt.Printf("creating user: %v %v", user, err)
	return user, err
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err
	}
	user, ok := db.Users[email]

	if !ok {
		return User{}, errors.New("user doesn't exist")
	}

	user.Password = password
	user.Name = name
	user.Age = age
	db.Users[email] = user
	err = c.updateDB(db)
	if err != nil {
		return User{}, err
	}
	fmt.Printf("updating user: %v %v", user, err)
	return user, err
}

func (c Client) GetUser(email string) (User, error) {
	db, err := c.readDB()
	user, ok := db.Users[email]
	if !ok {
		return User{}, errors.New("user doesn't exist")
	}
	fmt.Printf("getting user: %v %v", user, err)
	return user, err
}

func (c Client) DeleteUser(email string) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}
	delete(db.Users, email)
	err = c.updateDB(db)
	if err != nil {
		return err
	}

	return nil
}
