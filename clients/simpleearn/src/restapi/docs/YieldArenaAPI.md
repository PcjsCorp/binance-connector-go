# \YieldArenaAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetYieldArenaActivities**](YieldArenaAPI.md#GetYieldArenaActivities) | **Get** /sapi/v1/earn/arena/activities | Get Yield Arena Activities (USER_DATA)


## GetYieldArenaActivities

> GetYieldArenaActivitiesResponse GetYieldArenaActivities(ctx).RecvWindow(recvWindow).Execute()

Get Yield Arena Activities (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.YieldArenaAPI.GetYieldArenaActivities(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `YieldArenaAPI.GetYieldArenaActivities``: %v\n", err)
		return
	}

	// response from `GetYieldArenaActivities`: GetYieldArenaActivitiesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetYieldArenaActivitiesResponse**](GetYieldArenaActivitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

