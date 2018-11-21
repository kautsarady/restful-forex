package model_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/kautsarady/forex/model"
)

func TestModel(t *testing.T) {

	// Make database connection
	dao, err := model.Make(
		fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DBNAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")),
		true)
	if err != nil {
		log.Fatalf("Failed to make database connection: %v", err)
	}
	defer dao.DB.Close()

	tt := []string{
		"Add",
		"Input",
		"Track",
		"Trend",
		"Remove",
	}

	for _, tc := range tt {
		t.Run(tc, func(t *testing.T) {
			switch tc {
			case "Add":
				if err := dao.Add(model.Exchange{From: "IDR", To: "KRW"}); err != nil {
					t.Errorf("Error adding exchange: %v", err)
				}
				break
			case "Input":
				if err := dao.Input(
					model.Exchange{
						From: "IDR",
						To:   "KRW",
						Rate: 15.667,
						Date: time.Now(),
					}); err != nil {
					t.Errorf("Error input exchange: %v", err)
				}
				break
			case "Track":
				exs, err := dao.Track(time.Now())
				if err != nil || exs == nil {
					t.Errorf("Unexpected error/result tracking exchange: %v", err)
				}
				break
			case "Trend":
				exs, err := dao.Trend("IDR", "KRW", 15.5)
				if err != nil || exs == nil {
					t.Errorf("Unexpected error/result queriying trend exchange: %v", err)
				}
				break
			case "Remove":
				if err := dao.Remove(model.Exchange{From: "IDR", To: "KRW"}); err != nil {
					t.Errorf("Error removing exchange: %v", err)
				}
				break
			}
		})
	}
}
