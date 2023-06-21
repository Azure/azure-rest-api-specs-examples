package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7f3e601fd326ca910c3d2939b516e15581e7e41/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2023-05-01/examples/StorageTargets_ListByCache.json
func ExampleStorageTargetsClient_NewListByCachePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStorageTargetsClient().NewListByCachePager("scgroup", "sc1", nil)
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
		// page.StorageTargetsResult = armstoragecache.StorageTargetsResult{
		// 	Value: []*armstoragecache.StorageTarget{
		// 		{
		// 			Name: to.Ptr("st1"),
		// 			Type: to.Ptr("Microsoft.StorageCache/Cache/StorageTarget"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/caches/sc1/storagetargets/st1"),
		// 			SystemData: &armstoragecache.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
		// 			},
		// 			Properties: &armstoragecache.StorageTargetProperties{
		// 				AllocationPercentage: to.Ptr[int32](25),
		// 				Junctions: []*armstoragecache.NamespaceJunction{
		// 					{
		// 						NamespacePath: to.Ptr("/path/on/cache"),
		// 						NfsAccessPolicy: to.Ptr("default"),
		// 						NfsExport: to.Ptr("exp1"),
		// 						TargetPath: to.Ptr("/path/on/exp1"),
		// 					},
		// 					{
		// 						NamespacePath: to.Ptr("/path2/on/cache"),
		// 						NfsAccessPolicy: to.Ptr("default"),
		// 						NfsExport: to.Ptr("exp2"),
		// 						TargetPath: to.Ptr("/path2/on/exp2"),
		// 				}},
		// 				Nfs3: &armstoragecache.Nfs3Target{
		// 					Target: to.Ptr("10.0.44.44"),
		// 					UsageModel: to.Ptr("READ_ONLY"),
		// 				},
		// 				State: to.Ptr(armstoragecache.OperationalStateTypeReady),
		// 				TargetType: to.Ptr(armstoragecache.StorageTargetTypeNfs3),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("st2"),
		// 			Type: to.Ptr("Microsoft.StorageCache/Cache/StorageTarget"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/caches/sc1/storagetargets/st2"),
		// 			SystemData: &armstoragecache.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
		// 			},
		// 			Properties: &armstoragecache.StorageTargetProperties{
		// 				AllocationPercentage: to.Ptr[int32](50),
		// 				Clfs: &armstoragecache.ClfsTarget{
		// 					Target: to.Ptr("https://contoso123.blob.core.windows.net/contoso123"),
		// 				},
		// 				Junctions: []*armstoragecache.NamespaceJunction{
		// 					{
		// 						NamespacePath: to.Ptr("/some/arbitrary/place/on/cache"),
		// 						NfsAccessPolicy: to.Ptr("default"),
		// 						TargetPath: to.Ptr("/"),
		// 				}},
		// 				State: to.Ptr(armstoragecache.OperationalStateTypeReady),
		// 				TargetType: to.Ptr(armstoragecache.StorageTargetTypeClfs),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("st3"),
		// 			Type: to.Ptr("Microsoft.StorageCache/Cache/StorageTarget"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/caches/sc1/storagetargets/st3"),
		// 			SystemData: &armstoragecache.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
		// 			},
		// 			Properties: &armstoragecache.StorageTargetProperties{
		// 				AllocationPercentage: to.Ptr[int32](25),
		// 				Junctions: []*armstoragecache.NamespaceJunction{
		// 					{
		// 						NamespacePath: to.Ptr("/some/crazier/place/on/cache"),
		// 						NfsAccessPolicy: to.Ptr("default"),
		// 						NfsExport: to.Ptr(""),
		// 						TargetPath: to.Ptr("/"),
		// 				}},
		// 				State: to.Ptr(armstoragecache.OperationalStateTypeReady),
		// 				TargetType: to.Ptr(armstoragecache.StorageTargetTypeUnknown),
		// 				Unknown: &armstoragecache.UnknownTarget{
		// 					Attributes: map[string]*string{
		// 						"foo": to.Ptr("bar"),
		// 						"foo2": to.Ptr("test"),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
