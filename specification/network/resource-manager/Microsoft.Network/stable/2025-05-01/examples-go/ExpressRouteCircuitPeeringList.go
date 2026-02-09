package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ExpressRouteCircuitPeeringList.json
func ExampleExpressRouteCircuitPeeringsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExpressRouteCircuitPeeringsClient().NewListPager("rg1", "circuitName", nil)
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
		// page.ExpressRouteCircuitPeeringListResult = armnetwork.ExpressRouteCircuitPeeringListResult{
		// 	Value: []*armnetwork.ExpressRouteCircuitPeering{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName/peerings/MicrosoftPeering"),
		// 			Name: to.Ptr("MicrosoftPeering"),
		// 			Etag: to.Ptr("W/\"b2a25b98-2e6d-4d46-87f2-089de5f6fdf9\""),
		// 			Properties: &armnetwork.ExpressRouteCircuitPeeringPropertiesFormat{
		// 				AzureASN: to.Ptr[int32](12076),
		// 				ExpressRouteConnection: &armnetwork.ExpressRouteConnectionID{
		// 				},
		// 				GatewayManagerEtag: to.Ptr("103"),
		// 				IPv6PeeringConfig: &armnetwork.IPv6ExpressRouteCircuitPeeringConfig{
		// 					MicrosoftPeeringConfig: &armnetwork.ExpressRouteCircuitPeeringConfig{
		// 						AdvertisedCommunities: []*string{
		// 						},
		// 						AdvertisedPublicPrefixes: []*string{
		// 							to.Ptr("3FFE:FFFF:0:CD31::/120")},
		// 							AdvertisedPublicPrefixesState: to.Ptr(armnetwork.ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateValidationNeeded),
		// 							CustomerASN: to.Ptr[int32](23),
		// 							LegacyMode: to.Ptr[int32](0),
		// 							RoutingRegistryName: to.Ptr("ARIN"),
		// 						},
		// 						PrimaryPeerAddressPrefix: to.Ptr("3FFE:FFFF:0:CD30::/126"),
		// 						SecondaryPeerAddressPrefix: to.Ptr("3FFE:FFFF:0:CD30::4/126"),
		// 						State: to.Ptr(armnetwork.ExpressRouteCircuitPeeringStateEnabled),
		// 					},
		// 					LastModifiedBy: to.Ptr("Customer"),
		// 					MicrosoftPeeringConfig: &armnetwork.ExpressRouteCircuitPeeringConfig{
		// 						AdvertisedCommunities: []*string{
		// 						},
		// 						AdvertisedPublicPrefixes: []*string{
		// 							to.Ptr("123.1.0.0/24")},
		// 							AdvertisedPublicPrefixesState: to.Ptr(armnetwork.ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateValidationNeeded),
		// 							CustomerASN: to.Ptr[int32](23),
		// 							LegacyMode: to.Ptr[int32](0),
		// 							RoutingRegistryName: to.Ptr("ARIN"),
		// 						},
		// 						PeerASN: to.Ptr[int64](100),
		// 						PeeringType: to.Ptr(armnetwork.ExpressRoutePeeringTypeMicrosoftPeering),
		// 						PrimaryAzurePort: to.Ptr("A51-TEST-06GMR-CIS-1-PRI-A"),
		// 						PrimaryPeerAddressPrefix: to.Ptr("123.0.0.0/30"),
		// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						SecondaryAzurePort: to.Ptr("A51-TEST-06GMR-CIS-2-SEC-A"),
		// 						SecondaryPeerAddressPrefix: to.Ptr("123.0.0.4/30"),
		// 						State: to.Ptr(armnetwork.ExpressRoutePeeringStateEnabled),
		// 						VlanID: to.Ptr[int32](300),
		// 					},
		// 				},
		// 				{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName/peerings/AzurePrivatePeering"),
		// 					Name: to.Ptr("AzurePrivatePeering"),
		// 					Etag: to.Ptr("W/\"b2a25b98-2e6d-4d46-87f2-089de5f6fdf9\""),
		// 					Properties: &armnetwork.ExpressRouteCircuitPeeringPropertiesFormat{
		// 						AzureASN: to.Ptr[int32](12076),
		// 						ExpressRouteConnection: &armnetwork.ExpressRouteConnectionID{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/expressRouteGatewayName/expressRouteConnections/connectionName"),
		// 						},
		// 						GatewayManagerEtag: to.Ptr("103"),
		// 						IPv6PeeringConfig: &armnetwork.IPv6ExpressRouteCircuitPeeringConfig{
		// 							PrimaryPeerAddressPrefix: to.Ptr("3FFE:FFFF:0:CD30::/126"),
		// 							SecondaryPeerAddressPrefix: to.Ptr("3FFE:FFFF:0:CD30::4/126"),
		// 							State: to.Ptr(armnetwork.ExpressRouteCircuitPeeringStateEnabled),
		// 						},
		// 						LastModifiedBy: to.Ptr("Customer"),
		// 						PeerASN: to.Ptr[int64](100),
		// 						PeeringType: to.Ptr(armnetwork.ExpressRoutePeeringTypeAzurePrivatePeering),
		// 						PrimaryAzurePort: to.Ptr("A51-TEST-06GMR-CIS-1-PRI-A"),
		// 						PrimaryPeerAddressPrefix: to.Ptr("10.0.0.0/30"),
		// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						SecondaryAzurePort: to.Ptr("A51-TEST-06GMR-CIS-2-SEC-A"),
		// 						SecondaryPeerAddressPrefix: to.Ptr("10.0.0.4/30"),
		// 						State: to.Ptr(armnetwork.ExpressRoutePeeringStateEnabled),
		// 						VlanID: to.Ptr[int32](200),
		// 					},
		// 			}},
		// 		}
	}
}
