package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection URI
const uri = "mongodb://nikx-training-mongo:4ZR3wvYKDbgix7TRvJ8EsaBJMgQKGL165f1m1vLB1hRSDNIEhMGSAHobpHaPbca5IXKHGi5zG7kkACDbnNrowA==@nikx-training-mongo.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@nikx-training-mongo@"

var client, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
var coll = client.Database("mflix").Collection("movies")

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/movie", lessThan)
	router.GET("/movie", lessThaneqto)
	router.GET("/movie", equalTo)
	router.GET("/movie", greaterThan)
	router.GET("/movie", greaterThaneqto)
	router.GET("/movie", andSelec)
	router.GET("/movie", notSelec)

	router.Run("localhost:5700")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Next()
	}
}

type Movie struct {
	Title string `bson:"title,omitempty"`
	Id    string `bson:"_id,omitempty"`
	Rated string `bson:"rated,omitempty"`
}

//Displaying movies with rating less than 3
func lessThan(c *gin.Context) {
	size, _ := strconv.ParseInt(c.Query("size"), 10, 64)
	title := c.Query("title")
	opts := options.Find().SetLimit(size)
	filter := bson.M{"Rated": bson.M{"$lt": 3}}
	if title != "" {
		filter = bson.M{"title": title}
	}
	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			c.JSON(http.StatusNoContent, gin.H{"message": "No Content"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	var results []Movie
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	res := map[string]interface{}{"data": results}
	c.JSON(http.StatusOK, gin.H{"message": "Displaying data  with less than 3 ratings!", "Data": res})
}

//Displaying movies with rating less than or equal to 5
func lessThaneqto(c *gin.Context) {
	size, _ := strconv.ParseInt(c.Query("size"), 10, 64)
	title := c.Query("title")
	opts := options.Find().SetLimit(size)
	filter := bson.M{"Rated": bson.M{"$lte": 5}}
	if title != "" {
		filter = bson.M{"title": title}
	}
	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			c.JSON(http.StatusNoContent, gin.H{"message": "No Content"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	var results []Movie
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	res := map[string]interface{}{"data": results}
	c.JSON(http.StatusOK, gin.H{"message": "Displaying data with less than and equal to 5 ratings !", "Data": res})
}

//Displaying movies with rating equal to 10
func equalTo(c *gin.Context) {
	size, _ := strconv.ParseInt(c.Query("size"), 10, 64)
	title := c.Query("title")
	opts := options.Find().SetLimit(size)
	filter := bson.M{"Rated": bson.M{"$eq": 10}}
	if title != "" {
		filter = bson.M{"title": title}
	}
	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			c.JSON(http.StatusNoContent, gin.H{"message": "No Content"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	var results []Movie
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	res := map[string]interface{}{"data": results}
	c.JSON(http.StatusOK, gin.H{"message": "Displaying data with 10 ratings !", "Data": res})
}

//Displaying movies with rating greater than 5
func greaterThan(c *gin.Context) {
	size, _ := strconv.ParseInt(c.Query("size"), 10, 64)
	title := c.Query("title")
	opts := options.Find().SetLimit(size)
	filter := bson.M{"Rated": bson.M{"$gt": 5}}
	if title != "" {
		filter = bson.M{"title": title}
	}
	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			c.JSON(http.StatusNoContent, gin.H{"message": "No Content"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	var results []Movie
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	res := map[string]interface{}{"data": results}
	c.JSON(http.StatusOK, gin.H{"message": "Displaying data with greater than 5 ratings !", "Data": res})
}

//Displaying movies with rating greater than or equal to 7
func greaterThaneqto(c *gin.Context) {
	size, _ := strconv.ParseInt(c.Query("size"), 10, 64)
	title := c.Query("title")
	opts := options.Find().SetLimit(size)
	filter := bson.M{"Rated": bson.M{"$gte": 7}}
	if title != "" {
		filter = bson.M{"title": title}
	}
	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			c.JSON(http.StatusNoContent, gin.H{"message": "No Content"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	var results []Movie
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	res := map[string]interface{}{"data": results}
	c.JSON(http.StatusOK, gin.H{"message": "Displaying data with greater than and equal to 7 rating !", "Data": res})
}

//Displaying movies with  title "Minions" and a rating greater than 5
func andSelec(c *gin.Context) {
	size, _ := strconv.ParseInt(c.Query("size"), 10, 64)
	title := c.Query("title")
	opts := options.Find().SetLimit(size)
	filter := bson.M{ "$and": bson.A{bson.M{"Title": "Minions"},bson.M{"Rated": bson.M{"$gt": 5}},},}
	if title != "" {
		filter = bson.M{"title": title}
	}
	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			c.JSON(http.StatusNoContent, gin.H{"message": "No Content"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	var results []Movie
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	res := map[string]interface{}{"data": results}
	c.JSON(http.StatusOK, gin.H{"message": "Displaying data with title of the movie as 'minions' and with ratings greater than 5!", "Data": res})
}

//Displaying movies which do NOT have the title "Sabahat"
func notSelec(c *gin.Context) {
	size, _ := strconv.ParseInt(c.Query("size"), 10, 64)
	title := c.Query("title")
	opts := options.Find().SetLimit(size)
	filter := bson.M{ "$not": bson.A{bson.M{"Title": "Sabahat"},},}
	if title != "" {
		filter = bson.M{"title": title}
	}
	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			c.JSON(http.StatusNoContent, gin.H{"message": "No Content"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	var results []Movie
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	res := map[string]interface{}{"data": results}
	c.JSON(http.StatusOK, gin.H{"message": "Displaying data which do not have the title 'Sabahat'!", "Data": res})
}
