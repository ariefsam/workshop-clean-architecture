package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteRepository struct {
	db *gorm.DB
}

func NewSQLite(filename string) (repo *sqliteRepository, err error) {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		return
	}
	repo = &sqliteRepository{
		db: db,
	}
	return
}

func (s *sqliteRepository) Save(id string, tableName string, data any) (err error) {

	return
}

func (s *sqliteRepository) Get(id string, tableName string, data any) (err error) {

	return
}
