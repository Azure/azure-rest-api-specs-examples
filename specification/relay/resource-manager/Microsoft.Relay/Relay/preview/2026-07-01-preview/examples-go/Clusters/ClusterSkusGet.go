package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay/v2"
)

// Generated from example definition: 2026-07-01-preview/Clusters/ClusterSkusGet.json
func ExampleClustersClient_ListSKUs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().ListSKUs(ctx, "myResourceGroup", "testCluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrelay.ClustersClientListSKUsResponse{
	// 	ClusterSKUListResult: armrelay.ClusterSKUListResult{
	// 		Value: []*armrelay.ClusterSKUInfo{
	// 			{
	// 				ResourceType: to.Ptr("Microsoft.Relay/clusters"),
	// 				SKU: &armrelay.ClusterSKUDetails{
	// 					Name: to.Ptr(armrelay.ClusterSKUNameDedicated),
	// 					Tier: to.Ptr(armrelay.ClusterSKUTierDedicated),
	// 				},
	// 				Capacity: &armrelay.ClusterSKUCapacity{
	// 					AllowedValues: []*int32{
	// 						to.Ptr[int32](3),
	// 						to.Ptr[int32](4),
	// 						to.Ptr[int32](5),
	// 						to.Ptr[int32](6),
	// 						to.Ptr[int32](7),
	// 						to.Ptr[int32](8),
	// 						to.Ptr[int32](9),
	// 						to.Ptr[int32](10),
	// 					},
	// 					Default: to.Ptr[int32](3),
	// 					Maximum: to.Ptr[int32](10),
	// 					Minimum: to.Ptr[int32](3),
	// 					ScaleType: to.Ptr(armrelay.ClusterSKUScaleTypeAutomatic),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
