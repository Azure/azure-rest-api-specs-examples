package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2026-01-01-preview/AvsStorageContainers_Get_MaximumSet_Gen.json
func ExampleAvsStorageContainersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("11111111-1111-1111-1111-111111111111", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvsStorageContainersClient().Get(ctx, "rgpurestorage", "storagepool-01", "container-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpurestorageblock.AvsStorageContainersClientGetResponse{
	// 	AvsStorageContainer: armpurestorageblock.AvsStorageContainer{
	// 		Properties: &armpurestorageblock.AvsStorageContainerProperties{
	// 			Space: &armpurestorageblock.Space{
	// 				TotalUsed: to.Ptr[int64](1073741824),
	// 				Unique: to.Ptr[int64](268435456),
	// 				Snapshots: to.Ptr[int64](53687091236870910),
	// 				Shared: to.Ptr[int64](268435456),
	// 			},
	// 			ResourceName: to.Ptr("container-01"),
	// 			ProvisionedLimit: to.Ptr[int64](10737418240),
	// 			Datastore: to.Ptr("AVS-Datastore-01"),
	// 			Mounted: to.Ptr(true),
	// 		},
	// 		ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rgpurestorage/providers/PureStorage.Block/storagePools/storagepool-01/avsStorageContainers/container-01"),
	// 		Type: to.Ptr("PureStorage.Block/storagePools/avsStorageContainers"),
	// 		SystemData: &armpurestorageblock.SystemData{
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-01-16T07:25:56.721Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("admin@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-01-16T07:25:56.721Z"); return t}()),
	// 		},
	// 	},
	// }
}
