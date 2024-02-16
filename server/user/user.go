package user

import (
	"errors"
	"time"

	"app/db"

	"github.com/google/uuid"
)

func createUser(email string, hashedPassword string) (*db.User, error) {
	defer func() (*db.User, error) {
		if r := recover(); r != nil {
			return nil, errors.New("panic: createUser errored likely due to uuid generation")
		}

		return nil, nil
	}()

	id := uuid.NewString()
	user := &db.User{}
	err := db.DB.Get(user, `
    INSERT INTO user (
      id,
      email,
      password,
      created_at
    ) VALUES ($1, $2, $3, $3)
    RETURNING (id, email)
  `, id, email, hashedPassword, time.Now().UTC().String())
	if err != nil {
		return nil, err
	}

	return user, nil
}
