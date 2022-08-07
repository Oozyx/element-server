package models

import "github.com/ethereum/go-ethereum/common"

type User struct {
	ID uint `json:"id" gorm:"primary_key"`
	Address common.Address `json:"address" gorm:"unique"` 
	NFTs []NFT `json:"nfts"`
}

type NFT struct {
	Collection common.Address `json:"collection"`
	TokenID string `json:"TokenID"`
}