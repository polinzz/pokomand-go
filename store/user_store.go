package Store

import (
	"database/sql"
	"pokomand-go/Entity"
)

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db,
	}
}

type UserStore struct {
	*sql.DB
}

func (u *UserStore) GetUsers() ([]Entity.User, error) {
	var users []Entity.User

	rows, err := u.Query("SELECT * FROM Users")
	if err != nil {
		return []Entity.User{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var user Entity.User
		if err = rows.Scan(&user.ID, &user.FirstName, &user.LastName); err != nil {
			return []Entity.User{}, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return []Entity.User{}, err
	}

	return users, nil
}
