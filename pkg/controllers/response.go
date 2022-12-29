package controllers

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	//logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})

}
