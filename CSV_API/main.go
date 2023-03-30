package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Data struct {
	ID       int    `csv:"id"`
	Name     string `csv:"name"`
	Age      string `csv:"age"`
	Location string `csv:"location"`
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/data", fileExists(), getData)
	router.GET("/data/:id", fileExists(), getOneData)
	router.POST("/data", fileExists(), createData)
	router.PUT("/data/:id", fileExists(), updateData)
	router.DELETE("/data/:id", fileExists(), deleteData)

	if err := router.Run(":5700"); err != nil {
		log.Fatal(err)
	}
}

// Check if the CSV file exists
func fileExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := os.Stat("user.csv"); os.IsNotExist(err) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "CSV file not found"})
			return
		}
		c.Next()
	}
}

// Get data from the CSV file using GET method
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
		if i == 0 {
			continue
		}

		id, err := strconv.Atoi(row[0])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		data := Data{
			ID:       id,
			Name:     row[1],
			Age:      row[2],
			Location: row[3],
		}

		result = append(result, data)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Getting data successful!", "data": result})
}

// Get filtered data from the CSV file using GET method
func getOneData(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
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

	var result Data
	for i, row := range records {
		if i == 0 {
			continue
		}

		recordID, err := strconv.Atoi(row[0])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if recordID == id {
			result = Data{
				ID:       recordID,
				Name:     row[1],
				Age:      row[2],
				Location: row[3],
			}
		}
		// Get all data from the CSV file using GET method
func getAllData(c *gin.Context) {
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
		if i == 0 {
			continue
		}

		id, _ := strconv.Atoi(row[0])
		age, _ := strconv.Atoi(row[2])

		result = append(result, Data{
			ID:       id,
			Name:     row[1],
			Age:      age,
			Location: row[3],
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Getting data successful!", "data": result})
}

// Create new data in the CSV file using POST method
func createData(c *gin.Context) {
	var data Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if CSV file exists
	if _, err := os.Stat("user.csv"); os.IsNotExist(err) {
		file, err := os.Create("user.csv")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		header := []string{"id", "name", "age", "location"}
		if err := writer.Write(header); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	file, err := os.OpenFile("user.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		strconv.Itoa(data.ID),
		data.Name,
		strconv.Itoa(data.Age),
		data.Location,
	}

	if err := writer.Write(record); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data created successfully!", "data": data})
}

// Update data in the CSV file using PUT method
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

	file, err := os.OpenFile("user.csv", os.O_RDWR, 0644)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	records, err := reader.ReadAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
		}
		var updatedRecords [][]string
for _, record := range records {
	if record[0] == strconv.Itoa(id) {
		record[1] = data.Name
		record[2] = data.Age
		record[3] = data.Email
		record[4] = data.Address
		record[5] = data.PhoneNumber
	}
	updatedRecords = append(updatedRecords, record)
}

file.Truncate(0)
file.Seek(0, 0)

for _, updatedRecord := range updatedRecords {
	if err := writer.Write(updatedRecord); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}
