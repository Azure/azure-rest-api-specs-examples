package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ExpressRouteCrossConnectionGet.json
func ExampleExpressRouteCrossConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRouteCrossConnectionsClient().Get(ctx, "CrossConnection-SiliconValley", "<circuitServiceKey>", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExpressRouteCrossConnection = armnetwork.ExpressRouteCrossConnection{
	// 	Name: to.Ptr("<circuitServiceKey>"),
	// 	Type: to.Ptr("Microsoft.Network/expressRouteCrossConnections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/CrossConnection-SiliconValley/providers/Microsoft.Network/expressRouteCrossConnections/<circuitServiceKey>"),
	// 	Location: to.Ptr("brazilsouth"),
	// 	Etag: to.Ptr("W/\"c0e6477e-8150-4d4f-9bf6-bb10e6acb63a\""),
	// 	Properties: &armnetwork.ExpressRouteCrossConnectionProperties{
	// 		BandwidthInMbps: to.Ptr[int32](1000),
	// 		ExpressRouteCircuit: &armnetwork.ExpressRouteCircuitReference{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/ertest/providers/Microsoft.Network/expressRouteCircuits/er1"),
	// 		},
	// 		PeeringLocation: to.Ptr("SiliconValley"),
	// 		Peerings: []*armnetwork.ExpressRouteCrossConnectionPeering{
	// 		},
	// 		PrimaryAzurePort: to.Ptr("bvtazureixp01"),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		STag: to.Ptr[int32](2),
	// 		SecondaryAzurePort: to.Ptr("bvtazureixp01"),
	// 		ServiceProviderProvisioningState: to.Ptr(armnetwork.ServiceProviderProvisioningStateNotProvisioned),
	// 	},
	// }
}
