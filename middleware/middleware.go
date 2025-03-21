package middleware

import (
	"github.com/Omaroma/moneytransfer/apis"
	"github.com/gin-gonic/gin"
)

const PORT = ":8086"

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/addUser", apis.AddUser)
	router.DELETE("/deleteUser", apis.DeleteUser)
	router.POST("/transfer", apis.Transfer)
	router.GET("/userInfo", apis.GetUserInfo)
	router.GET("/allUsers", apis.GetAllUsers)

	return router
}
