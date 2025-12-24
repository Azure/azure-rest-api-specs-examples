package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ExpressRouteCircuitUpdateTags.json
func ExampleExpressRouteCircuitsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRouteCircuitsClient().UpdateTags(ctx, "ertest", "er1", armnetwork.TagsObject{
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
	// res.ExpressRouteCircuit = armnetwork.ExpressRouteCircuit{
	// 	Name: to.Ptr("er1"),
	// 	Type: to.Ptr("Microsoft.Network/expressRouteCircuits"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/ertest/providers/Microsoft.Network/expressRouteCircuits/er1"),
	// 	Location: to.Ptr("brazilsouth"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armnetwork.ExpressRouteCircuitPropertiesFormat{
	// 		AllowClassicOperations: to.Ptr(false),
	// 		Authorizations: []*armnetwork.ExpressRouteCircuitAuthorization{
	// 		},
	// 		CircuitProvisioningState: to.Ptr("Enabled"),
	// 		GatewayManagerEtag: to.Ptr(""),
	// 		Peerings: []*armnetwork.ExpressRouteCircuitPeering{
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ServiceKey: to.Ptr("0b392c2e-1e9d-46d7-b5e0-9ce90ca6b60c"),
	// 		ServiceProviderProperties: &armnetwork.ExpressRouteCircuitServiceProviderProperties{
	// 			BandwidthInMbps: to.Ptr[int32](1000),
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
