/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package binancesimpleearnrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/simpleearn/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// YieldArenaAPIService YieldArenaAPI Service
type YieldArenaAPIService Service

type ApiGetYieldArenaActivitiesRequest struct {
	ctx        context.Context
	ApiService *YieldArenaAPIService
	recvWindow *int64
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetYieldArenaActivitiesRequest) RecvWindow(recvWindow int64) ApiGetYieldArenaActivitiesRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetYieldArenaActivitiesRequest) Execute() (*common.RestApiResponse[models.GetYieldArenaActivitiesResponse], error) {
	return r.ApiService.GetYieldArenaActivitiesExecute(r)
}

/*
GetYieldArenaActivities Get Yield Arena Activities (USER_DATA)
Get /sapi/v1/earn/arena/activities

https://developers.binance.com/docs/simple_earn/yield-arena/earn/Get-Yield-Arena-Activities

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetYieldArenaActivitiesRequest
*/
func (a *YieldArenaAPIService) GetYieldArenaActivities(ctx context.Context) ApiGetYieldArenaActivitiesRequest {
	return ApiGetYieldArenaActivitiesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetYieldArenaActivitiesResponse
func (a *YieldArenaAPIService) GetYieldArenaActivitiesExecute(r ApiGetYieldArenaActivitiesRequest) (*common.RestApiResponse[models.GetYieldArenaActivitiesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/earn/arena/activities"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetYieldArenaActivitiesResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
