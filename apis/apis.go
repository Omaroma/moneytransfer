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
		return
	}

	userBalance, _ := strconv.ParseFloat(c.Query("balance"), 64)
	if userBalance < 0 {
		c.String(http.StatusBadRequest, "balance can't be negative")
		return
	}

	user := services.AddUser(userName, userBalance)
	c.JSON(http.StatusCreated, user)
}

func DeleteUser(c *gin.Context) {
	userId := c.Query("id")
	if userId == "" {
		c.String(http.StatusBadRequest, "you must enter user id")
		return
	}

	services.DeleteUser(userId)
	c.String(http.StatusOK, "user deleted successfully")
}

func Transfer(c *gin.Context) {
	senderId := c.Query("from")
	if senderId == "" {
		c.String(http.StatusBadRequest, "you must set sender Id")
		return
	}

	receiverId := c.Query("to")
	if receiverId == "" {
		c.String(http.StatusBadRequest, "you must set receiver Id")
		return
	}

	amount, _ := strconv.ParseFloat(c.Query("amount"), 64)
	if amount == 0 {
		c.String(http.StatusBadRequest, "please enter transfer amount")
		return
	}

	err := services.TransferMoney(senderId, receiverId, amount)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, "amount added successfully")
}

func GetUserInfo(c *gin.Context) {
	userId := c.Query("id")
	if userId == "" {
		c.String(http.StatusBadRequest, "you must enter user id")
		return
	}

	user := services.GetUser(userId)
	c.JSON(http.StatusOK, user)
}
