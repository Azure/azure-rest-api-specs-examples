package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ExpressRouteCrossConnectionUpdate.json
func ExampleExpressRouteCrossConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteCrossConnectionsClient().BeginCreateOrUpdate(ctx, "CrossConnection-SiliconValley", "<circuitServiceKey>", armnetwork.ExpressRouteCrossConnection{
		Properties: &armnetwork.ExpressRouteCrossConnectionProperties{
			ServiceProviderProvisioningState: to.Ptr(armnetwork.ServiceProviderProvisioningStateNotProvisioned),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ExpressRouteCrossConnectionsClientCreateOrUpdateResponse{
	// 	ExpressRouteCrossConnection: armnetwork.ExpressRouteCrossConnection{
	// 		Name: to.Ptr("<circuitServiceKey>"),
	// 		Type: to.Ptr("Microsoft.Network/expressRouteCrossConnections"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/CrossConnectionSiliconValley/providers/Microsoft.Network/expressRouteCrossConnections/<circuitServiceKey>"),
	// 		Location: to.Ptr("brazilsouth"),
	// 		Properties: &armnetwork.ExpressRouteCrossConnectionProperties{
	// 			BandwidthInMbps: to.Ptr[int32](1000),
	// 			ExpressRouteCircuit: &armnetwork.ExpressRouteCircuitReference{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ertest/providers/Microsoft.Network/expressRouteCircuits/er1"),
	// 			},
	// 			PeeringLocation: to.Ptr("SiliconValley"),
	// 			Peerings: []*armnetwork.ExpressRouteCrossConnectionPeering{
	// 			},
	// 			PrimaryAzurePort: to.Ptr("bvtazureixp01"),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningState("Enabled")),
	// 			STag: to.Ptr[int32](2),
	// 			SecondaryAzurePort: to.Ptr("bvtazureixp01"),
	// 			ServiceProviderProvisioningState: to.Ptr(armnetwork.ServiceProviderProvisioningStateNotProvisioned),
	// 		},
	// 	},
	// }
}
