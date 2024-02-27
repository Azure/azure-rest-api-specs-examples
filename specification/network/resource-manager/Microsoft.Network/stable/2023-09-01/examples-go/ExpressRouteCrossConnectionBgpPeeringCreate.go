package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/examples/ExpressRouteCrossConnectionBgpPeeringCreate.json
func ExampleExpressRouteCrossConnectionPeeringsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteCrossConnectionPeeringsClient().BeginCreateOrUpdate(ctx, "CrossConnection-SiliconValley", "<circuitServiceKey>", "AzurePrivatePeering", armnetwork.ExpressRouteCrossConnectionPeering{
		Properties: &armnetwork.ExpressRouteCrossConnectionPeeringProperties{
			IPv6PeeringConfig: &armnetwork.IPv6ExpressRouteCircuitPeeringConfig{
				PrimaryPeerAddressPrefix:   to.Ptr("3FFE:FFFF:0:CD30::/126"),
				SecondaryPeerAddressPrefix: to.Ptr("3FFE:FFFF:0:CD30::4/126"),
			},
			PeerASN:                    to.Ptr[int64](200),
			PrimaryPeerAddressPrefix:   to.Ptr("192.168.16.252/30"),
			SecondaryPeerAddressPrefix: to.Ptr("192.168.18.252/30"),
			VlanID:                     to.Ptr[int32](200),
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
	// res.ExpressRouteCrossConnectionPeering = armnetwork.ExpressRouteCrossConnectionPeering{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/CrossConnection-Boydton1DC/providers/Microsoft.Network/expressRouteCrossConnections/<circuitServiceKey>/peerings/AzurePrivatePeering"),
	// 	Name: to.Ptr("AzurePrivatePeering"),
	// 	Etag: to.Ptr("W/\"72090554-7e3b-43f2-80ad-99a9020dcb11\""),
	// 	Properties: &armnetwork.ExpressRouteCrossConnectionPeeringProperties{
	// 		AzureASN: to.Ptr[int32](12076),
	// 		GatewayManagerEtag: to.Ptr(""),
	// 		IPv6PeeringConfig: &armnetwork.IPv6ExpressRouteCircuitPeeringConfig{
	// 			PrimaryPeerAddressPrefix: to.Ptr("3FFE:FFFF:0:CD30::/126"),
	// 			SecondaryPeerAddressPrefix: to.Ptr("3FFE:FFFF:0:CD30::4/126"),
	// 			State: to.Ptr(armnetwork.ExpressRouteCircuitPeeringStateEnabled),
	// 		},
	// 		LastModifiedBy: to.Ptr("Customer"),
	// 		PeerASN: to.Ptr[int64](200),
	// 		PeeringType: to.Ptr(armnetwork.ExpressRoutePeeringTypeAzurePrivatePeering),
	// 		PrimaryAzurePort: to.Ptr(""),
	// 		PrimaryPeerAddressPrefix: to.Ptr("192.168.16.252/30"),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		SecondaryAzurePort: to.Ptr(""),
	// 		SecondaryPeerAddressPrefix: to.Ptr("192.168.18.252/30"),
	// 		State: to.Ptr(armnetwork.ExpressRoutePeeringStateEnabled),
	// 		VlanID: to.Ptr[int32](200),
	// 	},
	// }
}
