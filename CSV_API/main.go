package main

import (
	"encoding/csv"
	"log"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type Data struct {
	ID    int    `csv:"id"`
	name  string `csv:"name"`
	age  string  `csv:"age"`
	loc  string `csv:"location"`
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/data", getData)
	router.GET("/data", getOneData)
	router.POST("/data", createData)
	router.PUT("/data/:id", updateData)
	router.DELETE("/data/:id", deleteData)

	router.Run("localhost:5700")
}

//Checking if file exists
func fileExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := os.Stat("user.csv")
		if os.IsNotExist(err) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "CSV file not found"})
			return
		}
		c.Next()
	}
}


//Getting data from CSV file using GET method

func getData(c *gin.Context) {
	file, err := os.Open("user.csv")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	var result []Data
	for i, row := range records {
		// Skipping the header row as it would be column names
		if i == 0 {
			continue
		}

		id, err := strconv.Atoi(row[0])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Creating a new Data object
		data := Data{
			ID:    id,
			Name:  row[1],
			Age: row[2],
			Location: row[3],
		}

		result = append(result, data)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Getting data, Successfu!", "Data": result})
}


//Getting filtered data from CSV file using GET method

func getOneData(c *gin.Context) {
	file, err := os.Open("user.csv")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	params := c.Request.URL.Query()

	var filteredRecords [][]string
	for _, row := range records {
		if row[0] == "id" {
			continue
		}

		match := true
		for column, value := range params {
			if column >= len(row) {
				continue
			}

			if row[column] != value[0] {
				match = false
				break
			}
		}

		if match {
			filteredRecords = append(filteredRecords, row)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Records filtered successfully!", "data": filteredRecords})
}


//create new data in the CSV file using POST method

func createData(c *gin.Context) {
	var data Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := os.OpenFile("user.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{
		strconv.Itoa(data.ID),
		data.Name,
		data.age,
		data.loc,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data created successfully","Data": map[string]interface{}{"data": result}})
}


//updating data in the CSV file using PUT method

func updateData(c *gin.Context) {
idParam := c.Param("id")
id, err := strconv.Atoi(idParam)
if err != nil {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
	return
}

var data Data
if err := c.ShouldBindJSON(&data); err != nil {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}

file, err := os.Open("user.csv")
if err != nil {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}
defer file.Close()

reader := csv.NewReader(file)
records, err := reader.ReadAll()
if err != nil {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}

var found bool
for i, row := range records {
	// Skipping the header row as it is column names
	if i == 0 {
		continue
	}

	rowID, err := strconv.Atoi(row[0])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowID == id {
		records[i][1] = data.Name
		records[i][2] = data.Age
		records[i][3] = data.Location
		found = true
		break
	}
}

if !found {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found"})
	return
}

file, err = os.OpenFile("user.csv", os.O_WRONLY|os.O_TRUNC, 0644)
if err != nil {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}
defer file.Close()

writer := csv.NewWriter(file)
defer writer.Flush()

err = writer.WriteAll(records)
if err != nil {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}

c.JSON(http.StatusOK, gin.H{"message": "Records updated successfully"})
}


//deleting data from the CSV file using DELETE Method

func deleteData(c *gin.Context) {
idParam := c.Param("id")

id, err := strconv.Atoi(idParam)
if err != nil {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
	return
}

file, err := os.Open("user.csv")
if err != nil {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error})
defer file.Close()

reader := csv.NewReader(file)
records, err := reader.ReadAll()
if err != nil {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}

var found bool
for i, row := range records {
	// since the header row is column names
	if i == 0 {
		continue
	}

	rowID, err := strconv.Atoi(row[0])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowID == id {
		records = append(records[:i], records[i+1:]...)
		found = true
		break
	}
}

if !found {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found"})
	return
}

file, err = os.OpenFile("data.csv", os.O_WRONLY|os.O_TRUNC, 0644)
if err != nil {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}
defer file.Close()

writer := csv.NewWriter(file)
defer writer.Flush()

err = writer.WriteAll(records)
if err != nil {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}

c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}
