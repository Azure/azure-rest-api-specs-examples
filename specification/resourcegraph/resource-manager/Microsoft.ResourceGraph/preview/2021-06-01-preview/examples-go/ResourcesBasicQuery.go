package armresourcegraph_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesBasicQuery.json
func ExampleClient_Resources_basicQuery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcegraph.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Resources(ctx, armresourcegraph.QueryRequest{
		Query: to.Ptr("Resources | project id, name, type, location, tags | limit 3"),
		Subscriptions: []*string{
			to.Ptr("cfbbd179-59d2-4052-aa06-9270a38aa9d6")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryResponse = armresourcegraph.QueryResponse{
	// 	Count: to.Ptr[int64](3),
	// 	Data: []any{
	// 		map[string]any{
	// 			"name": "myNetworkInterface",
	// 			"type": "microsoft.network/networkinterfaces",
	// 			"id": "/subscriptions/cfbbd179-59d2-4052-aa06-9270a38aa9d6/resourceGroups/RG1/providers/Microsoft.Network/networkInterfaces/myNetworkInterface",
	// 			"location": "centralus",
	// 			"tags":map[string]any{
	// 				"tag1": "Value1",
	// 			},
	// 		},
	// 		map[string]any{
	// 			"name": "myVnet",
	// 			"type": "microsoft.network/virtualnetworks",
	// 			"id": "/subscriptions/cfbbd179-59d2-4052-aa06-9270a38aa9d6/resourceGroups/RG2/providers/Microsoft.Network/virtualNetworks/myVnet",
	// 			"location": "westus",
	// 			"tags":map[string]any{
	// 			},
	// 		},
	// 		map[string]any{
	// 			"name": "myPublicIp",
	// 			"type": "microsoft.network/publicipaddresses",
	// 			"id": "/subscriptions/cfbbd179-59d2-4052-aa06-9270a38aa9d6/resourceGroups/RG2/providers/Microsoft.Network/publicIPAddresses/myPublicIp",
	// 			"location": "westus",
	// 			"tags":map[string]any{
	// 			},
	// 		},
	// 	},
	// 	Facets: []armresourcegraph.FacetClassification{
	// 	},
	// 	ResultTruncated: to.Ptr(armresourcegraph.ResultTruncatedFalse),
	// 	TotalRecords: to.Ptr[int64](3),
	// }
}
