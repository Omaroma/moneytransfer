package services

import (
	"errors"
	"github.com/Omaroma/moneytransfer/models"
	"github.com/google/uuid"
)

func GetUser(id string) (user models.User) {
	return models.Users.Get(id)
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
