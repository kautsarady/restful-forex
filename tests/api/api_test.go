package api_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kautsarady/forex/api"
	"github.com/kautsarady/forex/model"
	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	tt := []struct {
		tName   string
		tMethod string
		tURL    string
		tExpect int
	}{
		{"input", "POST", fmt.Sprintf("/api/input/%s/%s?rate=%f&date=%s", "AUD", "USD", 0.73, "2018-11-16"), http.StatusOK},
		{"trend", "GET", fmt.Sprintf("/api/trend/%s/%s?avg=%f&vrn=%f", "AUD", "USD", 0.73, 0.01), http.StatusOK},
		{"remove1", "DELETE", fmt.Sprintf("/api/remove/%s/%s", "AUD", "USD"), http.StatusOK},
		{"track", "GET", fmt.Sprintf("/api/track?date=%s", "2018-11-16"), http.StatusOK},
		{"add", "POST", fmt.Sprintf("/api/add/%s/%s", "JMC", "JPY"), http.StatusOK},
		{"remove2", "DELETE", fmt.Sprintf("/api/remove/%s/%s", "JMC", "JPY"), http.StatusOK},
	}

	// Make database connection
	dao, err := model.Make(
		fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DBNAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")),
		true)
	if err != nil {
		log.Fatalf("Failed to make database connection: %v", err)
	}
	defer dao.DB.Close()

	// Make api controller
	gin.SetMode(gin.ReleaseMode)
	controller := api.Make(dao)

	for _, tc := range tt {
		t.Run(tc.tName, func(t *testing.T) {
			w := performRequest(controller.Router, tc.tMethod, tc.tURL)
			assert.Equal(t, http.StatusOK, w.Code)
		})
	}
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
