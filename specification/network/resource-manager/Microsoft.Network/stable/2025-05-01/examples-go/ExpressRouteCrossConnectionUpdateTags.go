package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ExpressRouteCrossConnectionUpdateTags.json
func ExampleExpressRouteCrossConnectionsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRouteCrossConnectionsClient().UpdateTags(ctx, "CrossConnection-SiliconValley", "<circuitServiceKey>", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExpressRouteCrossConnection = armnetwork.ExpressRouteCrossConnection{
	// 	Name: to.Ptr("er1"),
	// 	Type: to.Ptr("Microsoft.Network/expressRouteCrossConnections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/CrossConnectionSiliconValley/providers/Microsoft.Network/expressRouteCrossConnections/<circuitServiceKey>"),
	// 	Location: to.Ptr("brazilsouth"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
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
