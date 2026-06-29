package w3wprediction

import (
	BinanceW3wPredictionRestApi "github.com/binance/binance-connector-go/clients/w3wprediction/src/restapi"
	"github.com/binance/binance-connector-go/common/v2/common"
)

type BinanceW3wPredictionClient struct {
	RestApi *BinanceW3wPredictionRestApi.RestAPIClient
}

type Option func(*BinanceW3wPredictionClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceW3wPredictionClient) {
		c.RestApi = BinanceW3wPredictionRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceW3wPredictionClient(opts ...Option) *BinanceW3wPredictionClient {
	client := &BinanceW3wPredictionClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
