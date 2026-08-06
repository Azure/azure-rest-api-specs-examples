package armdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/discovery/armdiscovery"
)

// Generated from example definition: 2026-06-01/StorageAssets_ListByStorageContainer_MaximumSet_Gen.json
func ExampleStorageAssetsClient_NewListByStorageContainerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdiscovery.NewClientFactory("A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStorageAssetsClient().NewListByStorageContainerPager("rgdiscovery", "78d6139ad7238f844f", nil)
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
		// page = armdiscovery.StorageAssetsClientListByStorageContainerResponse{
		// 	StorageAssetListResult: armdiscovery.StorageAssetListResult{
		// 		Value: []*armdiscovery.StorageAsset{
		// 			{
		// 				ID: to.Ptr("/subscriptions/A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/storagecontainers/78d6139ad7238f844f/storageAssets/mpxydqppzup"),
		// 				Name: to.Ptr("mpxydqppzup"),
		// 				Tags: map[string]*string{
		// 					"key5959": to.Ptr("oougwvhtjmly"),
		// 				},
		// 				Location: to.Ptr("uksouth"),
		// 				Type: to.Ptr("Microsoft.Discovery/storagecontainers/storageAssets"),
		// 				SystemData: &armdiscovery.SystemData{
		// 					CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
		// 					CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
		// 					LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
		// 				},
		// 				Properties: &armdiscovery.StorageAssetProperties{
		// 					Description: to.Ptr("nopjazrozjrjeruobmiwm"),
		// 					ProvisioningState: to.Ptr(armdiscovery.ProvisioningStateSucceeded),
		// 					Path: to.Ptr("oakrihezlavfyobbhmgqmzowzw"),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
