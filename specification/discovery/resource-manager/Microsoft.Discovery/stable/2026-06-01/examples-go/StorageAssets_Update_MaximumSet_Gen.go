package armdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/discovery/armdiscovery"
)

// Generated from example definition: 2026-06-01/StorageAssets_Update_MaximumSet_Gen.json
func ExampleStorageAssetsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdiscovery.NewClientFactory("A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageAssetsClient().BeginUpdate(ctx, "rgdiscovery", "0d1a35c1a97ad3a5d9", "06a17210ad733943d1", armdiscovery.StorageAsset{
		Properties: &armdiscovery.StorageAssetProperties{
			Description: to.Ptr("hesvphxbjqrgbmdntuohbtmmllmqj"),
		},
		Tags: map[string]*string{
			"key6257": to.Ptr("fspetaqmcldlgu"),
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
	// res = armdiscovery.StorageAssetsClientUpdateResponse{
	// 	StorageAsset: armdiscovery.StorageAsset{
	// 		ID: to.Ptr("/subscriptions/A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/storageContainers/0d1a35c1a97ad3a5d9/storageAssets/06a17210ad733943d1"),
	// 		Name: to.Ptr("06a17210ad733943d1"),
	// 		Tags: map[string]*string{
	// 			"key5959": to.Ptr("oougwvhtjmly"),
	// 		},
	// 		Location: to.Ptr("uksouth"),
	// 		Type: to.Ptr("Microsoft.Discovery/storageContainers/storageAssets"),
	// 		SystemData: &armdiscovery.SystemData{
	// 			CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
	// 			CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
	// 			LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 		},
	// 		Properties: &armdiscovery.StorageAssetProperties{
	// 			Description: to.Ptr("nopjazrozjrjeruobmiwm"),
	// 			ProvisioningState: to.Ptr(armdiscovery.ProvisioningStateSucceeded),
	// 			Path: to.Ptr("oakrihezlavfyobbhmgqmzowzw"),
	// 		},
	// 	},
	// }
}
