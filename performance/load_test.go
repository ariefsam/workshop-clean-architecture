package performance

import (
	"expense-tracker/idgenerator"
	"expense-tracker/repository"
	"expense-tracker/service/activity"
	"expense-tracker/service/expense-tracker/spending"
	"log"
	"testing"
	"time"

	"gorm.io/gorm"
)

type spendingGorm struct {
	gorm.Model
	spending.Spending
}

func (s spendingGorm) TableName() string {
	return "spending"
}

type spendingActivityGorm struct {
	gorm.Model
	activity.Activity
}

func (s spendingActivityGorm) TableName() string {
	return "activity"
}

func TestLoadSQLite(t *testing.T) {
	repoService, err := repository.NewSQLite("tbl.db")
	if err != nil {
		log.Println(err)
	}
	modelSpending := spendingGorm{}
	modelActivity := spendingActivityGorm{}
	repoService.AutoMigrate(&modelSpending)
	repoService.AutoMigrate(&modelActivity)

	start := time.Now()
	for i := 0; i < 100000; i++ {
		spendingService := spending.NewSpending(repoService)
		id := idgenerator.Generate()
		spendingService.Create(id, "Belanja bulanan "+id, float64(i)*100)
	}
	end := time.Now()
	log.Println(end.Sub(start), " seconds")
}
