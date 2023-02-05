package repository_test

import (
	"expense-tracker/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type dataTest struct {
	gorm.Model
	DataID string `gorm:"index"`
	Name   string
}

// get tablename of dataTest
func (d dataTest) TableName() string {
	return "data_test"
}

func (d dataTest) PrimaryKey() string {
	return "data_id"
}

func TestSaveGet(t *testing.T) {
	repo, err := repository.NewSQLite("file::memory:?cache=shared")
	assert.NoError(t, err)

	data := dataTest{
		DataID: "id001",
		Name:   "Arief",
	}

	// data migration
	err = repo.AutoMigrate(data)
	assert.NoError(t, err)

	// save data
	err = repo.Save(data.DataID, "data_test", &data)
	assert.NoError(t, err)

	// get data
	var dataGet dataTest
	err = repo.Get(data.DataID, "data_test", &dataGet)
	assert.NoError(t, err)
	assert.Equal(t, data.DataID, dataGet.DataID)
	assert.Equal(t, data.Name, dataGet.Name)

}
