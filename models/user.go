package models

import "github.com/ethereum/go-ethereum/common"



type User struct {
	ID uint `json:"id" gorm:"primary_key"`
	Address common.Address `json:"address" gorm:"unique"` 
	NumOwned uint `json:"num_owned"`
}