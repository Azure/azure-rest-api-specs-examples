package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v12"
)

// Generated from example definition: 2026-01-01/ExpressRouteCircuitCreateOnExpressRouteLag.json
func ExampleExpressRouteCircuitsClient_BeginCreateOrUpdate_createExpressRouteCircuitOnExpressRouteLag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteCircuitsClient().BeginCreateOrUpdate(ctx, "rg1", "expressRouteCircuit1", armnetwork.ExpressRouteCircuit{
		Location: to.Ptr("eastus2euap"),
		Properties: &armnetwork.ExpressRouteCircuitPropertiesFormat{
			BandwidthInGbps:           to.Ptr[float32](5),
			EnableDirectPortRateLimit: to.Ptr(true),
			ExpressRouteLag: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName"),
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
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ExpressRouteCircuitsClientCreateOrUpdateResponse{
	// 	ExpressRouteCircuit: armnetwork.ExpressRouteCircuit{
	// 		Name: to.Ptr("expressRouteCircuit1"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuit1"),
	// 		Etag: to.Ptr("W/\"75acaf40-7411-4909-82de-23af323d5feb\""),
	// 		Type: to.Ptr("Microsoft.Network/expressRouteCircuits"),
	// 		Location: to.Ptr("eastus2euap"),
	// 		Properties: &armnetwork.ExpressRouteCircuitPropertiesFormat{
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateUpdating),
	// 			Peerings: []*armnetwork.ExpressRouteCircuitPeering{
	// 			},
	// 			Authorizations: []*armnetwork.ExpressRouteCircuitAuthorization{
	// 			},
	// 			CircuitProvisioningState: to.Ptr("Disabled"),
	// 			AllowClassicOperations: to.Ptr(false),
	// 			ServiceKey: to.Ptr("a0466ed3-8a36-4036-9aab-25cb33c6e5a1"),
	// 			ServiceProviderProvisioningState: to.Ptr(armnetwork.ServiceProviderProvisioningStateProvisioned),
	// 			GlobalReachEnabled: to.Ptr(false),
	// 			ExpressRouteLag: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName"),
	// 			},
	// 			EnableDirectPortRateLimit: to.Ptr(true),
	// 			BandwidthInGbps: to.Ptr[float32](5),
	// 			ResiliencyLevel: to.Ptr(armnetwork.ResiliencyLevelStandard),
	// 		},
	// 		SKU: &armnetwork.ExpressRouteCircuitSKU{
	// 			Name: to.Ptr("Premium_MeteredData"),
	// 			Tier: to.Ptr(armnetwork.ExpressRouteCircuitSKUTierPremium),
	// 			Family: to.Ptr(armnetwork.ExpressRouteCircuitSKUFamilyMeteredData),
	// 		},
	// 	},
	// }
}
