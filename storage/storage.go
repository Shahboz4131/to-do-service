package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/Shahboz4131/template-service/storage/postgres"
	"github.com/Shahboz4131/template-service/storage/repo"
)

// IStorage ...
// type IStorage interface {
// 	User() repo.UserStorageI
// }

// IStorage ...
type IStorage interface {
	Task() repo.TaskStorageI
}

// type storagePg struct {
// 	db       *sqlx.DB
// 	userRepo repo.UserStorageI
// }

type storagePg struct {
	db       *sqlx.DB
	taskRepo repo.TaskStorageI
}

// NewStoragePg ...
// func NewStoragePg(db *sqlx.DB) *storagePg {
// 	return &storagePg{
// 		db:       db,
// 		userRepo: postgres.NewUserRepo(db),
// 	}
// }

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		taskRepo: postgres.NewTaskRepo(db),
	}
}

// func (s storagePg) User() repo.UserStorageI {
// 	return s.userRepo
// }

func (s storagePg) Task() repo.TaskStorageI {
	return s.taskRepo
}
