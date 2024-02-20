package user

import (
	"errors"
	"time"

	"app/db"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IUserService interface {
	GetUserByID(ID string) (db.User, error)
	CreateUser(email, salt string, hashedPassword string) (db.User, error)
}

type UserService struct {
	db *sqlx.DB
}

func (s *UserService) GetUserByID(ID string) (db.User, error) {
	user := db.User{}
	err := s.db.Get(&user, `
    FROM user
    SELECT id, email
    WHERE id=$1
    VALUES=($1)
  `, ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *UserService) CreateUser(email string, salt string, hashedPassword string) (db.User, error) {
	defer func() (*db.User, error) {
		if r := recover(); r != nil {
			return nil, errors.New("panic: createUser errored likely due to uuid generation")
		}

		return nil, nil
	}()

	id := uuid.NewString()
	user := db.User{}
	err := s.db.Get(&user, `
    INSERT INTO user (
      id,
      email,
      salt,
      password,
      created_at
    ) VALUES ($1, $2, $3, $4, $5)
    RETURNING (id, email)
  `, id, email, salt, hashedPassword, time.Now().UTC().String())
	if err != nil {
		return user, err
	}

	return user, nil
}

func New(db *sqlx.DB) *UserService {
	return &UserService{
		db: db,
	}
}
