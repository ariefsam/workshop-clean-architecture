package spending_test

import (
	"expense-tracker/repository"
	"expense-tracker/service/expense-tracker/spending"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSpendingSimple(t *testing.T) {
	mockRepository := repository.NewMock()
	ucSpending := spending.NewSpending(mockRepository)

	ucSpending.Create("id001", "Belanja bulanan", 1000000)

	assert.True(t, mockRepository.IsSaved("spending"))
	assert.True(t, mockRepository.IsSaved("activity"))

}

func TestCreateSpendingErrorAmount(t *testing.T) {
	mockRepository := repository.NewMock()
	ucSpending := spending.NewSpending(mockRepository)

	err := ucSpending.Create("id001", "Belanja bulanan", -1000000)
	assert.Error(t, err)
}
