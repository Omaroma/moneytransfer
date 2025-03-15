package apis

import (
	"github.com/Omaroma/moneytransfer/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddUser(c *gin.Context) {
	userName := c.Query("name")
	if userName == "" {
		c.String(http.StatusBadRequest, "you must set user name")
	}

	userBalance, _ := strconv.ParseFloat(c.Query("balance"), 64)
	if userBalance < 0 {
		c.String(http.StatusBadRequest, "balance can't be negative")
		return
	}

	services.AddUser(userName, userBalance)
	c.String(http.StatusCreated, "user added successfully")
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
