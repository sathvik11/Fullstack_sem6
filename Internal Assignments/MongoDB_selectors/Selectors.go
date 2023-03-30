package main

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection URI
const uri = "mongodb://nikx-training-mongo:4ZR3wvYKDbgix7TRvJ8EsaBJMgQKGL165f1m1vLB1hRSDNIEhMGSAHobpHaPbca5IXKHGi5zG7kkACDbnNrowA==@nikx-training-mongo.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@nikx-training-mongo@"

var client, _ = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
var coll = client.Database("mflix").Collection("movies")

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Content-Length"},
	}))

	router.GET("/movies", searchMovies)

	router.Run("localhost:5700")
}

type Movie struct {
	Title string `bson:"title,omitempty"`
	Id    string `bson:"_id,omitempty"`
	Rated string `bson:"rated,omitempty"`
}

type SearchFilter struct {
	Title  string `json:"title,omitempty"`
	Rating int    `json:"rating,omitempty"`
}

func searchMovies(c *gin.Context) {
	var searchFilter SearchFilter
	if err := c.ShouldBindQuery(&searchFilter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{}
	if searchFilter.Title != "" {
		filter["title"] = searchFilter.Title
	}
	if searchFilter.Rating != 0 {
		filter["rated"] = bson.M{"$eq": searchFilter.Rating}
	}

	opts := options.Find().SetLimit(50)
	cursor, err := coll.Find(context.Background(), filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "No documents found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var results []Movie
	if err := cursor.All(context.Background(), &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No documents found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}
