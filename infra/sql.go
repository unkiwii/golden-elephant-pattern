package infra

import (
	"errors"

	"github.com/unkiwii/golden-elephant-pattern/models"
)

type UserRepo struct{}

func (repo UserRepo) UserByID(id int) (models.User, error) {
	rows, err := db.Query(`SELECT * FROM users WHERE id = $1`, id)
	if err != nil {
		return models.User{}, err
	}

	if !rows.Next() {
		return models.User{}, errors.New("user not found")
	}

	var user models.User
	rows.Scan(&user.ID, &user.City, &user.Email)
	return user, nil
}

func (repo UserRepo) UsersByID(ids []int) ([]models.User, error) {
	var (
		users = make([]models.User, len(ids))
		err   error
	)
	for i, id := range ids {
		users[i], err = repo.UserByID(id)
		if err != nil {
			return users, err
		}
	}
	return users, nil
}
