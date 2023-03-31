package armresourcegraph_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesFirstPageQuery.json
func ExampleClient_Resources_firstPageQuery() {
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
		Options: &armresourcegraph.QueryRequestOptions{
			Skip: to.Ptr[int32](0),
			Top:  to.Ptr[int32](3),
		},
		Query: to.Ptr("Resources | where name contains 'test' | project id, name, type, location"),
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
	// 	SkipToken: to.Ptr("eyAibm8iOiAibHVjayIsICJidXQiOiAibmljZSIsICJ0cnkiOiAiISIgfQ=="),
	// 	Count: to.Ptr[int64](3),
	// 	Data: []any{
	// 		map[string]any{
	// 			"name": "yetanothertest_OsDisk_1_f396cbcb625a457bb69fe2abf5975820",
	// 			"type": "microsoft.compute/disks",
	// 			"id": "/subscriptions/cfbbd179-59d2-4052-aa06-9270a38aa9d6/resourceGroups/RG1/providers/Microsoft.Compute/disks/yetanothertest_OsDisk_1_f396cbcb625a457bb69fe2abf5975820",
	// 			"location": "eastus",
	// 		},
	// 		map[string]any{
	// 			"name": "TestAA",
	// 			"type": "microsoft.automation/automationaccounts",
	// 			"id": "/subscriptions/cfbbd179-59d2-4052-aa06-9270a38aa9d6/resourceGroups/RG2/providers/Microsoft.Automation/automationAccounts/TestAA",
	// 			"location": "westcentralus",
	// 		},
	// 		map[string]any{
	// 			"name": "TestRB",
	// 			"type": "microsoft.automation/automationaccounts/runbooks",
	// 			"id": "/subscriptions/cfbbd179-59d2-4052-aa06-9270a38aa9d6/resourceGroups/RG2/providers/Microsoft.Automation/automationAccounts/TestAA/runbooks/TestRB",
	// 			"location": "westcentralus",
	// 		},
	// 	},
	// 	Facets: []armresourcegraph.FacetClassification{
	// 	},
	// 	ResultTruncated: to.Ptr(armresourcegraph.ResultTruncatedFalse),
	// 	TotalRecords: to.Ptr[int64](386),
	// }
}
