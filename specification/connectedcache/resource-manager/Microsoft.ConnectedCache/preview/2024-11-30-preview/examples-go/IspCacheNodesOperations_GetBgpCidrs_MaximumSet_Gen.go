package armconnectedcache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedcache/armconnectedcache"
)

// Generated from example definition: 2024-11-30-preview/IspCacheNodesOperations_GetBgpCidrs_MaximumSet_Gen.json
func ExampleIspCacheNodesOperationsClient_GetBgpCidrs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedcache.NewClientFactory("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIspCacheNodesOperationsClient().GetBgpCidrs(ctx, "rgConnectedCache", "MccRPTest1", "MCCCachenode1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconnectedcache.IspCacheNodesOperationsClientGetBgpCidrsResponse{
	// 	MccCacheNodeBgpCidrDetails: &armconnectedcache.MccCacheNodeBgpCidrDetails{
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/rgConnectedCache/providers/Microsoft.ConnectedCache/ispCustomers/MccRPTest1/ispCacheNodes/MCCCachenode1"),
	// 		Type: to.Ptr("Microsoft.ConnectedCache/ispCustomers/ispCacheNodes"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armconnectedcache.BgpCidrsConfiguration{
	// 			BgpCidrs: []*string{
	// 				to.Ptr("10.1.0.1/23"),
	// 				to.Ptr("10.1.0.1/32"),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key5911": to.Ptr("qllqoxckxvnyhdegqubvqtfediery"),
	// 		},
	// 		SystemData: &armconnectedcache.SystemData{
	// 			CreatedBy: to.Ptr("gambtqj"),
	// 			CreatedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.771Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("qomgaceiessgnuogz"),
	// 			LastModifiedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.771Z"); return t}()),
	// 		},
	// 	},
	// }
}
