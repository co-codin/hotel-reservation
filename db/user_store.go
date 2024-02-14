package db

import "github.com/co-codin/hotel-reservation/types"

type UserStore interface {
	GetUserByID(string) (*types.User, error)
}

type MongoUserStore struct {
}
