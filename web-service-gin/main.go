package main

import(
	
	"net/http"


	"github.com/gin-gonic/gin"
)

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var albums = []album {
	{ID: "1", Title: "Blue", Artist: "John", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughn", Artist: "Sarah", Price: 39.99},
	{ID: "4", Title: "Guff", Artist: "Melquiades", Price: 9.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}