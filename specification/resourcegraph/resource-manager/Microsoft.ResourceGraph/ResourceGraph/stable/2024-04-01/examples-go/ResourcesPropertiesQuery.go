package armresourcegraph_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

// Generated from example definition: 2024-04-01/ResourcesPropertiesQuery.json
func ExampleClient_Resources_accessAPropertiesField() {
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
		Query: to.Ptr("Resources | where type =~ 'Microsoft.Compute/virtualMachines' | summarize count() by tostring(properties.storageProfile.osDisk.osType)"),
		Subscriptions: []*string{
			to.Ptr("cfbbd179-59d2-4052-aa06-9270a38aa9d6"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresourcegraph.ClientResourcesResponse{
	// 	QueryResponse: armresourcegraph.QueryResponse{
	// 		Count: to.Ptr[int64](2),
	// 		Data: []any{
	// 			map[string]any{
	// 				"count": 7,
	// 				"properties_storageProfile_osDisk_osType": "Linux",
	// 			},
	// 			map[string]any{
	// 				"count": 23,
	// 				"properties_storageProfile_osDisk_osType": "Windows",
	// 			},
	// 		},
	// 		Facets: []armresourcegraph.FacetClassification{
	// 		},
	// 		ResultTruncated: to.Ptr(armresourcegraph.ResultTruncatedFalse),
	// 		TotalRecords: to.Ptr[int64](2),
	// 	},
	// }
}
