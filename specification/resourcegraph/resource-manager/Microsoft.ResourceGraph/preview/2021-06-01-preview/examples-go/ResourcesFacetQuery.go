package armresourcegraph_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesFacetQuery.json
func ExampleClient_Resources_queryWithAFacetRequest() {
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
		Facets: []*armresourcegraph.FacetRequest{
			{
				Expression: to.Ptr("location"),
				Options: &armresourcegraph.FacetRequestOptions{
					Top:       to.Ptr[int32](3),
					SortOrder: to.Ptr(armresourcegraph.FacetSortOrderDesc),
				},
			},
			{
				Expression: to.Ptr("properties.storageProfile.osDisk.osType"),
				Options: &armresourcegraph.FacetRequestOptions{
					Top:       to.Ptr[int32](3),
					SortOrder: to.Ptr(armresourcegraph.FacetSortOrderDesc),
				},
			},
			{
				Expression: to.Ptr("nonExistingColumn"),
				Options: &armresourcegraph.FacetRequestOptions{
					Top:       to.Ptr[int32](3),
					SortOrder: to.Ptr(armresourcegraph.FacetSortOrderDesc),
				},
			},
			{
				Expression: to.Ptr("resourceGroup"),
				Options: &armresourcegraph.FacetRequestOptions{
					Top:       to.Ptr[int32](3),
					SortBy:    to.Ptr("tolower(resourceGroup)"),
					SortOrder: to.Ptr(armresourcegraph.FacetSortOrderAsc),
				},
			},
			{
				Expression: to.Ptr("resourceGroup"),
				Options: &armresourcegraph.FacetRequestOptions{
					Top:    to.Ptr[int32](3),
					Filter: to.Ptr("resourceGroup contains 'test'"),
				},
			}},
		Query: to.Ptr("Resources | where type =~ 'Microsoft.Compute/virtualMachines' | project id, name, location, resourceGroup, properties.storageProfile.osDisk.osType | limit 5"),
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
	// 	Count: to.Ptr[int64](5),
	// 	Data: []any{
	// 		map[string]any{
	// 			"name": "myTestVm",
	// 			"id": "/subscriptions/cfbbd179-59d2-4052-aa06-9270a38aa9d6/resourceGroups/B-TEST-RG/providers/Microsoft.Compute/virtualMachines/myTestVm",
	// 			"location": "eastus",
	// 			"properties_storageProfile_osDisk_osType": "Windows",
	// 			"resourceGroup": "B-TEST-RG",
	// 		},
	// 		map[string]any{
	// 			"name": "myTestAccountVm",
	// 			"id": "/subscriptions/cfbbd179-59d2-4052-aa06-9270a38aa9d6/resourceGroups/c-rg/providers/Microsoft.Compute/virtualMachines/myTestAccountVm",
	// 			"location": "westcentralus",
	// 			"properties_storageProfile_osDisk_osType": "Windows",
	// 			"resourceGroup": "c-rg",
	// 		},
	// 		map[string]any{
	// 			"name": "yetanothertest",
	// 			"id": "/subscriptions/cfbbd179-59d2-4052-aa06-9270a38aa9d6/resourceGroups/I-RG/providers/Microsoft.Compute/virtualMachines/yetanothertest",
	// 			"location": "eastus",
	// 			"properties_storageProfile_osDisk_osType": "Linux",
	// 			"resourceGroup": "I-RG",
	// 		},
	// 		map[string]any{
	// 			"name": "drafttest1bux4cv7a7q3aw",
	// 			"id": "/subscriptions/cfbbd179-59d2-4052-aa06-9270a38aa9d6/resourceGroups/x-test-rg/providers/Microsoft.Compute/virtualMachines/drafttest1bux4cv7a7q3aw",
	// 			"location": "southcentralus",
	// 			"properties_storageProfile_osDisk_osType": "Linux",
	// 			"resourceGroup": "x-test-rg",
	// 		},
	// 		map[string]any{
	// 			"name": "testvmntp25370",
	// 			"id": "/subscriptions/cfbbd179-59d2-4052-aa06-9270a38aa9d6/resourceGroups/y-rg/providers/Microsoft.Compute/virtualMachines/testvmntp25370",
	// 			"location": "eastus",
	// 			"properties_storageProfile_osDisk_osType": "Windows",
	// 			"resourceGroup": "y-rg",
	// 		},
	// 	},
	// 	Facets: []armresourcegraph.FacetClassification{
	// 		&armresourcegraph.FacetResult{
	// 			Expression: to.Ptr("location"),
	// 			ResultType: to.Ptr("FacetResult"),
	// 			Count: to.Ptr[int32](3),
	// 			Data: []any{
	// 				map[string]any{
	// 					"count": float64(3),
	// 					"location": "eastus",
	// 				},
	// 				map[string]any{
	// 					"count": float64(1),
	// 					"location": "southcentralus",
	// 				},
	// 				map[string]any{
	// 					"count": float64(1),
	// 					"location": "westcentralus",
	// 				},
	// 			},
	// 			TotalRecords: to.Ptr[int64](3),
	// 		},
	// 		&armresourcegraph.FacetResult{
	// 			Expression: to.Ptr("properties.storageProfile.osDisk.osType"),
	// 			ResultType: to.Ptr("FacetResult"),
	// 			Count: to.Ptr[int32](2),
	// 			Data: []any{
	// 				map[string]any{
	// 					"count": float64(2),
	// 					"properties_storageProfile_osDisk_osType": "Linux",
	// 				},
	// 				map[string]any{
	// 					"count": float64(3),
	// 					"properties_storageProfile_osDisk_osType": "Windows",
	// 				},
	// 			},
	// 			TotalRecords: to.Ptr[int64](2),
	// 		},
	// 		&armresourcegraph.FacetError{
	// 			Expression: to.Ptr("nonExistingColumn"),
	// 			ResultType: to.Ptr("FacetError"),
	// 			Errors: []*armresourcegraph.ErrorDetails{
	// 				{
	// 					Code: to.Ptr("NoValidColumns"),
	// 					Message: to.Ptr("No valid columns in facet expression."),
	// 				},
	// 				{
	// 					Code: to.Ptr("InvalidColumnNames"),
	// 					Message: to.Ptr("Invalid column names: [nonExistingColumn]."),
	// 			}},
	// 		},
	// 		&armresourcegraph.FacetResult{
	// 			Expression: to.Ptr("resourceGroup"),
	// 			ResultType: to.Ptr("FacetResult"),
	// 			Count: to.Ptr[int32](3),
	// 			Data: []any{
	// 				map[string]any{
	// 					"count": float64(1),
	// 					"resourceGroup": "B-TEST-RG",
	// 				},
	// 				map[string]any{
	// 					"count": float64(1),
	// 					"resourceGroup": "c-rg",
	// 				},
	// 				map[string]any{
	// 					"count": float64(1),
	// 					"resourceGroup": "I-RG",
	// 				},
	// 			},
	// 			TotalRecords: to.Ptr[int64](5),
	// 		},
	// 		&armresourcegraph.FacetResult{
	// 			Expression: to.Ptr("resourceGroup"),
	// 			ResultType: to.Ptr("FacetResult"),
	// 			Count: to.Ptr[int32](2),
	// 			Data: []any{
	// 				map[string]any{
	// 					"count": float64(1),
	// 					"resourceGroup": "B-TEST-RG",
	// 				},
	// 				map[string]any{
	// 					"count": float64(1),
	// 					"resourceGroup": "x-test-rg",
	// 				},
	// 			},
	// 			TotalRecords: to.Ptr[int64](2),
	// 	}},
	// 	ResultTruncated: to.Ptr(armresourcegraph.ResultTruncatedFalse),
	// 	TotalRecords: to.Ptr[int64](5),
	// }
}
