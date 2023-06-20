package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e2b5f9323c4214408969a6e953b4075cfdc693b6/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2023-05-01/examples/StorageTargets_CreateOrUpdate.json
func ExampleStorageTargetsClient_BeginCreateOrUpdate_storageTargetsCreateOrUpdate() {
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
			Junctions: []*armstoragecache.NamespaceJunction{
				{
					NamespacePath:   to.Ptr("/path/on/cache"),
					NfsAccessPolicy: to.Ptr("default"),
					NfsExport:       to.Ptr("exp1"),
					TargetPath:      to.Ptr("/path/on/exp1"),
				},
				{
					NamespacePath:   to.Ptr("/path2/on/cache"),
					NfsAccessPolicy: to.Ptr("rootSquash"),
					NfsExport:       to.Ptr("exp2"),
					TargetPath:      to.Ptr("/path2/on/exp2"),
				}},
			Nfs3: &armstoragecache.Nfs3Target{
				Target:            to.Ptr("10.0.44.44"),
				UsageModel:        to.Ptr("READ_ONLY"),
				VerificationTimer: to.Ptr[int32](30),
			},
			TargetType: to.Ptr(armstoragecache.StorageTargetTypeNfs3),
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
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
	// 	},
	// 	Properties: &armstoragecache.StorageTargetProperties{
	// 		Junctions: []*armstoragecache.NamespaceJunction{
	// 			{
	// 				NamespacePath: to.Ptr("/path/on/cache"),
	// 				NfsAccessPolicy: to.Ptr("default"),
	// 				NfsExport: to.Ptr("exp1"),
	// 				TargetPath: to.Ptr("/path/on/exp1"),
	// 			},
	// 			{
	// 				NamespacePath: to.Ptr("/path2/on/cache"),
	// 				NfsAccessPolicy: to.Ptr("rootSquash"),
	// 				NfsExport: to.Ptr("exp2"),
	// 				TargetPath: to.Ptr("/path2/on/exp2"),
	// 		}},
	// 		Nfs3: &armstoragecache.Nfs3Target{
	// 			Target: to.Ptr("10.0.44.44"),
	// 			UsageModel: to.Ptr("READ_ONLY"),
	// 			VerificationTimer: to.Ptr[int32](30),
	// 		},
	// 		State: to.Ptr(armstoragecache.OperationalStateTypeReady),
	// 		TargetType: to.Ptr(armstoragecache.StorageTargetTypeNfs3),
	// 	},
	// }
}
