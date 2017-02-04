package handler

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/rakeshkumargupt/models"
	"github.com/rakeshkumargupt/webpush"
)


//Function to call send web push
//Input is json
func SendWebpush(c *gin.Context)  {

	//Getting input data after binding json to UiWebpushModel struct
	inputWebpush := c.Keys["form_data"].(*models.UiWebpushModel)
	err := webpush.WebPush(*inputWebpush)
	if err != nil {
		c.JSON(501, gin.H{"Message": "Something wrong with the form parameters"})
	}
	c.JSON(201, gin.H{"Message": "Webpush Success!!"})
	return
}
