// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Telecom Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package eci

import (
	"context"
	"fmt"
	"net/http"

	"github.com/telecom-cloud/client-go/pkg/common/config"
	"github.com/telecom-cloud/client-go/pkg/openapi"
	"github.com/telecom-cloud/client-go/pkg/protocol"

	price "github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/price"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
)

type PriceClient interface {
	DescribePrice(context context.Context, req *price.DescribeContainerGroupPriceRequest, reqOpt ...config.RequestOption) (resp *price.DescribeContainerGroupPriceResponse, rawResponse *protocol.Response, err error)

	BatchDescribePrice(context context.Context, req *price.BatchDescribeContainerGroupPriceRequest, reqOpt ...config.RequestOption) (resp *price.BatchDescribeContainerGroupPriceResponse, rawResponse *protocol.Response, err error)
}

type priceClient struct {
	client *HttpClient
}

func NewPriceClient(hostUrl string, ops ...Option) (PriceClient, error) {
	opts := GetOptions(append(ops, WithHostUrl(hostUrl))...)
	cli, err := NewHttpClient(opts)
	if err != nil {
		return nil, err
	}
	return &priceClient{
		client: cli,
	}, nil
}

func (s *priceClient) DescribePrice(ctx context.Context, req *price.DescribeContainerGroupPriceRequest, reqOpt ...config.RequestOption) (resp *price.DescribeContainerGroupPriceResponse, rawResponse *protocol.Response, err error) {
	openapiResp := &openapi.OpenapiResponse{}
	openapiResp.ReturnObj = &resp
	ret, err := s.client.R().
		SetContext(ctx).
		AddHeaders(map[string]string{
			"regionId": req.GetRegionId(),
		}).
		SetBodyParam(req).
		SetRequestOption(reqOpt...).
		SetResult(openapiResp).
		Execute(http.MethodPost, "/eci/api/v1/containers/describeContainerGroupPrice")
	if err != nil {
		return nil, nil, err
	}

	rawResponse = ret.RawResponse
	return resp, rawResponse, nil
}

func (s *priceClient) BatchDescribePrice(ctx context.Context, req *price.BatchDescribeContainerGroupPriceRequest, reqOpt ...config.RequestOption) (resp *price.BatchDescribeContainerGroupPriceResponse, rawResponse *protocol.Response, err error) {
	openapiResp := &openapi.OpenapiResponse{}
	openapiResp.ReturnObj = &resp
	ret, err := s.client.R().
		SetContext(ctx).
		AddHeaders(map[string]string{
			"regionId": req.GetRegionId(),
		}).
		SetBodyParam(req).
		SetRequestOption(reqOpt...).
		SetResult(openapiResp).
		Execute(http.MethodPost, "/eci/api/v1/containers/batchDescribeContainerGroupPrice")
	if err != nil {
		return nil, nil, err
	}

	rawResponse = ret.RawResponse
	return resp, rawResponse, nil
}

var defaultPriceClient, _ = NewPriceClient(baseDomain)

func ConfigDefaultPriceClient(ops ...Option) (err error) {
	defaultPriceClient, err = NewPriceClient(baseDomain, ops...)
	return
}

func DescribePrice(context context.Context, req *price.DescribeContainerGroupPriceRequest, reqOpt ...config.RequestOption) (resp *price.DescribeContainerGroupPriceResponse, rawResponse *protocol.Response, err error) {
	return defaultPriceClient.DescribePrice(context, req, reqOpt...)
}

func BatchDescribePrice(context context.Context, req *price.BatchDescribeContainerGroupPriceRequest, reqOpt ...config.RequestOption) (resp *price.BatchDescribeContainerGroupPriceResponse, rawResponse *protocol.Response, err error) {
	return defaultPriceClient.BatchDescribePrice(context, req, reqOpt...)
}
