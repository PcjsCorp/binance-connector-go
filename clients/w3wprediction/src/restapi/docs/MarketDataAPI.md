# \MarketDataAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetMarketDetail**](MarketDataAPI.md#GetMarketDetail) | **Get** /sapi/v1/w3w/wallet/prediction/market/detail | Get Market Detail
[**ListPredictionCategories**](MarketDataAPI.md#ListPredictionCategories) | **Get** /sapi/v1/w3w/wallet/prediction/category/list | List Prediction Categories
[**ListPredictionMarkets**](MarketDataAPI.md#ListPredictionMarkets) | **Get** /sapi/v1/w3w/wallet/prediction/market/list | List Prediction Markets
[**MarketSearch**](MarketDataAPI.md#MarketSearch) | **Get** /sapi/v1/w3w/wallet/prediction/market/search | Market Search
[**QueryLastTradePrice**](MarketDataAPI.md#QueryLastTradePrice) | **Get** /sapi/v1/w3w/wallet/prediction/order-book/last-trade-price | Query Last Trade Price
[**QueryOrderBook**](MarketDataAPI.md#QueryOrderBook) | **Get** /sapi/v1/w3w/wallet/prediction/order-book | Query Order Book


## GetMarketDetail

> GetMarketDetailResponse GetMarketDetail(ctx).MarketTopicId(marketTopicId).Execute()

Get Market Detail


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	marketTopicId := int64(4229564) // int64 | Market topic ID. Must be > 0

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetMarketDetail(context.Background()).MarketTopicId(marketTopicId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetMarketDetail``: %v\n", err)
		return
	}

	// response from `GetMarketDetail`: GetMarketDetailResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **marketTopicId** | **int64** | Market topic ID. Must be &gt; 0 | 

### Return type

[**GetMarketDetailResponse**](GetMarketDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ListPredictionCategories

> ListPredictionCategoriesResponse ListPredictionCategories(ctx).Execute()

List Prediction Categories


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.ListPredictionCategories(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.ListPredictionCategories``: %v\n", err)
		return
	}

	// response from `ListPredictionCategories`: ListPredictionCategoriesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**ListPredictionCategoriesResponse**](ListPredictionCategoriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ListPredictionMarkets

> ListPredictionMarketsResponse ListPredictionMarkets(ctx).L1Category(l1Category).L2Category(l2Category).SortBy(sortBy).OrderBy(orderBy).Offset(offset).Limit(limit).Execute()

List Prediction Markets


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	l1Category := "crypto" // string | Level-1 category filter (optional)
	l2Category := "up-down" // string | Level-2 category filter (optional)
	sortBy := models.ListPredictionMarketsSortByParameterRecommended // ListPredictionMarketsSortByParameter | Sort field. Enum: `RECOMMENDED`, `VOLUME`, `PARTICIPANTS`, `CREATED_TIME`, `END_DATE` (optional)
	orderBy := models.ListPredictionMarketsOrderByParameterAsc // ListPredictionMarketsOrderByParameter | Sort direction. Enum: `ASC`, `DESC` (optional)
	offset := int32(0) // int32 | Pagination offset. Default `0` (optional)
	limit := int32(20) // int32 | Page size. Default `20`, range 1–100 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.ListPredictionMarkets(context.Background()).L1Category(l1Category).L2Category(l2Category).SortBy(sortBy).OrderBy(orderBy).Offset(offset).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.ListPredictionMarkets``: %v\n", err)
		return
	}

	// response from `ListPredictionMarkets`: ListPredictionMarketsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **l1Category** | **string** | Level-1 category filter | 
 **l2Category** | **string** | Level-2 category filter | 
 **sortBy** | [**ListPredictionMarketsSortByParameter**](ListPredictionMarketsSortByParameter.md) | Sort field. Enum: &#x60;RECOMMENDED&#x60;, &#x60;VOLUME&#x60;, &#x60;PARTICIPANTS&#x60;, &#x60;CREATED_TIME&#x60;, &#x60;END_DATE&#x60; | 
 **orderBy** | [**ListPredictionMarketsOrderByParameter**](ListPredictionMarketsOrderByParameter.md) | Sort direction. Enum: &#x60;ASC&#x60;, &#x60;DESC&#x60; | 
 **offset** | **int32** | Pagination offset. Default &#x60;0&#x60; | 
 **limit** | **int32** | Page size. Default &#x60;20&#x60;, range 1–100 | 

### Return type

[**ListPredictionMarketsResponse**](ListPredictionMarketsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarketSearch

> MarketSearchResponse MarketSearch(ctx).Query(query).TopK(topK).Execute()

Market Search


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	query := "BTC price" // string | Search keyword. Not blank
	topK := int32(20) // int32 | Max number of results to return. Default `20`, range 1–50 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.MarketSearch(context.Background()).Query(query).TopK(topK).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.MarketSearch``: %v\n", err)
		return
	}

	// response from `MarketSearch`: MarketSearchResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Search keyword. Not blank | 
 **topK** | **int32** | Max number of results to return. Default &#x60;20&#x60;, range 1–50 | 

### Return type

[**MarketSearchResponse**](MarketSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryLastTradePrice

> QueryLastTradePriceResponse QueryLastTradePrice(ctx).MarketId(marketId).Execute()

Query Last Trade Price


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	marketId := int64(5567895) // int64 | Market ID. Must be > 0

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.QueryLastTradePrice(context.Background()).MarketId(marketId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.QueryLastTradePrice``: %v\n", err)
		return
	}

	// response from `QueryLastTradePrice`: QueryLastTradePriceResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **marketId** | **int64** | Market ID. Must be &gt; 0 | 

### Return type

[**QueryLastTradePriceResponse**](QueryLastTradePriceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryOrderBook

> QueryOrderBookResponse QueryOrderBook(ctx).Vendor(vendor).MarketId(marketId).TokenId(tokenId).Execute()

Query Order Book


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	vendor := "predict_fun" // string | Vendor identifier (e.g. `predict_fun`)
	marketId := int64(5567895) // int64 | Market ID. Must be > 0
	tokenId := "112233" // string | Prediction outcome token ID

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.QueryOrderBook(context.Background()).Vendor(vendor).MarketId(marketId).TokenId(tokenId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.QueryOrderBook``: %v\n", err)
		return
	}

	// response from `QueryOrderBook`: QueryOrderBookResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **vendor** | **string** | Vendor identifier (e.g. &#x60;predict_fun&#x60;) | 
 **marketId** | **int64** | Market ID. Must be &gt; 0 | 
 **tokenId** | **string** | Prediction outcome token ID | 

### Return type

[**QueryOrderBookResponse**](QueryOrderBookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

