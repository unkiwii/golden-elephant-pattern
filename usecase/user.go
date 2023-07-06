package usecase

import (
	"fmt"

	"github.com/unkiwii/golden-elephant-pattern/infra"
	"github.com/unkiwii/golden-elephant-pattern/models"
)

type UserProvider struct{}

func (provider UserProvider) UsersByID(ids ...int64) ([]models.User, error) {
	for i := 0; i < len(ids); i++ {
		id := ids[i]
		if id < 0 {
			return nil, fmt.Errorf("invalid id '%d'", id)
		}
	}

	uids := []int{}
	for _, id := range ids {
		uids = append(uids, int(id))
	}

	return infra.UserRepo{}.UsersByID(uids)
}
