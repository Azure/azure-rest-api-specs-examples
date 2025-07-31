package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2024-09-01/Clusters_ListZones.json
func ExampleClustersClient_ListZones_clustersListZones() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().ListZones(ctx, "group1", "cloud1", "cluster1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armavs.ClustersClientListZonesResponse{
	// 	ClusterZoneList: &armavs.ClusterZoneList{
	// 		Zones: []*armavs.ClusterZone{
	// 			{
	// 				Hosts: []*string{
	// 					to.Ptr("fakehost22.nyc1.kubernetes.center"),
	// 					to.Ptr("fakehost23.nyc1.kubernetes.center"),
	// 					to.Ptr("fakehost24.nyc1.kubernetes.center"),
	// 				},
	// 				Zone: to.Ptr("2"),
	// 			},
	// 		},
	// 	},
	// }
}
