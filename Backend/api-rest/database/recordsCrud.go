package database

import (

	"github.com/gin-gonic/gin"
	"virtual-plant-nursery/api-rest/models"
	
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type RecordRequestBody struct {
	//RecordID string `json:"record_id"`
	Reason   string `json:"reason"`
//	PlantID  string `json:"plant_id"`
}

func RecordCreate(c *gin.Context) {
	body := RecordRequestBody{}
	c.BindJSON(&body)

	record := &models.Record{
	//	RecordID: body.RecordID,
		Reason:   body.Reason,
	//	PlantID:  body.PlantID,
	}

	result := db.Create(&record)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &record)
}

func RecordGet(c *gin.Context) {
	var records []models.Record
	db.Find(&records)
	c.JSON(200, &records)
	return
}

func RecordGetById(c *gin.Context) {
	id := c.Param("id")
	var record models.Record
	db.First(&record, id)
	c.JSON(200, &record)
	return
}

func RecordUpdate(c *gin.Context) {
	id := c.Param("id")
	var record models.Record
	db.First(&record, id)

	body := RecordRequestBody{}
	c.BindJSON(&body)

	data := &models.Record{
	//	RecordID: body.RecordID,
		Reason:   body.Reason,
	//	PlantID:  body.PlantID,
	}

	result := db.Model(&record).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &record)
}

func RecordDelete(c *gin.Context) {
	id := c.Param("id")
	var record models.Record
	db.Delete(&record, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}