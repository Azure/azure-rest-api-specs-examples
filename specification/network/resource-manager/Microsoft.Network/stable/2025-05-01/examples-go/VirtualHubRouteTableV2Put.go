package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualHubRouteTableV2Put.json
func ExampleVirtualHubRouteTableV2SClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualHubRouteTableV2SClient().BeginCreateOrUpdate(ctx, "rg1", "virtualHub1", "virtualHubRouteTable1a", armnetwork.VirtualHubRouteTableV2{
		Properties: &armnetwork.VirtualHubRouteTableV2Properties{
			AttachedConnections: []*string{
				to.Ptr("All_Vnets")},
			Routes: []*armnetwork.VirtualHubRouteV2{
				{
					DestinationType: to.Ptr("CIDR"),
					Destinations: []*string{
						to.Ptr("20.10.0.0/16"),
						to.Ptr("20.20.0.0/16")},
					NextHopType: to.Ptr("IPAddress"),
					NextHops: []*string{
						to.Ptr("10.0.0.68")},
				},
				{
					DestinationType: to.Ptr("CIDR"),
					Destinations: []*string{
						to.Ptr("0.0.0.0/0")},
					NextHopType: to.Ptr("IPAddress"),
					NextHops: []*string{
						to.Ptr("10.0.0.68")},
				}},
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
	// res.VirtualHubRouteTableV2 = armnetwork.VirtualHubRouteTableV2{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeTables/virtualHubRouteTable1a"),
	// 	Name: to.Ptr("virtualHubRouteTable1a"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VirtualHubRouteTableV2Properties{
	// 		AttachedConnections: []*string{
	// 			to.Ptr("All_Vnets")},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			Routes: []*armnetwork.VirtualHubRouteV2{
	// 				{
	// 					DestinationType: to.Ptr("CIDR"),
	// 					Destinations: []*string{
	// 						to.Ptr("20.10.0.0/16"),
	// 						to.Ptr("20.20.0.0/16")},
	// 						NextHopType: to.Ptr("IPAddress"),
	// 						NextHops: []*string{
	// 							to.Ptr("10.0.0.68")},
	// 						},
	// 						{
	// 							DestinationType: to.Ptr("CIDR"),
	// 							Destinations: []*string{
	// 								to.Ptr("0.0.0.0/0")},
	// 								NextHopType: to.Ptr("IPAddress"),
	// 								NextHops: []*string{
	// 									to.Ptr("10.0.0.68")},
	// 							}},
	// 						},
	// 					}
}
