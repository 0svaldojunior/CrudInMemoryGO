package database

import (
	"CrudInMemoryGo/log"

	"github.com/google/uuid"
)

type id uuid.UUID

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

type database struct {
	data map[string]User
}

var databaseInstance *database

func NewDatabase() *database {
	if databaseInstance == nil {
		databaseInstance = &database{
			data: make(map[string]User),
		}
	}

	return databaseInstance
}

func (database *database) CreateUser(user User) string {
	byteArray := id(uuid.New())

	id, err := uuid.FromBytes(byteArray[:])
	if err != nil {
		log.Error("Error creating user ID:", err)
	}

	stringID := id.String()
	user.ID = stringID
	database.data[stringID] = user

	return stringID
}

func (database *database) GetUser(id string) (User, bool) {
	user, ok := database.data[id]
	return user, ok
}

func (database *database) UpdateUser(id string, user User) bool {
	existingUser, ok := database.data[id]

	if !ok {
		return false
	}

	if len(user.FirstName) > 0 {
		existingUser.FirstName = user.FirstName
	}
	if len(user.LastName) > 0 {
		existingUser.LastName = user.LastName
	}
	if len(user.Biography) > 0 {
		existingUser.Biography = user.Biography
	}

	database.data[id] = existingUser

	return true
}

func (database *database) DeleteUser(id string) bool {
	if _, ok := database.data[id]; !ok {
		return false
	}

	delete(database.data, id)
	return true
}

func (database *database) GetAllUsers() []User {
	users := make([]User, 0, len(database.data))
	for _, user := range database.data {
		users = append(users, user)
	}

	return users
}
