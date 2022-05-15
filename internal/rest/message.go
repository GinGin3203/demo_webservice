package rest

import "github.com/gin-gonic/gin"

func msgErr(err error) gin.H {
	return gin.H{"message": err.Error()}
}

var msgOK = gin.H{"message": "OK"}
