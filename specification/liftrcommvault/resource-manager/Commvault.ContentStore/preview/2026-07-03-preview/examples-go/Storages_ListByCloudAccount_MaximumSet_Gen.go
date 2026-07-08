package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-07-03-preview/Storages_ListByCloudAccount_MaximumSet_Gen.json
func ExampleStoragesClient_NewListByCloudAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("65D4E6D7-7063-4C4B-BAC5-13C45474009E", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStoragesClient().NewListByCloudAccountPager("rgcommvault", "sample-cloudAccountName", nil)
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
		// page = armcommvaultcontentstore.StoragesClientListByCloudAccountResponse{
		// 	StorageListResult: armcommvaultcontentstore.StorageListResult{
		// 		Value: []*armcommvaultcontentstore.Storage{
		// 			{
		// 				Properties: &armcommvaultcontentstore.StorageProperties{
		// 					Location: to.Ptr("idxxvyqhyxvahalfmkbjxklfbimouxnasnxdmaydqv"),
		// 					StorageType: to.Ptr(armcommvaultcontentstore.StorageTypeAirGapProtect),
		// 					Vendor: to.Ptr(armcommvaultcontentstore.VendorAzureBlobStorage),
		// 					Class: to.Ptr(armcommvaultcontentstore.StorageClassTypeCool),
		// 					ProvisioningState: to.Ptr(armcommvaultcontentstore.ResourceProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/65D4E6D7-7063-4C4B-BAC5-13C45474009E/resourceGroups/rgcommvault/providers/Commvault.ContentStore/cloudAccounts/myCloudAccount/storages/myStorage"),
		// 				Name: to.Ptr("igfptpjqzg"),
		// 				Type: to.Ptr("r"),
		// 				SystemData: &armcommvaultcontentstore.SystemData{
		// 					CreatedBy: to.Ptr("wg"),
		// 					CreatedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-01T06:14:57.355Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("hbpzxzzwhqfy"),
		// 					LastModifiedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-01T06:14:57.355Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
