package report

import (
	"errors"
	"log"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	ActivityID string
	Name       string
	Data       string
	Time       time.Time
}

var ActivityData map[string]Activity
var lastID uint

func loadData() {
	if ActivityData == nil {
		ActivityData = make(map[string]Activity)
	}
	filename := "../performance/tbl.db"
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		return
	}

	activities := []Activity{}
	err = db.Table("activity").Where("id", ">", lastID).Find(&activities).Error
	if err != nil {
		log.Println(err)
		return
	}
	for _, activity := range activities {
		ActivityData[activity.ActivityID] = activity
		lastID = activity.ID
	}
}

func TotalByAmount() (total float64) {
	loadData()

	for x := 0; x < 1000; x++ {
		start := time.Now().UnixNano()
		dat, ok := ActivityData["9FJGL-AVRz"]
		if ok {
			log.Println(dat)
		}
		end := time.Now().UnixNano()
		log.Println(end-start, " nano seconds")

		start = time.Now().UnixMilli()
		activitySearch, err := SearchActivityMap("OpJ")
		if err != nil {
			log.Println(err)
		} else {
			log.Println(activitySearch)
		}
		end = time.Now().UnixMilli()
		log.Println(end-start, " milli seconds")
	}
	return
}

// search from map string activity
func SearchActivityMap(name string) (activity Activity, err error) {
	for _, activity := range ActivityData {
		// log.Println(activity.Name, name)

		if strings.Contains(activity.Name, name) {
			return activity, nil
		}
	}
	err = errors.New("not found")
	return
}
