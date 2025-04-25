package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGatewayListTrace.json
func ExampleGatewayClient_ListTrace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewayClient().ListTrace(ctx, "rg1", "apimService1", "gw1", armapimanagement.GatewayListTraceContract{
		TraceID: to.Ptr("CrDvXXXXXXXXXXXXXVU3ZA2-1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Value = map[string]any{
	// 	"serviceName": "apimService1",
	// 	"traceEntries": map[string]any{
	// 		"inbound":[]any{
	// 			map[string]any{
	// 				"data":map[string]any{
	// 					"request":map[string]any{
	// 						"method": "GET",
	// 						"headers":[]any{
	// 							map[string]any{
	// 								"name": "Host",
	// 								"value": "proxy.msitesting.net",
	// 							},
	// 						},
	// 						"url": "https://proxy.msitesting.net/6452XXXXXXXXXXXX9c2facb1/64524dXXXXXXXXXXXX2facb3?subscription-key=117313e70XXXXXXXXXXXX38ba4658ca3",
	// 					},
	// 				},
	// 				"elapsed": "00:00:00.2983315",
	// 				"source": "api-inspector",
	// 				"timestamp": "2023-05-03T12:03:04.6899436Z",
	// 			},
	// 			map[string]any{
	// 				"data":map[string]any{
	// 					"configuration":map[string]any{
	// 						"api":map[string]any{
	// 							"from": "/6452XXXXXXXXXXXX9c2facb1",
	// 							"revision": "1",
	// 							"to":map[string]any{
	// 								"path": "/",
	// 								"host": "msitesting.net",
	// 								"isDefaultPort": true,
	// 								"port": float64(80),
	// 								"query":map[string]any{
	// 								},
	// 								"queryString": "",
	// 								"scheme": "http",
	// 							},
	// 							"version": nil,
	// 						},
	// 						"operation":map[string]any{
	// 							"method": "GET",
	// 							"uriTemplate": "/64524dXXXXXXXXXXXX2facb3",
	// 						},
	// 						"product": "-",
	// 						"user": "-",
	// 					},
	// 				},
	// 				"elapsed": "00:00:00.3046329",
	// 				"source": "api-inspector",
	// 				"timestamp": "2023-05-03T12:03:04.6969650Z",
	// 			},
	// 			map[string]any{
	// 				"data": "Origin header was missing or empty and the request was classified as not cross-domain. CORS policy was not applied.",
	// 				"elapsed": "00:00:00.5972352",
	// 				"source": "cors",
	// 				"timestamp": "2023-05-03T12:03:04.9901631Z",
	// 			},
	// 			map[string]any{
	// 				"data":map[string]any{
	// 					"message":[]any{
	// 						"Response status code was set to 200",
	// 						"Response status reason was set to 'OK'",
	// 					},
	// 				},
	// 				"elapsed": "00:00:00.6107970",
	// 				"source": "set-status",
	// 				"timestamp": "2023-05-03T12:03:05.0031202Z",
	// 			},
	// 			map[string]any{
	// 				"data":map[string]any{
	// 					"message": "Return response was applied",
	// 					"response":map[string]any{
	// 						"headers":[]any{
	// 						},
	// 						"status":map[string]any{
	// 							"code": "OK",
	// 							"reason": "OK",
	// 						},
	// 					},
	// 				},
	// 				"elapsed": "00:00:00.6164228",
	// 				"source": "return-response",
	// 				"timestamp": "2023-05-03T12:03:05.0086543Z",
	// 			},
	// 		},
	// 		"outbound":[]any{
	// 			map[string]any{
	// 				"data":map[string]any{
	// 					"message": "Response has been sent to the caller in full",
	// 				},
	// 				"elapsed": "00:00:00.6510195",
	// 				"source": "transfer-response",
	// 				"timestamp": "2023-05-03T12:03:05.0438287Z",
	// 			},
	// 		},
	// 	},
	// 	"traceId": "1e0447d4-XXXX-XXXX-XXXX-dbdb8098a0d3",
	// }
}
