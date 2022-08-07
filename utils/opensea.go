package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"element.com/m/models"
	"github.com/ethereum/go-ethereum/common"
)

type OpenseaAssetsReponse struct {
	Assets []OpenseaAsset `json:"assets"`
}

type OpenseaAsset struct {
	TokenId string `json:"token_id"`
	Collection OpenseaCollection `json:"asset_contract"`
}

type OpenseaCollection struct {
	Address string `json:"address"`
}

func FetchUserNfts(address string) []models.NFT {
	// build params
	params := url.Values{}
	params.Add("owner", address)


	// make request
	resp, err := http.Get(os.Getenv("OPENSEA_API_URL") + "assets?" + params.Encode())
	if err != nil {
		fmt.Println("Error http request")
		return []models.NFT{}
	}

	// get response body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error read response")
		return []models.NFT{}
	}
	
	// unmarshal response
	var openseaResponse OpenseaAssetsReponse
	err = json.Unmarshal(body, &openseaResponse)
	if err != nil {
		fmt.Println("Error unmarshal")
		return[]models.NFT{}
	}

	// create slice of NFTs and return
	var nfts []models.NFT
	for _, asset := range openseaResponse.Assets {
		nft := models.NFT{Collection: common.HexToAddress(asset.Collection.Address), TokenID: asset.TokenId }
		nfts = append(nfts, nft)
	}
	return nfts
}
