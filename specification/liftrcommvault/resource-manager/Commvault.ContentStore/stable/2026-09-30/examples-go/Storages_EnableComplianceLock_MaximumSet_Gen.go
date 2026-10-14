package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-09-30/Storages_EnableComplianceLock_MaximumSet_Gen.json
func ExampleStoragesClient_EnableComplianceLock() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("65D4E6D7-7063-4C4B-BAC5-13C45474009E", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStoragesClient().EnableComplianceLock(ctx, "rgcommvault", "myCloudAccount", "myStorage", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcommvaultcontentstore.StoragesClientEnableComplianceLockResponse{
	// 	Storage: armcommvaultcontentstore.Storage{
	// 		Properties: &armcommvaultcontentstore.StorageProperties{
	// 			Location: to.Ptr("eastus2"),
	// 			StorageType: to.Ptr(armcommvaultcontentstore.StorageTypeAirGapProtect),
	// 			Vendor: to.Ptr(armcommvaultcontentstore.VendorAzureBlobStorage),
	// 			Class: to.Ptr(armcommvaultcontentstore.StorageClassTypeHot),
	// 			ProvisioningState: to.Ptr(armcommvaultcontentstore.ResourceProvisioningStateSucceeded),
	// 			ComplianceLockStatus: to.Ptr(armcommvaultcontentstore.ComplianceLockStatusEnabled),
	// 		},
	// 		ID: to.Ptr("/subscriptions/65D4E6D7-7063-4C4B-BAC5-13C45474009E/resourceGroups/rgcommvault/providers/Commvault.ContentStore/cloudAccounts/myCloudAccount/storages/myStorage"),
	// 		Name: to.Ptr("myStorage"),
	// 		Type: to.Ptr("Commvault.ContentStore/cloudAccounts/storages"),
	// 		SystemData: &armcommvaultcontentstore.SystemData{
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.June, 15, 10, 30, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.July, 10, 14, 22, 0, 0, time.UTC)),
	// 		},
	// 	},
	// }
}
