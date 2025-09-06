package handler

import (
	"cdr/module/core/model"
	"cdr/module/core/queue"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func RegisterDB(db *gorm.DB) {
	DB = db
}

func APICallHandler(r *gin.Engine) {
	group := r.Group("/cdr")
	{
		group.POST("/import", ImportCDR)
	}
}

func ImportCDR(c *gin.Context) {
	var cdr model.CDR
	if err := c.ShouldBindJSON(&cdr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Publish to RabbitMQ queue for asynchronous processing
	if queue.R == nil {
		// fallback: insert directly if queue not configured
		if err := DB.Where("call_id = ?", cdr.CallID).Assign(cdr).FirstOrCreate(&cdr).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "CDR inserted directly (no queue)"})
		return
	}

	if err := queue.R.PublishCDR(c.Request.Context(), cdr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "CDR queued for processing"})
}
