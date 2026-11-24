package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay/v2"
)

// Generated from example definition: 2026-07-01-preview/Clusters/ClusterGet.json
func ExampleClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().Get(ctx, "myResourceGroup", "testCluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrelay.ClustersClientGetResponse{
	// 	Cluster: armrelay.Cluster{
	// 		ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/myResourceGroup/providers/Microsoft.Relay/clusters/testCluster"),
	// 		Location: to.Ptr("South Central US"),
	// 		Name: to.Ptr("testCluster"),
	// 		Properties: &armrelay.ClusterProperties{
	// 			MetricID: to.Ptr("PROD-SN-VR51001"),
	// 			ProvisioningState: to.Ptr(armrelay.ClusterProvisioningStateSucceeded),
	// 			Status: to.Ptr("Active"),
	// 			SupportsScaling: to.Ptr(true),
	// 			ZoneRedundant: to.Ptr(true),
	// 		},
	// 		SKU: &armrelay.ClusterSKU{
	// 			Name: to.Ptr(armrelay.ClusterSKUNameDedicated),
	// 			Tier: to.Ptr(armrelay.ClusterSKUTierDedicated),
	// 			Capacity: to.Ptr[int32](3),
	// 		},
	// 		SystemData: &armrelay.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2026, time.July, 15, 20, 30, 0, 0, time.UTC)),
	// 			CreatedBy: to.Ptr("user@example.com"),
	// 			CreatedByType: to.Ptr(armrelay.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.July, 15, 20, 35, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@example.com"),
	// 			LastModifiedByType: to.Ptr(armrelay.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 		},
	// 		Type: to.Ptr("Microsoft.Relay/clusters"),
	// 	},
	// }
}
