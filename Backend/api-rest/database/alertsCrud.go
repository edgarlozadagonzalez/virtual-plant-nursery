package database

import (
	"virtual-plant-nursery/api-rest/models"

	"github.com/gin-gonic/gin"
)

type AlertRequestBody struct {
	AlertID  uint   `gorm:"primaryKey" json:"alert_id"`
	Reason   string `gorm:"reason" json:"reason"`
	DateTime string `gorm:"date_time" json:"date_time"`
	PlantID  string `gorm:"plant_id" json:"plant_id"`
}

func AlertCreate(c *gin.Context) {
	body := AlertRequestBody{}
	c.BindJSON(&body)

	alert := &models.Alert{
		//		AlertID:  body.AlertID,
		Reason: body.Reason,
		//	DateTime: body.DateTime,
		//	PlantID:  body.PlantID,
	}

	result := Db.Create(&alert)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to insert"})
		return
	}

	c.JSON(200, &alert)
}

func AlertGet(c *gin.Context) {
	var alerts []models.Alert
	Db.Find(&alerts)
	c.JSON(200, &alerts)
	return
}

func AlertGetByID(c *gin.Context) {
	id := c.Param("id")
	var alert models.Alert
	Db.First(&alert, id)
	c.JSON(200, &alert)
	return
}

func AlertUpdate(c *gin.Context) {
	id := c.Param("id")
	var alert models.Alert
	Db.First(&alert, id)

	body := AlertRequestBody{}
	c.BindJSON(&body)

	data := &models.Alert{
		//		AlertID:  body.AlertID,
		Reason: body.Reason,
		//	DateTime: body.DateTime,
		//	PlantID:  body.PlantID,
	}

	result := Db.Model(&alert).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &alert)
}

func AlertDelete(c *gin.Context) {
	id := c.Param("id")
	var alert models.Alert
	Db.Delete(&alert, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}
