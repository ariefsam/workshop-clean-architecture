package spending

import (
	"encoding/json"
	"errors"
	"expense-tracker/service/activity"
	"time"

	"github.com/ariefsam/eventsam/idgenerator"
)

type spending struct {
	repo repository
}

type Spending struct {
	SpendingID string
	Name       string
	Amount     float64
}

type repository interface {
	Save(id string, tableName string, data any) (err error)
	Get(id string, tableName string, data any) (err error)
}

func NewSpending(repo repository) (s *spending) {
	s = &spending{
		repo: repo,
	}
	return
}

func (s *spending) Create(id, name string, amount float64) (err error) {

	if amount <= 0 {
		err = errors.New("amount must be greater than 0")
		return
	}

	spending := Spending{
		SpendingID: id,
		Name:       name,
		Amount:     amount,
	}
	err = s.repo.Save(id, "spending", spending)
	if err != nil {
		return
	}

	dataSpending, _ := json.Marshal(spending)

	activityData := activity.Activity{
		ActivityID: idgenerator.Generate(),
		Name:       "Create Spending",
		Data:       string(dataSpending),
		Time:       time.Now(),
	}
	err = s.repo.Save(activityData.ActivityID, "activity", activityData)

	return
}
