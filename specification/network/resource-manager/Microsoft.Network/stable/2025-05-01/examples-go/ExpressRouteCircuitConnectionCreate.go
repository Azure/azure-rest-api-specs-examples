package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ExpressRouteCircuitConnectionCreate.json
func ExampleExpressRouteCircuitConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteCircuitConnectionsClient().BeginCreateOrUpdate(ctx, "rg1", "ExpressRouteARMCircuitA", "AzurePrivatePeering", "circuitConnectionUSAUS", armnetwork.ExpressRouteCircuitConnection{
		Properties: &armnetwork.ExpressRouteCircuitConnectionPropertiesFormat{
			AddressPrefix:    to.Ptr("10.0.0.0/29"),
			AuthorizationKey: to.Ptr("946a1918-b7a2-4917-b43c-8c4cdaee006a"),
			ExpressRouteCircuitPeering: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid1/resourceGroups/dedharcktinit/providers/Microsoft.Network/expressRouteCircuits/dedharcktlocal/peerings/AzurePrivatePeering"),
			},
			IPv6CircuitConnectionConfig: &armnetwork.IPv6CircuitConnectionConfig{
				AddressPrefix: to.Ptr("aa:bb::/125"),
			},
			PeerExpressRouteCircuitPeering: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid2/resourceGroups/dedharcktpeer/providers/Microsoft.Network/expressRouteCircuits/dedharcktremote/peerings/AzurePrivatePeering"),
			},
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
	// res.ExpressRouteCircuitConnection = armnetwork.ExpressRouteCircuitConnection{
	// 	ID: to.Ptr("/subscriptions/subid1/resourceGroups/dedharcktinit/providers/Microsoft.Network/expressRouteCircuits/ExpressRouteARMCircuitA/peerings/AzurePrivatePeering/connections/circuitConnectionUSAUS"),
	// 	Name: to.Ptr("circuitConnectionUSAUS"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.ExpressRouteCircuitConnectionPropertiesFormat{
	// 		AddressPrefix: to.Ptr("10.0.0.0/24"),
	// 		AuthorizationKey: to.Ptr("946a1918-b7a2-4917-b43c-8c4cdaee006a"),
	// 		CircuitConnectionStatus: to.Ptr(armnetwork.CircuitConnectionStatusConnected),
	// 		ExpressRouteCircuitPeering: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid1/resourceGroups/dedharcktinit/providers/Microsoft.Network/expressRouteCircuits/dedharcktlocal/peerings/AzurePrivatePeering"),
	// 		},
	// 		IPv6CircuitConnectionConfig: &armnetwork.IPv6CircuitConnectionConfig{
	// 			AddressPrefix: to.Ptr("aa:bb::1/125"),
	// 			CircuitConnectionStatus: to.Ptr(armnetwork.CircuitConnectionStatusConnected),
	// 		},
	// 		PeerExpressRouteCircuitPeering: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid2/resourceGroups/dedharcktpeer/providers/Microsoft.Network/expressRouteCircuits/dedharcktremote/peerings/AzurePrivatePeering"),
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
