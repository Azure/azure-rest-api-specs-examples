package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/220ad9c6554fc7d6d10a89bdb441c1e3b36e3285/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountListBlobInventoryPolicy.json
func ExampleBlobInventoryPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBlobInventoryPoliciesClient().NewListPager("res7687", "sto9699", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ListBlobInventoryPolicy = armstorage.ListBlobInventoryPolicy{
		// 	Value: []*armstorage.BlobInventoryPolicy{
		// 		{
		// 			Name: to.Ptr("DefaultInventoryPolicy"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/inventoryPolicies"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7687/providers/Microsoft.Storage/storageAccounts/sto9699/inventoryPolicies/default"),
		// 			Properties: &armstorage.BlobInventoryPolicyProperties{
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-05T02:53:39.093Z"); return t}()),
		// 				Policy: &armstorage.BlobInventoryPolicySchema{
		// 					Type: to.Ptr(armstorage.InventoryRuleTypeInventory),
		// 					Enabled: to.Ptr(true),
		// 					Rules: []*armstorage.BlobInventoryPolicyRule{
		// 						{
		// 							Name: to.Ptr("inventoryPolicyRule1"),
		// 							Definition: &armstorage.BlobInventoryPolicyDefinition{
		// 								Format: to.Ptr(armstorage.FormatCSV),
		// 								Filters: &armstorage.BlobInventoryPolicyFilter{
		// 									BlobTypes: []*string{
		// 										to.Ptr("blockBlob"),
		// 										to.Ptr("appendBlob"),
		// 										to.Ptr("pageBlob")},
		// 										IncludeBlobVersions: to.Ptr(true),
		// 										IncludeSnapshots: to.Ptr(true),
		// 										PrefixMatch: []*string{
		// 											to.Ptr("inventoryprefix1"),
		// 											to.Ptr("inventoryprefix2")},
		// 										},
		// 										ObjectType: to.Ptr(armstorage.ObjectTypeBlob),
		// 										Schedule: to.Ptr(armstorage.ScheduleDaily),
		// 										SchemaFields: []*string{
		// 											to.Ptr("Name"),
		// 											to.Ptr("Creation-Time"),
		// 											to.Ptr("Last-Modified"),
		// 											to.Ptr("Content-Length"),
		// 											to.Ptr("Content-MD5"),
		// 											to.Ptr("BlobType"),
		// 											to.Ptr("AccessTier"),
		// 											to.Ptr("AccessTierChangeTime"),
		// 											to.Ptr("Snapshot"),
		// 											to.Ptr("VersionId"),
		// 											to.Ptr("IsCurrentVersion"),
		// 											to.Ptr("Metadata")},
		// 										},
		// 										Destination: to.Ptr("container1"),
		// 										Enabled: to.Ptr(true),
		// 								}},
		// 							},
		// 						},
		// 				}},
		// 			}
	}
}
