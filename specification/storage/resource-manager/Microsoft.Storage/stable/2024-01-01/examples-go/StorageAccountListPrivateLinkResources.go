package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86c6306649b02e542117adb46c61e8019dbd78e9/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/StorageAccountListPrivateLinkResources.json
func ExamplePrivateLinkResourcesClient_ListByStorageAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().ListByStorageAccount(ctx, "res6977", "sto2527", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResourceListResult = armstorage.PrivateLinkResourceListResult{
	// 	Value: []*armstorage.PrivateLinkResource{
	// 		{
	// 			Name: to.Ptr("blob"),
	// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/privateLinkResources"),
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/privateLinkResources/blob"),
	// 			Properties: &armstorage.PrivateLinkResourceProperties{
	// 				GroupID: to.Ptr("blob"),
	// 				RequiredMembers: []*string{
	// 					to.Ptr("blob")},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.blob.core.windows.net")},
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("blob_secondary"),
	// 					Type: to.Ptr("Microsoft.Storage/storageAccounts/privateLinkResources"),
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/privateLinkResources/blob_secondary"),
	// 					Properties: &armstorage.PrivateLinkResourceProperties{
	// 						GroupID: to.Ptr("blob_secondary"),
	// 						RequiredMembers: []*string{
	// 							to.Ptr("blob_secondary")},
	// 							RequiredZoneNames: []*string{
	// 								to.Ptr("privatelink.blob.core.windows.net")},
	// 							},
	// 						},
	// 						{
	// 							Name: to.Ptr("table"),
	// 							Type: to.Ptr("Microsoft.Storage/storageAccounts/privateLinkResources"),
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/privateLinkResources/table"),
	// 							Properties: &armstorage.PrivateLinkResourceProperties{
	// 								GroupID: to.Ptr("table"),
	// 								RequiredMembers: []*string{
	// 									to.Ptr("table")},
	// 									RequiredZoneNames: []*string{
	// 										to.Ptr("privatelink.table.core.windows.net")},
	// 									},
	// 								},
	// 								{
	// 									Name: to.Ptr("table_secondary"),
	// 									Type: to.Ptr("Microsoft.Storage/storageAccounts/privateLinkResources"),
	// 									ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/privateLinkResources/table_secondary"),
	// 									Properties: &armstorage.PrivateLinkResourceProperties{
	// 										GroupID: to.Ptr("table_secondary"),
	// 										RequiredMembers: []*string{
	// 											to.Ptr("table_secondary")},
	// 											RequiredZoneNames: []*string{
	// 												to.Ptr("privatelink.table.core.windows.net")},
	// 											},
	// 										},
	// 										{
	// 											Name: to.Ptr("dfs"),
	// 											Type: to.Ptr("Microsoft.Storage/storageAccounts/privateLinkResources"),
	// 											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/privateLinkResources/dfs"),
	// 											Properties: &armstorage.PrivateLinkResourceProperties{
	// 												GroupID: to.Ptr("dfs"),
	// 												RequiredMembers: []*string{
	// 													to.Ptr("dfs")},
	// 													RequiredZoneNames: []*string{
	// 														to.Ptr("privatelink.dfs.core.windows.net")},
	// 													},
	// 												},
	// 												{
	// 													Name: to.Ptr("dfs_secondary"),
	// 													Type: to.Ptr("Microsoft.Storage/storageAccounts/privateLinkResources"),
	// 													ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/privateLinkResources/dfs_secondary"),
	// 													Properties: &armstorage.PrivateLinkResourceProperties{
	// 														GroupID: to.Ptr("dfs_secondary"),
	// 														RequiredMembers: []*string{
	// 															to.Ptr("dfs_secondary")},
	// 															RequiredZoneNames: []*string{
	// 																to.Ptr("privatelink.dfs.core.windows.net")},
	// 															},
	// 													}},
	// 												}
}
