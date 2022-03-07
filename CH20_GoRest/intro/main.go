package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{
		"1", "Blue Train", "John Coltrane", 56.99,
	},
	{
		"2", "Jeru", "Gerry Mulligan", 17.99,
	},
	{
		"3", "Sarah Vaughan and Clifford Brown", "Sarah Vaughan", 39.99,
	},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumById) // colon preceeding an item in the path -- means is path variable

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) // serialize the albums (slice) into JSON -- pretty print
	// c.JSON(http.StatusOK, albums)      // serialize the albums (slice) into JSON -- minimize
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil { // Bind request body to newAlbum
		c.IndentedJSON(http.StatusInternalServerError, struct {
			Message string `json:"message"`
		}{"Fatal error can't bind the request body into model data"})
	} else {
		albums = append(albums, newAlbum)
		c.IndentedJSON(http.StatusCreated, newAlbum)
	}
}

type person struct {
	Name string `form:"name"`
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	var newPerson person

	if err := c.ShouldBind(&newPerson); err == nil {
		log.Println(newPerson.Name)
	}

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, struct {
		Message string `json:"message"`
	}{"Album is Not Found!"})
}
