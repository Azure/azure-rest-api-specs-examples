package armresourcegraph_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

// Generated from example definition: 2020-09-01-preview/ResourceChangeDetails.json
func ExampleClient_ResourceChangeDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcegraph.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ResourceChangeDetails(ctx, armresourcegraph.ResourceChangeDetailsRequestParameters{
		ChangeIDs: []*string{
			to.Ptr("53dc0515-b86b-4bc2-979b-e4694ab4a556"),
		},
		ResourceIDs: []*string{
			to.Ptr("/subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresourcegraph.ClientResourceChangeDetailsResponse{
	// 	ResourceChangeDataArray: []*armresourcegraph.ResourceChangeData{
	// 		{
	// 			AfterSnapshot: &armresourcegraph.ResourceChangeDataAfterSnapshot{
	// 				Content: map[string]any{
	// 					"name": "mystorageaccount",
	// 					"type": "Microsoft.Storage/storageAccounts",
	// 					"id": "/subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount",
	// 					"kind": "Storage",
	// 					"location": "westus",
	// 					"properties": map[string]any{
	// 						"creationTime": "2018-07-27T18:37:21.7708872Z",
	// 						"encryption": map[string]any{
	// 							"keySource": "Microsoft.Storage",
	// 							"services": map[string]any{
	// 								"blob": map[string]any{
	// 									"enabled": true,
	// 									"lastEnabledTime": "2018-07-27T18:37:21.8333895Z",
	// 								},
	// 								"file": map[string]any{
	// 									"enabled": true,
	// 									"lastEnabledTime": "2018-07-27T18:37:21.8333895Z",
	// 								},
	// 							},
	// 						},
	// 						"networkAcls": map[string]any{
	// 							"bypass": "AzureServices",
	// 							"defaultAction": "Allow",
	// 							"ipRules": []any{
	// 							},
	// 							"virtualNetworkRules": []any{
	// 							},
	// 						},
	// 						"primaryEndpoints": map[string]any{
	// 							"blob": "https://mystorageaccount.blob.core.windows.net/",
	// 							"file": "https://mystorageaccount.file.core.windows.net/",
	// 							"queue": "https://mystorageaccount.queue.core.windows.net/",
	// 							"table": "https://mystorageaccount.table.core.windows.net/",
	// 						},
	// 						"primaryLocation": "westus",
	// 						"provisioningState": "Succeeded",
	// 						"statusOfPrimary": "available",
	// 						"supportsHttpsTrafficOnly": true,
	// 					},
	// 					"sku": map[string]any{
	// 						"name": "Standard_LRS",
	// 						"tier": "Standard",
	// 					},
	// 					"tags": map[string]any{
	// 					},
	// 				},
	// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-31T01:54:24.42Z"); return t}()),
	// 			},
	// 			BeforeSnapshot: &armresourcegraph.ResourceChangeDataBeforeSnapshot{
	// 				Content: map[string]any{
	// 					"name": "mystorageaccount",
	// 					"type": "Microsoft.Storage/storageAccounts",
	// 					"id": "/subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount",
	// 					"kind": "Storage",
	// 					"location": "westus",
	// 					"properties": map[string]any{
	// 						"creationTime": "2018-07-27T18:37:21.7708872Z",
	// 						"encryption": map[string]any{
	// 							"keySource": "Microsoft.Storage",
	// 							"services": map[string]any{
	// 								"blob": map[string]any{
	// 									"enabled": true,
	// 									"lastEnabledTime": "2018-07-27T18:37:21.8333895Z",
	// 								},
	// 								"file": map[string]any{
	// 									"enabled": true,
	// 									"lastEnabledTime": "2018-07-27T18:37:21.8333895Z",
	// 								},
	// 							},
	// 						},
	// 						"networkAcls": map[string]any{
	// 							"bypass": "AzureServices",
	// 							"defaultAction": "Allow",
	// 							"ipRules": []any{
	// 							},
	// 							"virtualNetworkRules": []any{
	// 							},
	// 						},
	// 						"primaryEndpoints": map[string]any{
	// 							"blob": "https://mystorageaccount.blob.core.windows.net/",
	// 							"file": "https://mystorageaccount.file.core.windows.net/",
	// 							"queue": "https://mystorageaccount.queue.core.windows.net/",
	// 							"table": "https://mystorageaccount.table.core.windows.net/",
	// 						},
	// 						"primaryLocation": "westus",
	// 						"provisioningState": "Succeeded",
	// 						"statusOfPrimary": "available",
	// 						"supportsHttpsTrafficOnly": false,
	// 					},
	// 					"sku": map[string]any{
	// 						"name": "Standard_LRS",
	// 						"tier": "Standard",
	// 					},
	// 					"tags": map[string]any{
	// 					},
	// 				},
	// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-31T01:32:05.993Z"); return t}()),
	// 			},
	// 			ChangeID: to.Ptr("53dc0515-b86b-4bc2-979b-e4694ab4a556"),
	// 			ResourceID: to.Ptr("/subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount"),
	// 		},
	// 	},
	// }
}
