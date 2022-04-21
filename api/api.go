package api

import (
	"net/http"

	"github.com/eas-kaine/discord-bot/models"
	"github.com/gin-gonic/gin"
)

// List of all quotes as JSON
func getQuotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, quotes)
}

// Locates the quote by id
// Parameter sent by the client, then returns that quote as a response.
func getQuotesByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of quotes, to find the corresponding value
	for _, a := range quotes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "quote not found"})
}

// Adds a quote from JSON recieved in the request body
func postQuotes(c *gin.Context) {
	var newQuote models.Quote

	// Call BindJSON to bind the received JSON to newQuote
	if err := c.BindJSON(&newQuote); err != nil {
		return
	}

	// Add the new quote to the slice
	quotes = append(quotes, newQuote)
	c.IndentedJSON(http.StatusCreated, newQuote)
}

// Quotes slice for quote data
var quotes = []models.Quote{
    {Text: "The greatest glory in living lies not in never falling, but in rising every time we fall.", Author: "Nelson Mandela"},
    {Text: "The way to get started is to quit talking and begin doing.", Author: "Walt Disney"},
    {Text: "Your time is limited, so don't waste it living someone else's life. Don't be trapped by dogma – which is living with the results of other people's thinking.", Author: "Steve Jobs"},
}

// List of all facts as JSON
func getFacts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, facts)
}

// Locates the facts by id
// Parameter sent by the client, then returns that fact as a response.
func getFactsByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of facts, to find the corresponding value
	for _, a := range facts {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "fact not found"})
}

// Adds a fact from JSON recieved in the request body
func postFacts(c *gin.Context) {
	var newFact models.Fact

	// Call BindJSON to bind the received JSON to newFact
	if err := c.BindJSON(&newFact); err != nil {
		return
	}

	// Add the new fact to the slice
	facts = append(facts, newFact)
	c.IndentedJSON(http.StatusCreated, newFact)
}

// Facts slice for fact data
var facts = []models.Fact{
    {Subject: "Food", Text: "Avocados are a fruit, not a vegetable."},
    {Subject: "Travel", Text: "The Eiffel Tower can be 15 cm taller during the summer."},
    {Subject: "Music", Text: "The Spice Girls were originally a band called Touch."},
}

func Run() {
	router := gin.Default()
	router.GET("/quotes", getQuotes)
	router.GET("/quotes/:id", getQuotesByID)
	router.POST("/quotes", postQuotes)

	router.GET("/facts", getFacts)
	router.GET("/facts/:id", getFactsByID)
	router.POST("/facts", postFacts)
	
	router.Run("localhost:3000")
}