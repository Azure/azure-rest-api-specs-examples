package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/ExpressRouteCircuitCreateOnExpressRoutePort.json
func ExampleExpressRouteCircuitsClient_BeginCreateOrUpdate_createExpressRouteCircuitOnExpressRoutePort() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteCircuitsClient().BeginCreateOrUpdate(ctx, "rg1", "expressRouteCircuit1", armnetwork.ExpressRouteCircuit{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.ExpressRouteCircuitPropertiesFormat{
			AuthorizationKey: to.Ptr("b0be57f5-1fba-463b-adec-ffe767354cdd"),
			BandwidthInGbps:  to.Ptr[float32](10),
			ExpressRoutePort: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName"),
			},
		},
		SKU: &armnetwork.ExpressRouteCircuitSKU{
			Name:   to.Ptr("Premium_MeteredData"),
			Family: to.Ptr(armnetwork.ExpressRouteCircuitSKUFamilyMeteredData),
			Tier:   to.Ptr(armnetwork.ExpressRouteCircuitSKUTierPremium),
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
	// 	Name: to.Ptr("expressRouteCircuit1"),
	// 	Type: to.Ptr("Microsoft.Network/expressRouteCircuits"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuit1"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.ExpressRouteCircuitPropertiesFormat{
	// 		AllowClassicOperations: to.Ptr(false),
	// 		AuthorizationKey: to.Ptr("b0be57f5-1fba-463b-adec-ffe767354cdd"),
	// 		AuthorizationStatus: to.Ptr("Enabled"),
	// 		Authorizations: []*armnetwork.ExpressRouteCircuitAuthorization{
	// 		},
	// 		BandwidthInGbps: to.Ptr[float32](10),
	// 		CircuitProvisioningState: to.Ptr("Enabled"),
	// 		ExpressRoutePort: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName"),
	// 		},
	// 		GatewayManagerEtag: to.Ptr("20"),
	// 		Peerings: []*armnetwork.ExpressRouteCircuitPeering{
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ServiceKey: to.Ptr("d281f746-ee01-4d00-8b0a-edec4833772b"),
	// 		ServiceProviderProvisioningState: to.Ptr(armnetwork.ServiceProviderProvisioningStateProvisioned),
	// 	},
	// 	SKU: &armnetwork.ExpressRouteCircuitSKU{
	// 		Name: to.Ptr("Premium_MeteredData"),
	// 		Family: to.Ptr(armnetwork.ExpressRouteCircuitSKUFamilyMeteredData),
	// 		Tier: to.Ptr(armnetwork.ExpressRouteCircuitSKUTierPremium),
	// 	},
	// }
}
