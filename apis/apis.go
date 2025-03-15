package apis

import "github.com/gin-gonic/gin"

func AddUser(c *gin.Context) {
	c.String(200, "user added successfully")
}

func DeleteUser(c *gin.Context) {
	c.String(200, "user deleted successfully")
}

func Transfer(c *gin.Context) {
	c.String(200, "amount added successfully")
}

func GetBalance(c *gin.Context) {
	c.String(200, "balance is: 0")
}
