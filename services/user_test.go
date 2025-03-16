package services

import (
	"errors"
	"github.com/Omaroma/moneytransfer/models"
	"testing"
)

func TestAddUser(t *testing.T) {
	models.InitUserStore()
	wantName := "Omar"
	wantBalance := float64(100)

	got := AddUser("Omar", 100)
	if got.Id == "" {
		t.Errorf("id wan't generated")
	}

	if got.Name != wantName && got.Balance != wantBalance {
		t.Errorf("got name %v, balance %v want name %v, balance %v", got.Name, got.Balance, wantName, wantBalance)
	}
}

func TestTransferMoney(t *testing.T) {
	models.InitUserStore()

	sender := AddUser("Mark", 100)
	receiver := AddUser("Jane", 50)

	var tests = []struct {
		caseName string
		from, to string
		amount   float64
		wantErr  error
	}{
		{"Sender is not registered",
			"1c60a2ae-8672-4843-9164-f1b82ca068f5", receiver.Id, 100, errors.New("sender doesn't exist")},

		{"Receiver is not registered",
			sender.Id, "1c60a2ae-8672-4843-9164-f1b82ca068f5", 100, errors.New("receiver doesn't exist")},

		{"Not enough cash to do transaction",
			sender.Id, receiver.Id, 700, errors.New("insufficient funds")},

		{"Successful transaction",
			sender.Id, receiver.Id, 60, nil},
	}

	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			gotErr := TransferMoney(tt.from, tt.to, tt.amount)
			if gotErr != nil && tt.wantErr != nil {
				if errors.Is(gotErr, tt.wantErr) {
					t.Errorf("got error %v, want %v", gotErr, tt.wantErr)
				}
			}
		})
	}
}
