package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/PeerExpressRouteCircuitConnectionGet.json
func ExamplePeerExpressRouteCircuitConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPeerExpressRouteCircuitConnectionsClient().Get(ctx, "rg1", "ExpressRouteARMCircuitA", "AzurePrivatePeering", "60aee347-e889-4a42-8c1b-0aae8b1e4013", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PeerExpressRouteCircuitConnection = armnetwork.PeerExpressRouteCircuitConnection{
	// 	ID: to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/ExpressRouteARMCircuitA/peerings/AzurePrivatePeering/peerConnections/60aee347-e889-4a42-8c1b-0aae8b1e4013"),
	// 	Name: to.Ptr("60aee347-e889-4a42-8c1b-0aae8b1e4013"),
	// 	Etag: to.Ptr("W/\"6ffbbb06-da20-44ca-a34f-280c4653b1e9\""),
	// 	Properties: &armnetwork.PeerExpressRouteCircuitConnectionPropertiesFormat{
	// 		AddressPrefix: to.Ptr("20.0.0.0/29"),
	// 		AuthResourceGUID: to.Ptr(""),
	// 		CircuitConnectionStatus: to.Ptr(armnetwork.CircuitConnectionStatusConnected),
	// 		ConnectionName: to.Ptr("circuitConnectionWestusEastus"),
	// 		ExpressRouteCircuitPeering: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/ExpressRouteARMCircuitA/peerings/AzurePrivatePeering"),
	// 		},
	// 		PeerExpressRouteCircuitPeering: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/ExpressRouteARMCircuitB/peerings/AzurePrivatePeering"),
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
