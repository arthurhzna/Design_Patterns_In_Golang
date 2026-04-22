package main

import (
	"fmt"
)

// Subject
type Database interface {
	GetUser(id int) string
}

// Real Subject
type RealDB struct{}

func (r *RealDB) GetUser(id int) string {
	fmt.Println("querying real database...")
	return "Arthur"
}

// Proxy
type DBProxy struct {
	real *RealDB
}

func (p *DBProxy) GetUser(id int) string {
	fmt.Println("checking cache...")

	// additional behavior before accessing the real database
	if id == 1 {
		fmt.Println("cache hit")
		return "Arthur (from cache)"
	}

	fmt.Println("cache miss")
	return p.real.GetUser(id)
}

func main() {
	var db Database = &DBProxy{
		real: &RealDB{},
	}

	fmt.Println(db.GetUser(1))
	fmt.Println("------")
	fmt.Println(db.GetUser(2))
}

// Proxy

// A wrapper that controls access to another object.

// Used for:

// authentication
// authorization
// caching
// lazy loading
// remote communication

//----------------------------------------------------------

// Decorator

// A wrapper that adds new behavior/features to another object.

// Used for:

// logging
// metrics
// compression
// formatting
// extra responsibilities

// Proxy     = gatekeeper
// Decorator = enhancer
