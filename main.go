package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gobuffalo/packr"
	"github.com/kautsarady/forex/api"
	"github.com/kautsarady/forex/model"

	_ "github.com/lib/pq"
)

func main() {

	cs := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DBNAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	withDummy := true

	// Make database connection
	dao, err := model.Make(cs, withDummy)
	if err != nil {
		log.Fatalf("Failed to make database connection: %v [RETRY]", err)
	}
	defer dao.DB.Close()

	// Make api controller
	controller := api.Make(dao)

	// Serve web page
	controller.Router.StaticFS("/web", packr.NewBox("./public"))

	// Run http server
	addr := ":" + os.Getenv("PORT")
	if err := controller.Router.Run(addr); err != nil {
		log.Fatalf("Failed to run controller.Router: %v", err)
	}
}
