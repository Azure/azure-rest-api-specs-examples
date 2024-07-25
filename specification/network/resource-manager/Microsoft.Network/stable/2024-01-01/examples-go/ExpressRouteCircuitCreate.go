package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/ExpressRouteCircuitCreate.json
func ExampleExpressRouteCircuitsClient_BeginCreateOrUpdate_createExpressRouteCircuit() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteCircuitsClient().BeginCreateOrUpdate(ctx, "rg1", "circuitName", armnetwork.ExpressRouteCircuit{
		Location: to.Ptr("Brazil South"),
		Properties: &armnetwork.ExpressRouteCircuitPropertiesFormat{
			AllowClassicOperations: to.Ptr(false),
			Authorizations:         []*armnetwork.ExpressRouteCircuitAuthorization{},
			Peerings:               []*armnetwork.ExpressRouteCircuitPeering{},
			ServiceProviderProperties: &armnetwork.ExpressRouteCircuitServiceProviderProperties{
				BandwidthInMbps:     to.Ptr[int32](200),
				PeeringLocation:     to.Ptr("Silicon Valley"),
				ServiceProviderName: to.Ptr("Equinix"),
			},
		},
		SKU: &armnetwork.ExpressRouteCircuitSKU{
			Name:   to.Ptr("Standard_MeteredData"),
			Family: to.Ptr(armnetwork.ExpressRouteCircuitSKUFamilyMeteredData),
			Tier:   to.Ptr(armnetwork.ExpressRouteCircuitSKUTierStandard),
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
	// res.ExpressRouteCircuit = armnetwork.ExpressRouteCircuit{
	// 	Name: to.Ptr("circuitName"),
	// 	Type: to.Ptr("Microsoft.Network/expressRouteCircuits"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName"),
	// 	Location: to.Ptr("Brazil South"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.ExpressRouteCircuitPropertiesFormat{
	// 		AllowClassicOperations: to.Ptr(false),
	// 		Authorizations: []*armnetwork.ExpressRouteCircuitAuthorization{
	// 		},
	// 		CircuitProvisioningState: to.Ptr("Enabled"),
	// 		Peerings: []*armnetwork.ExpressRouteCircuitPeering{
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ServiceKey: to.Ptr("a1410692-0000-4ceb-b94a-b90b94d398d1"),
	// 		ServiceProviderProperties: &armnetwork.ExpressRouteCircuitServiceProviderProperties{
	// 			BandwidthInMbps: to.Ptr[int32](200),
	// 			PeeringLocation: to.Ptr("Silicon Valley"),
	// 			ServiceProviderName: to.Ptr("Equinix"),
	// 		},
	// 		ServiceProviderProvisioningState: to.Ptr(armnetwork.ServiceProviderProvisioningStateNotProvisioned),
	// 	},
	// 	SKU: &armnetwork.ExpressRouteCircuitSKU{
	// 		Name: to.Ptr("Standard_MeteredData"),
	// 		Family: to.Ptr(armnetwork.ExpressRouteCircuitSKUFamilyMeteredData),
	// 		Tier: to.Ptr(armnetwork.ExpressRouteCircuitSKUTierStandard),
	// 	},
	// }
}
