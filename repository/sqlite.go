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

func (s *sqliteRepository) AutoMigrate(data any) (err error) {
	err = s.db.AutoMigrate(data)
	return
}

func (s *sqliteRepository) Save(id string, tableName string, data any) (err error) {
	err = s.db.Table(tableName).Create(data).Error
	return
}

func (s *sqliteRepository) Get(id string, tableName string, data any) (err error) {
	err = s.db.Table(tableName).First(data, "data_id = ?", id).Error
	return
}
