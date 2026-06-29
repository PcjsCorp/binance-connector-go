package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	GetRedeemStatus()
}

func GetRedeemStatus() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.W3wPredictionRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceW3wPredictionClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.RedeemAPI.GetRedeemStatus(context.Background()).WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").TxHash("0xabc123def456789abcdef123456789abcdef123456789abcdef123456789abcd").Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
