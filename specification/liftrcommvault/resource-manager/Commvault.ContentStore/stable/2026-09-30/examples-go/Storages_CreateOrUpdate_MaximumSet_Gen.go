package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-09-30/Storages_CreateOrUpdate_MaximumSet_Gen.json
func ExampleStoragesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("65D4E6D7-7063-4C4B-BAC5-13C45474009E", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStoragesClient().BeginCreateOrUpdate(ctx, "rgcommvault", "sample-cloudAccountName", "sample-storageName", armcommvaultcontentstore.Storage{
		Properties: &armcommvaultcontentstore.StorageProperties{
			Location:    to.Ptr("idxxvyqhyxvahalfmkbjxklfbimouxnasnxdmaydqv"),
			StorageType: to.Ptr(armcommvaultcontentstore.StorageTypeAirGapProtect),
			Vendor:      to.Ptr(armcommvaultcontentstore.VendorAzureBlobStorage),
			Class:       to.Ptr(armcommvaultcontentstore.StorageClassTypeCool),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcommvaultcontentstore.StoragesClientCreateOrUpdateResponse{
	// 	Storage: armcommvaultcontentstore.Storage{
	// 		Properties: &armcommvaultcontentstore.StorageProperties{
	// 			Location: to.Ptr("idxxvyqhyxvahalfmkbjxklfbimouxnasnxdmaydqv"),
	// 			StorageType: to.Ptr(armcommvaultcontentstore.StorageTypeAirGapProtect),
	// 			Vendor: to.Ptr(armcommvaultcontentstore.VendorAzureBlobStorage),
	// 			Class: to.Ptr(armcommvaultcontentstore.StorageClassTypeCool),
	// 			ProvisioningState: to.Ptr(armcommvaultcontentstore.ResourceProvisioningStateSucceeded),
	// 			ComplianceLockStatus: to.Ptr(armcommvaultcontentstore.ComplianceLockStatusEnabled),
	// 		},
	// 		ID: to.Ptr("/subscriptions/65D4E6D7-7063-4C4B-BAC5-13C45474009E/resourceGroups/rgcommvault/providers/Commvault.ContentStore/cloudAccounts/myCloudAccount/storages/myStorage"),
	// 		Name: to.Ptr("igfptpjqzg"),
	// 		Type: to.Ptr("r"),
	// 		SystemData: &armcommvaultcontentstore.SystemData{
	// 			CreatedBy: to.Ptr("wg"),
	// 			CreatedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.April, 1, 6, 14, 57, 355000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("hbpzxzzwhqfy"),
	// 			LastModifiedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.April, 1, 6, 14, 57, 355000000, time.UTC)),
	// 		},
	// 	},
	// }
}
