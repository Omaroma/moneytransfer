package services

import (
	"errors"
	"github.com/Omaroma/moneytransfer/models"
	"github.com/google/uuid"
	"log"
)

func GetUser(id string) (user models.User) {
	return models.Users.Get(id)
}

func GetAllUsers() []models.User {
	return models.Users.GetAllData()
}

func AddUser(name string, balance float64) (user models.User) {
	user.Id = uuid.New().String()
	user.Name = name
	user.Balance = balance

	models.Users.Add(user)
	return user
}

func DeleteUser(id string) {
	models.Users.Delete(id)
}

func TransferMoney(from, to string, amount float64) error {
	sender := models.Users.Get(from)
	if sender.Id == "" {
		return errors.New("sender doesn't exist")
	}

	receiver := models.Users.Get(to)
	if receiver.Id == "" {
		return errors.New("receiver doesn't exist")
	}

	if sender.Balance < amount {
		return errors.New("insufficient funds")
	}

	newSenderBalance := sender.Balance - amount
	newReceiverBalance := receiver.Balance + amount
	models.Users.UpdateBalance(sender.Id, newSenderBalance)
	models.Users.UpdateBalance(receiver.Id, newReceiverBalance)

	return nil
}

func InitUsers() {
	models.InitUserStore()
	userMark := AddUser("Mark", 100)
	log.Println("created user:", userMark)
	userJane := AddUser("Jane", 50)
	log.Println("created user:", userJane)
	userAdam := AddUser("Adam", 0)
	log.Println("created user:", userAdam)
}
