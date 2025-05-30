package armconnectedcache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedcache/armconnectedcache"
)

// Generated from example definition: 2023-05-01-preview/CacheNodesOperations_Update_MaximumSet_Gen.json
func ExampleCacheNodesOperationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedcache.NewClientFactory("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCacheNodesOperationsClient().Update(ctx, "rgConnectedCache", "wlrwpdbcv", armconnectedcache.PatchResource{
		Tags: map[string]*string{
			"key5032": to.Ptr("esiuyjbpcwkpqriqiqztxuocv"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconnectedcache.CacheNodesOperationsClientUpdateResponse{
	// 	CacheNodePreviewResource: &armconnectedcache.CacheNodePreviewResource{
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/DoTest/providers/Microsoft.ConnectedCache/cacheNodes/MccRPTest2"),
	// 		Name: to.Ptr("MCCTPTest2"),
	// 		Type: to.Ptr("Microsoft.ConnectedCache/cacheNodes"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armconnectedcache.CacheNodeOldResponse{
	// 			ProvisioningState: to.Ptr(armconnectedcache.ProvisioningStateSucceeded),
	// 			StatusCode: to.Ptr("movtzupooyhdqk"),
	// 			StatusText: to.Ptr("bjnsrpzaofjntleoesjwammgbi"),
	// 			StatusDetails: to.Ptr("quuziibkwtgf"),
	// 			Status: to.Ptr("vxwmuwtqimapfw"),
	// 			Error: &armconnectedcache.ErrorDetail{
	// 				Code: to.Ptr("dzxbdigdjbdbdclvxkxmfutgcbjf"),
	// 				Message: to.Ptr("ifabxmzinicoximnsjkmhdpdgkw"),
	// 				Details: []*armconnectedcache.ErrorDetail{
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key8256": to.Ptr("oreqiywrjkmate"),
	// 		},
	// 		SystemData: &armconnectedcache.SystemData{
	// 			CreatedBy: to.Ptr("yqwxlhphavoggkcwg"),
	// 			CreatedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-31T00:19:33.838Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("knekx"),
	// 			LastModifiedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-31T00:19:33.838Z"); return t}()),
	// 		},
	// 	},
	// }
}
