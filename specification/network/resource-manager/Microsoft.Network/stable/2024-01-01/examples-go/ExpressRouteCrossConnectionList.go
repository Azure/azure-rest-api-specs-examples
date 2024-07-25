package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/ExpressRouteCrossConnectionList.json
func ExampleExpressRouteCrossConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExpressRouteCrossConnectionsClient().NewListPager(&armnetwork.ExpressRouteCrossConnectionsClientListOptions{Filter: nil})
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
		// page.ExpressRouteCrossConnectionListResult = armnetwork.ExpressRouteCrossConnectionListResult{
		// 	Value: []*armnetwork.ExpressRouteCrossConnection{
		// 		{
		// 			Name: to.Ptr("<circuitServiceKey>"),
		// 			Type: to.Ptr("Microsoft.Network/expressRouteCrossConnections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/CrossConnectionSiliconValley/providers/Microsoft.Network/expressRouteCrossConnections/<circuitServiceKey>"),
		// 			Location: to.Ptr("brazilsouth"),
		// 			Properties: &armnetwork.ExpressRouteCrossConnectionProperties{
		// 				BandwidthInMbps: to.Ptr[int32](1000),
		// 				ExpressRouteCircuit: &armnetwork.ExpressRouteCircuitReference{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/ertest/providers/Microsoft.Network/expressRouteCircuits/er1"),
		// 				},
		// 				PeeringLocation: to.Ptr("SiliconValley"),
		// 				Peerings: []*armnetwork.ExpressRouteCrossConnectionPeering{
		// 				},
		// 				PrimaryAzurePort: to.Ptr("bvtazureixp01"),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				STag: to.Ptr[int32](2),
		// 				SecondaryAzurePort: to.Ptr("bvtazureixp01"),
		// 				ServiceProviderProvisioningState: to.Ptr(armnetwork.ServiceProviderProvisioningStateNotProvisioned),
		// 			},
		// 	}},
		// }
	}
}
