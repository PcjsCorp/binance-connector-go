package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/clients/w3wprediction/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	CreateInboundTransfer()
}

func CreateInboundTransfer() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.W3wPredictionRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceW3wPredictionClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.TransferAPI.CreateInboundTransfer(context.Background()).WalletId("5b5c1ec3be4e4416a5872b21c1ca5d20").WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").FromTokenAmount("1000000000000000000").AccountType(models.PlaceOrderAccountTypeParameterSpot).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
