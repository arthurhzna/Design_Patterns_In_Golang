package main
package main

import "fmt"

// =========================
// Business side
// =========================

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	GetUserByID(id int) User
}

func PrintUser(repo UserRepository, id int) {
	user := repo.GetUserByID(id)
	fmt.Println("user:", user.ID, user.Name)
}

// =========================
// External / existing DB
// =========================

type LegacyDB struct{}

func (db *LegacyDB) QueryUser(id int) map[string]any {
	return map[string]any{
		"id":   id,
		"name": "Arthur",
	}
}

// =========================
// Adapter
// =========================

type LegacyDBAdapter struct {
	db *LegacyDB
}

func (a *LegacyDBAdapter) GetUserByID(id int) User {
	row := a.db.QueryUser(id)

	return User{
		ID:   row["id"].(int),
		Name: row["name"].(string),
	}
}

// =========================
// Main
// =========================

func main() {
	legacyDB := &LegacyDB{}              // external
	repo := &LegacyDBAdapter{db: legacyDB} // adapter

	PrintUser(repo, 1) // business logic 
}