package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9975d3476c05bcc6bd9535ad3dfb564e6a168fa5/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/VirtualNetworkPeeringSync.json
func ExampleVirtualNetworkPeeringsClient_BeginCreateOrUpdate_syncPeering() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkPeeringsClient().BeginCreateOrUpdate(ctx, "peerTest", "vnet1", "peer", armnetwork.VirtualNetworkPeering{
		Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
			AllowForwardedTraffic:     to.Ptr(true),
			AllowGatewayTransit:       to.Ptr(false),
			AllowVirtualNetworkAccess: to.Ptr(true),
			RemoteVirtualNetwork: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
			},
			UseRemoteGateways: to.Ptr(false),
		},
	}, &armnetwork.VirtualNetworkPeeringsClientBeginCreateOrUpdateOptions{SyncRemoteAddressSpace: to.Ptr(armnetwork.SyncRemoteAddressSpaceTrue)})
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
	// res.VirtualNetworkPeering = armnetwork.VirtualNetworkPeering{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet1/virtualNetworkPeerings/peer"),
	// 	Name: to.Ptr("peer"),
	// 	Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
	// 		AllowForwardedTraffic: to.Ptr(true),
	// 		AllowGatewayTransit: to.Ptr(false),
	// 		AllowVirtualNetworkAccess: to.Ptr(true),
	// 		PeeringState: to.Ptr(armnetwork.VirtualNetworkPeeringStateInitiated),
	// 		PeeringSyncLevel: to.Ptr(armnetwork.VirtualNetworkPeeringLevelFullyInSync),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RemoteAddressSpace: &armnetwork.AddressSpace{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("12.0.0.0/8")},
	// 			},
	// 			RemoteBgpCommunities: &armnetwork.VirtualNetworkBgpCommunities{
	// 				RegionalCommunity: to.Ptr("12076:50004"),
	// 				VirtualNetworkCommunity: to.Ptr("12076:20002"),
	// 			},
	// 			RemoteVirtualNetwork: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
	// 			},
	// 			RemoteVirtualNetworkAddressSpace: &armnetwork.AddressSpace{
	// 				AddressPrefixes: []*string{
	// 					to.Ptr("12.0.0.0/8")},
	// 				},
	// 				UseRemoteGateways: to.Ptr(false),
	// 			},
	// 		}
}
