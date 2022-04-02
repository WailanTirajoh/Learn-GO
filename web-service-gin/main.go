package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{
		ID:     "1",
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	},
	{
		ID:     "2",
		Title:  "Jeru",
		Artist: "Gerry Mulligan",
		Price:  17.99,
	},
	{
		ID:     "3",
		Title:  "Sarah Vaughan and Clifford Brown",
		Artist: "Sarah Vaughan",
		Price:  39.99,
	},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Error"})
		return
	}

	for _, album := range albums {
		if album.ID == newAlbum.ID {
			c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Data already exists for id " + newAlbum.ID})
			return
		}
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("album")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func destroyAlbumByID(c *gin.Context) {
	id := c.Param("album")

	index := SliceIndex(len(albums), func(i int) bool { return albums[i].ID == id })

	if index < 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album for id " + id + " not found"})
		return
	}

	albums = removeSliceByIndex(albums, index)
	c.IndentedJSON(http.StatusOK, albums)
}

// Utils
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func removeSliceByIndex(slice []album, s int) []album {
	return append(slice[:s], slice[s+1:]...)
}

// End Utils

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:album", getAlbumByID)
	router.DELETE("/albums/:album", destroyAlbumByID)

	router.Run("localhost:8080")
}
