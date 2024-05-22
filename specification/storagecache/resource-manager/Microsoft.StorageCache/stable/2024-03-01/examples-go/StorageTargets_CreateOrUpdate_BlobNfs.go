package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/StorageTargets_CreateOrUpdate_BlobNfs.json
func ExampleStorageTargetsClient_BeginCreateOrUpdate_storageTargetsCreateOrUpdateBlobNfs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageTargetsClient().BeginCreateOrUpdate(ctx, "scgroup", "sc1", "st1", armstoragecache.StorageTarget{
		Properties: &armstoragecache.StorageTargetProperties{
			BlobNfs: &armstoragecache.BlobNfsTarget{
				Target:            to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/blofnfs/blobServices/default/containers/blobnfs"),
				UsageModel:        to.Ptr("READ_WRITE"),
				VerificationTimer: to.Ptr[int32](28800),
				WriteBackTimer:    to.Ptr[int32](3600),
			},
			Junctions: []*armstoragecache.NamespaceJunction{
				{
					NamespacePath: to.Ptr("/blobnfs"),
				}},
			TargetType: to.Ptr(armstoragecache.StorageTargetTypeBlobNfs),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StorageTarget = armstoragecache.StorageTarget{
	// 	Name: to.Ptr("st1"),
	// 	Type: to.Ptr("Microsoft.StorageCache/Cache/StorageTarget"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/caches/sc1/storagetargets/st1"),
	// 	SystemData: &armstoragecache.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
	// 	},
	// 	Properties: &armstoragecache.StorageTargetProperties{
	// 		BlobNfs: &armstoragecache.BlobNfsTarget{
	// 			Target: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/blofnfs/blobServices/default/containers/blobnfs"),
	// 			UsageModel: to.Ptr("READ_WRITE"),
	// 			VerificationTimer: to.Ptr[int32](28800),
	// 			WriteBackTimer: to.Ptr[int32](3600),
	// 		},
	// 		Junctions: []*armstoragecache.NamespaceJunction{
	// 			{
	// 				NamespacePath: to.Ptr("/blobnfs"),
	// 		}},
	// 		State: to.Ptr(armstoragecache.OperationalStateTypeReady),
	// 		TargetType: to.Ptr(armstoragecache.StorageTargetTypeBlobNfs),
	// 	},
	// }
}
