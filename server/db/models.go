package db

import "time"

type User struct {
	Id        string    
	Email     string    
	Password  string    
  CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Sites struct {
	Id        string    
	Name      string    
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
  TreeId string `db:"tree_id"`
}

type SiteDomTrees struct {
  Id string
  Tree string
}
