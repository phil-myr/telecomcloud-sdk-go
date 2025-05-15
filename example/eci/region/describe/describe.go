package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/region"
)

var (
	accessKey  = "b1accda96cb74be390009d2144923466"
	secretKey  = "64caa9d195f14f5da3e973f7d076b6e3"
	baseDomain = "https://eci-global.ctapi-test.ctyun.cn:21443"
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_ECI_DOMAIN")
	if domain != "" {
		baseDomain = domain
	}
}

func main() {
	config := &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []eci.Option{
		eci.WithClientConfig(config),
	}
	client, err := eci.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	req := &region.DescribeRegionRequest{
		RegionId: "b342b77ef26b11ecb0ac0242ac110002",
	}
	resp, raw, err := client.Region().DescribeRegion(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
