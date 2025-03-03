package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/it"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/it/types/order"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = ""
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

	options := []it.Option{
		it.WithClientConfig(config),
		it.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}
	client, err := it.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	orderDetail := Order{
		AutoPay: true,
		Orders: []OrderItem{
			{
				CycleCnt:    0,
				CycleType:   101,
				InstanceCnt: 1,
				Items: []Item{
					{
						ItemConfig: map[string]interface{}{
							"version":  "v1",
							"edition":  "basic",
							"billMode": "2",
							"regionId": "b342b77ef26b11ecb0ac0242ac110002",
							"azInfo": []map[string]string{
								{
									"azName": "cn-xinan1-1A",
								},
							},
							"name": "eci-zc11x1tl2g6c46vjrkvepxaiyz",
							"extJson": map[string]interface{}{
								"busiChannel": "010",
								"clusterName": "eci-zc11x1tl2g6c46vjrkvepxaiyz",
								"prodId":      "12710101",
								"attrMap": map[string]string{
									"cpu":               "1",
									"memory":            "2",
									"restartPolicy":     "Always",
									"instancePayAmount": "1",
									"vpcUuid":           "vpc-gbji827p0c",
									"subnetUuid":        "subnet-de0fpe2yab",
									"securityGroupUuid": "sg-lqmyzwighs",
								},
							},
						},
						ItemValue:    2,
						Master:       true,
						ResourceType: "PAAS_ECI",
						ServiceTag:   "PAAS",
					},
					{
						ItemConfig: map[string]interface{}{
							"Version": "v1",
							"Type":    "mem",
						},
						ItemValue:    2,
						Master:       false,
						ResourceType: "PAAS_ECI_BASIC",
						ServiceTag:   "PAAS",
					},
					{
						ItemConfig: map[string]interface{}{
							"Version": "v1",
							"Type":    "cpu",
						},
						ItemValue:    1,
						Master:       false,
						ResourceType: "PAAS_ECI_BASIC",
						ServiceTag:   "PAAS",
					},
				},
			},
		},
		CustomInfo: CustomInfo{
			Email:    "",
			Identity: Identity{AccountId: "8aee44f097114efaaead22acf2ed4194"},
		},
	}

	ctx := context.Background()
	req := &order.PlaceOnDemandNewPurchaseOrderRequest{
		// Fill in the request parameters
		OrderDetailJson: orderDetail.Marshal(),
	}
	resp, raw, err := client.Order().PlaceOnDemandNewPurchaseOrder(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
