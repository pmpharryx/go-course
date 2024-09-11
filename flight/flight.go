package flight

import (
	"database/sql"
	"net/http"
	"strconv"

	"log"

	"github.com/gin-gonic/gin"
)

type Flight struct {
	ID          int    `json:"id"`
	Number      int    `json:"number"`
	AirlineCode string `json:"airline_code"`
	Destination string `json:"destination"`
	Arrival     string `json:"arrival"`
}

type flightHandler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) flightHandler {
	return flightHandler{db: db}
}

func (f flightHandler) Create(c *gin.Context) {
	var flight Flight

	err := c.ShouldBindJSON(&flight)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	stmt, err := f.db.Prepare("INSERT INTO flights (number, airline_code, destination, arrival) values ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = stmt.QueryRow(
		flight.Number,
		flight.AirlineCode,
		flight.Destination,
		flight.Arrival,
	).Scan(&flight.ID)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, flight)
}

func (f flightHandler) GetAll(c *gin.Context) {
	stmt, err := f.db.Prepare("SELECT id, number, airline_code, destination, arrival FROM flights")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var flights = []Flight{}

	for rows.Next() {
		var flight Flight

		err := rows.Scan(&flight.ID,
			&flight.Number,
			&flight.AirlineCode,
			&flight.Destination,
			&flight.Arrival,
		)
		if err != nil {
			log.Print("can't Scan row into variable", err)
		}

		flights = append(flights, flight)
	}

	c.JSON(http.StatusOK, flights)
}

func (f flightHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	stmt, err := f.db.Prepare("SELECT id, number, airline_code, destination, arrival FROM flights WHERE id=$1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var flight Flight
	err = stmt.QueryRow(id).Scan(&flight.ID,
		&flight.Number,
		&flight.AirlineCode,
		&flight.Destination,
		&flight.Arrival,
	)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, flight)
}

func (f flightHandler) UpdateById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var flight Flight
	err = c.ShouldBindJSON(&flight)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	flight.ID = id

	stmt, err := f.db.Prepare("UPDATE flights SET number=$1, airline_code=$2, destination=$3, arrival=$4 WHERE id=$5")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = stmt.Exec(flight.Number, flight.AirlineCode, flight.Destination, flight.Arrival, flight.ID)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, flight)
}

func (f flightHandler) DeleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	stmt, err := f.db.Prepare("DELETE FROM flights WHERE id=$1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
