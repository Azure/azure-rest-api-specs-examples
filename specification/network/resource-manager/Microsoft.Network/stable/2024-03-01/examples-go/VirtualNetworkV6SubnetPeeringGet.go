package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/VirtualNetworkV6SubnetPeeringGet.json
func ExampleVirtualNetworkPeeringsClient_Get_getV6SubnetPeering() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkPeeringsClient().Get(ctx, "peerTest", "vnet1", "peer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetworkPeering = armnetwork.VirtualNetworkPeering{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet1/virtualNetworkPeerings/peer"),
	// 	Name: to.Ptr("peer"),
	// 	Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
	// 		AllowForwardedTraffic: to.Ptr(true),
	// 		AllowGatewayTransit: to.Ptr(false),
	// 		AllowVirtualNetworkAccess: to.Ptr(true),
	// 		EnableOnlyIPv6Peering: to.Ptr(true),
	// 		LocalAddressSpace: &armnetwork.AddressSpace{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("2002:2002::/64")},
	// 			},
	// 			LocalSubnetNames: []*string{
	// 				to.Ptr("Subnet1")},
	// 				LocalVirtualNetworkAddressSpace: &armnetwork.AddressSpace{
	// 					AddressPrefixes: []*string{
	// 						to.Ptr("2002:2002::/64")},
	// 					},
	// 					PeerCompleteVnets: to.Ptr(false),
	// 					PeeringState: to.Ptr(armnetwork.VirtualNetworkPeeringStateInitiated),
	// 					PeeringSyncLevel: to.Ptr(armnetwork.VirtualNetworkPeeringLevelFullyInSync),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RemoteAddressSpace: &armnetwork.AddressSpace{
	// 						AddressPrefixes: []*string{
	// 							to.Ptr("2001:2001::/64")},
	// 						},
	// 						RemoteBgpCommunities: &armnetwork.VirtualNetworkBgpCommunities{
	// 							RegionalCommunity: to.Ptr("12076:50004"),
	// 							VirtualNetworkCommunity: to.Ptr("12076:20002"),
	// 						},
	// 						RemoteSubnetNames: []*string{
	// 							to.Ptr("Subnet2")},
	// 							RemoteVirtualNetwork: &armnetwork.SubResource{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
	// 							},
	// 							RemoteVirtualNetworkAddressSpace: &armnetwork.AddressSpace{
	// 								AddressPrefixes: []*string{
	// 									to.Ptr("2001:2001::/64")},
	// 								},
	// 								UseRemoteGateways: to.Ptr(false),
	// 							},
	// 						}
}
