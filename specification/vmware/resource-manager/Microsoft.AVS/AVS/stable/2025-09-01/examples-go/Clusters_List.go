package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/Clusters_List.json
func ExampleClustersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListPager("group1", "cloud1", nil)
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
		// page = armavs.ClustersClientListResponse{
		// 	ClusterList: armavs.ClusterList{
		// 		Value: []*armavs.Cluster{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1"),
		// 				Name: to.Ptr("cluster1"),
		// 				SKU: &armavs.SKU{
		// 					Name: to.Ptr("AV20"),
		// 				},
		// 				Properties: &armavs.ClusterProperties{
		// 					ClusterSize: to.Ptr[int32](3),
		// 					Hosts: []*string{
		// 						to.Ptr("fakehost22.nyc1.kubernetes.center"),
		// 						to.Ptr("fakehost23.nyc1.kubernetes.center"),
		// 						to.Ptr("fakehost24.nyc1.kubernetes.center"),
		// 					},
		// 					ProvisioningState: to.Ptr(armavs.ClusterProvisioningStateSucceeded),
		// 				},
		// 				Type: to.Ptr("Microsoft.AVS/privateClouds/clusters"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
