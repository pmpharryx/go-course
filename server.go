package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pmpharryx/go-course/flight"
	_ "github.com/lib/pq"
)

func initDatabase() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	
	log.Print("Connected to database successfully")
	return db
}

func main() {
	db := initDatabase()
	defer db.Close()

	r := gin.Default()

	flightHandler := flight.NewHandler(db)

	r.POST("/flights", flightHandler.Create)
	r.GET("/flights", flightHandler.GetAll)
	r.GET("/flights/:id", flightHandler.GetById)
	r.PUT("/flights/:id", flightHandler.UpdateById)
	r.DELETE("/flights/:id", flightHandler.DeleteById)

	r.Run()
}
