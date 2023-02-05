package repository

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteRepository struct {
	db *gorm.DB
}

type sqliteRepo interface {
	PrimaryKey() string
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
	sqliteData, ok := data.(sqliteRepo)
	if !ok {
		err = errors.New("data must be implement PrimaryKey()")
		return
	}

	err = s.db.Table(tableName).First(data, sqliteData.PrimaryKey()+"=?", id).Error
	return
}
