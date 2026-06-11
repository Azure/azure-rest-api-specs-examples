package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/VirtualNetworkPeeringList.json
func ExampleVirtualNetworkPeeringsClient_NewListPager_listPeerings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkPeeringsClient().NewListPager("peerTest", "vnet1", nil)
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
		// page = armnetwork.VirtualNetworkPeeringsClientListResponse{
		// 	VirtualNetworkPeeringListResult: armnetwork.VirtualNetworkPeeringListResult{
		// 		Value: []*armnetwork.VirtualNetworkPeering{
		// 			{
		// 				Name: to.Ptr("peer"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet1/virtualNetworkPeerings/peer"),
		// 				Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
		// 					AllowForwardedTraffic: to.Ptr(true),
		// 					AllowGatewayTransit: to.Ptr(false),
		// 					AllowVirtualNetworkAccess: to.Ptr(true),
		// 					PeeringState: to.Ptr(armnetwork.VirtualNetworkPeeringStateInitiated),
		// 					PeeringSyncLevel: to.Ptr(armnetwork.VirtualNetworkPeeringLevelFullyInSync),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					RemoteAddressSpace: &armnetwork.AddressSpace{
		// 						AddressPrefixes: []*string{
		// 							to.Ptr("12.0.0.0/8"),
		// 						},
		// 					},
		// 					RemoteBgpCommunities: &armnetwork.VirtualNetworkBgpCommunities{
		// 						RegionalCommunity: to.Ptr("12076:50004"),
		// 						VirtualNetworkCommunity: to.Ptr("12076:20002"),
		// 					},
		// 					RemoteVirtualNetwork: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
		// 					},
		// 					RemoteVirtualNetworkAddressSpace: &armnetwork.AddressSpace{
		// 						AddressPrefixes: []*string{
		// 							to.Ptr("12.0.0.0/8"),
		// 						},
		// 					},
		// 					UseRemoteGateways: to.Ptr(false),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("peer"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet1/virtualNetworkPeerings/peer2"),
		// 				Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
		// 					AllowForwardedTraffic: to.Ptr(false),
		// 					AllowGatewayTransit: to.Ptr(false),
		// 					AllowVirtualNetworkAccess: to.Ptr(true),
		// 					PeeringState: to.Ptr(armnetwork.VirtualNetworkPeeringStateInitiated),
		// 					PeeringSyncLevel: to.Ptr(armnetwork.VirtualNetworkPeeringLevelFullyInSync),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					RemoteAddressSpace: &armnetwork.AddressSpace{
		// 						AddressPrefixes: []*string{
		// 							to.Ptr("13.0.0.0/8"),
		// 						},
		// 					},
		// 					RemoteBgpCommunities: &armnetwork.VirtualNetworkBgpCommunities{
		// 						RegionalCommunity: to.Ptr("12076:50004"),
		// 						VirtualNetworkCommunity: to.Ptr("12076:20003"),
		// 					},
		// 					RemoteVirtualNetwork: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet3"),
		// 					},
		// 					RemoteVirtualNetworkAddressSpace: &armnetwork.AddressSpace{
		// 						AddressPrefixes: []*string{
		// 							to.Ptr("13.0.0.0/8"),
		// 						},
		// 					},
		// 					UseRemoteGateways: to.Ptr(false),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
