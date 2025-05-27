package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VirtualNetworkPeeringCreateWithRemoteVirtualNetworkEncryption.json
func ExampleVirtualNetworkPeeringsClient_BeginCreateOrUpdate_createPeeringWithRemoteVirtualNetworkEncryption() {
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
	}, &armnetwork.VirtualNetworkPeeringsClientBeginCreateOrUpdateOptions{SyncRemoteAddressSpace: nil})
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
	// 			RemoteVirtualNetworkEncryption: &armnetwork.VirtualNetworkEncryption{
	// 				Enabled: to.Ptr(true),
	// 				Enforcement: to.Ptr(armnetwork.VirtualNetworkEncryptionEnforcementAllowUnencrypted),
	// 			},
	// 			UseRemoteGateways: to.Ptr(false),
	// 		},
	// 	}
}
