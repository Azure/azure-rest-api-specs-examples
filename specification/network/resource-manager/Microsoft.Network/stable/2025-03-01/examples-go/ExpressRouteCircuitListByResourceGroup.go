package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ExpressRouteCircuitListByResourceGroup.json
func ExampleExpressRouteCircuitsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExpressRouteCircuitsClient().NewListPager("rg1", nil)
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
		// page.ExpressRouteCircuitListResult = armnetwork.ExpressRouteCircuitListResult{
		// 	Value: []*armnetwork.ExpressRouteCircuit{
		// 		{
		// 			Name: to.Ptr("circuitName1"),
		// 			Type: to.Ptr("Microsoft.Network/expressRouteCircuits"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName1"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"832b28c3-f5fd-4d2a-a2cb-6e4a2fe452b3\""),
		// 			Properties: &armnetwork.ExpressRouteCircuitPropertiesFormat{
		// 				AllowClassicOperations: to.Ptr(false),
		// 				Authorizations: []*armnetwork.ExpressRouteCircuitAuthorization{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName/authorizations/MyAuthorization1"),
		// 						Name: to.Ptr("MyAuthorization1"),
		// 						Etag: to.Ptr("W/\"832b28c3-f5fd-4d2a-a2cb-6e4a2fe452b3\""),
		// 						Properties: &armnetwork.AuthorizationPropertiesFormat{
		// 							AuthorizationKey: to.Ptr("authkey"),
		// 							AuthorizationUseStatus: to.Ptr(armnetwork.AuthorizationUseStatusAvailable),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						},
		// 				}},
		// 				CircuitProvisioningState: to.Ptr("Enabled"),
		// 				GatewayManagerEtag: to.Ptr("113"),
		// 				Peerings: []*armnetwork.ExpressRouteCircuitPeering{
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ServiceKey: to.Ptr("a1410692-ed3b-4ceb-b94a-b90b95d398d1"),
		// 				ServiceProviderProperties: &armnetwork.ExpressRouteCircuitServiceProviderProperties{
		// 					BandwidthInMbps: to.Ptr[int32](200),
		// 					PeeringLocation: to.Ptr("peeringLocation"),
		// 					ServiceProviderName: to.Ptr("providerName"),
		// 				},
		// 				ServiceProviderProvisioningState: to.Ptr(armnetwork.ServiceProviderProvisioningStateProvisioned),
		// 			},
		// 			SKU: &armnetwork.ExpressRouteCircuitSKU{
		// 				Name: to.Ptr("Standard_MeteredData"),
		// 				Family: to.Ptr(armnetwork.ExpressRouteCircuitSKUFamilyMeteredData),
		// 				Tier: to.Ptr(armnetwork.ExpressRouteCircuitSKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("circuitName2"),
		// 			Type: to.Ptr("Microsoft.Network/expressRouteCircuits"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName2"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"e33c875f-48df-4a91-b7d3-eb95b5ddbb89\""),
		// 			Properties: &armnetwork.ExpressRouteCircuitPropertiesFormat{
		// 				AllowClassicOperations: to.Ptr(false),
		// 				Authorizations: []*armnetwork.ExpressRouteCircuitAuthorization{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName2/authorizations/MyAuthorization2"),
		// 						Name: to.Ptr("MyAuthorization2"),
		// 						Etag: to.Ptr("W/\"e33c875f-48df-4a91-b7d3-eb95b5ddbb89\""),
		// 						Properties: &armnetwork.AuthorizationPropertiesFormat{
		// 							AuthorizationKey: to.Ptr("authkey"),
		// 							AuthorizationUseStatus: to.Ptr(armnetwork.AuthorizationUseStatusAvailable),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						},
		// 				}},
		// 				CircuitProvisioningState: to.Ptr("Enabled"),
		// 				GatewayManagerEtag: to.Ptr(""),
		// 				Peerings: []*armnetwork.ExpressRouteCircuitPeering{
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ServiceKey: to.Ptr("6569625a-9ba4-498b-9719-14d778eef609"),
		// 				ServiceProviderProperties: &armnetwork.ExpressRouteCircuitServiceProviderProperties{
		// 					BandwidthInMbps: to.Ptr[int32](200),
		// 					PeeringLocation: to.Ptr("peeringLocation"),
		// 					ServiceProviderName: to.Ptr("providerName"),
		// 				},
		// 				ServiceProviderProvisioningState: to.Ptr(armnetwork.ServiceProviderProvisioningStateNotProvisioned),
		// 			},
		// 			SKU: &armnetwork.ExpressRouteCircuitSKU{
		// 				Name: to.Ptr("Standard_MeteredData"),
		// 				Family: to.Ptr(armnetwork.ExpressRouteCircuitSKUFamilyMeteredData),
		// 				Tier: to.Ptr(armnetwork.ExpressRouteCircuitSKUTierStandard),
		// 			},
		// 	}},
		// }
	}
}
