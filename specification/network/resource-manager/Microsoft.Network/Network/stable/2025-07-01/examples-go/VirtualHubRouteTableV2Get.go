package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/VirtualHubRouteTableV2Get.json
func ExampleVirtualHubRouteTableV2SClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualHubRouteTableV2SClient().Get(ctx, "rg1", "virtualHub1", "virtualHubRouteTable1a", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.VirtualHubRouteTableV2SClientGetResponse{
	// 	VirtualHubRouteTableV2: armnetwork.VirtualHubRouteTableV2{
	// 		Name: to.Ptr("virtualHubRouteTable1a"),
	// 		Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeTables/virtualHubRouteTable1a"),
	// 		Properties: &armnetwork.VirtualHubRouteTableV2Properties{
	// 			AttachedConnections: []*string{
	// 				to.Ptr("All_Vnets"),
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			Routes: []*armnetwork.VirtualHubRouteV2{
	// 				{
	// 					DestinationType: to.Ptr("CIDR"),
	// 					Destinations: []*string{
	// 						to.Ptr("20.10.0.0/16"),
	// 						to.Ptr("20.20.0.0/16"),
	// 					},
	// 					NextHopType: to.Ptr("IPAddress"),
	// 					NextHops: []*string{
	// 						to.Ptr("10.0.0.68"),
	// 					},
	// 				},
	// 				{
	// 					DestinationType: to.Ptr("CIDR"),
	// 					Destinations: []*string{
	// 						to.Ptr("0.0.0.0/0"),
	// 					},
	// 					NextHopType: to.Ptr("IPAddress"),
	// 					NextHops: []*string{
	// 						to.Ptr("10.0.0.68"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
