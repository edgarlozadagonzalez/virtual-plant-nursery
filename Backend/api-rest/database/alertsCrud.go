package database

import (

	"github.com/gin-gonic/gin"
	"virtual-plant-nursery/api-rest/models"
	
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type AlertRequestBody struct {
//	AlertID  string `json:"alert_id"`
	Reason   string `json:"reason"`
	//DateTime string `json:"date_time"`
	PlantID  string `json:"plant_id"`
}

func AlertCreate(c *gin.Context) {
	body := AlertRequestBody{}
	c.BindJSON(&body)

	alert := &models.Alert{
//		AlertID:  body.AlertID,
		Reason:   body.Reason,
	//	DateTime: body.DateTime,
	//	PlantID:  body.PlantID,
	}


	result := db.Create(&alert)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to insert"})
		return
	}

	c.JSON(200, &alert)
}

func AlertGet(c *gin.Context) {
	var alerts []models.Alert
	db.Find(&alerts)
	c.JSON(200, &alerts)
	return
}

func AlertGetByID(c *gin.Context) {
	id := c.Param("id")
	var alert models.Alert
	db.First(&alert, id)
	c.JSON(200, &alert)
	return
}

func AlertUpdate(c *gin.Context) {
	id := c.Param("id")
	var alert models.Alert
	db.First(&alert, id)

	body := AlertRequestBody{}
	c.BindJSON(&body)

	data := &models.Alert{
//		AlertID:  body.AlertID,
		Reason:   body.Reason,
	//	DateTime: body.DateTime,
	//	PlantID:  body.PlantID,
	}

	result := db.Model(&alert).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &alert)
}

func AlertDelete(c *gin.Context) {
	id := c.Param("id")
	var alert models.Alert
	db.Delete(&alert, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}