package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VirtualHubListByResourceGroup.json
func ExampleVirtualHubsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualHubsClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.ListVirtualHubsResult = armnetwork.ListVirtualHubsResult{
		// 	Value: []*armnetwork.VirtualHub{
		// 		{
		// 			Name: to.Ptr("virtualHub1"),
		// 			Type: to.Ptr("Microsoft.Network/virtualHubs"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1"),
		// 			Location: to.Ptr("West US"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.VirtualHubProperties{
		// 				AddressPrefix: to.Ptr("10.10.1.0/24"),
		// 				AllowBranchToBranchTraffic: to.Ptr(false),
		// 				HubRoutingPreference: to.Ptr(armnetwork.HubRoutingPreferenceExpressRoute),
		// 				PreferredRoutingGateway: to.Ptr(armnetwork.PreferredRoutingGatewayExpressRoute),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				RoutingState: to.Ptr(armnetwork.RoutingStateProvisioned),
		// 				SKU: to.Ptr("Basic"),
		// 				VirtualHubRouteTableV2S: []*armnetwork.VirtualHubRouteTableV2{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2/routeTables/virtualHubRouteTable2"),
		// 						Name: to.Ptr("rt2a"),
		// 						Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 						Properties: &armnetwork.VirtualHubRouteTableV2Properties{
		// 							AttachedConnections: []*string{
		// 								to.Ptr("All_Vnets")},
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								Routes: []*armnetwork.VirtualHubRouteV2{
		// 									{
		// 										DestinationType: to.Ptr("CIDR"),
		// 										Destinations: []*string{
		// 											to.Ptr("20.10.0.0/16"),
		// 											to.Ptr("20.20.0.0/16")},
		// 											NextHopType: to.Ptr("IPAddress"),
		// 											NextHops: []*string{
		// 												to.Ptr("10.0.0.68")},
		// 											},
		// 											{
		// 												DestinationType: to.Ptr("CIDR"),
		// 												Destinations: []*string{
		// 													to.Ptr("0.0.0.0/0")},
		// 													NextHopType: to.Ptr("IPAddress"),
		// 													NextHops: []*string{
		// 														to.Ptr("10.0.0.68")},
		// 												}},
		// 											},
		// 									}},
		// 									VirtualRouterAsn: to.Ptr[int64](65515),
		// 									VirtualRouterAutoScaleConfiguration: &armnetwork.VirtualRouterAutoScaleConfiguration{
		// 										MinCapacity: to.Ptr[int32](2),
		// 									},
		// 									VirtualRouterIPs: []*string{
		// 										to.Ptr("10.10.1.12"),
		// 										to.Ptr("10.10.1.13")},
		// 										VirtualWan: &armnetwork.SubResource{
		// 											ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWans/virtualWan1"),
		// 										},
		// 									},
		// 								},
		// 								{
		// 									Name: to.Ptr("virtualHub2"),
		// 									Type: to.Ptr("Microsoft.Network/virtualHubs"),
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2"),
		// 									Location: to.Ptr("East US"),
		// 									Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 									Properties: &armnetwork.VirtualHubProperties{
		// 										AddressPrefix: to.Ptr("210.10.1.0/24"),
		// 										AllowBranchToBranchTraffic: to.Ptr(false),
		// 										HubRoutingPreference: to.Ptr(armnetwork.HubRoutingPreferenceExpressRoute),
		// 										PreferredRoutingGateway: to.Ptr(armnetwork.PreferredRoutingGatewayExpressRoute),
		// 										ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 										RoutingState: to.Ptr(armnetwork.RoutingStateProvisioned),
		// 										SKU: to.Ptr("Basic"),
		// 										VirtualHubRouteTableV2S: []*armnetwork.VirtualHubRouteTableV2{
		// 											{
		// 												ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2/routeTables/virtualHubRouteTable2"),
		// 												Name: to.Ptr("rt2a"),
		// 												Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 												Properties: &armnetwork.VirtualHubRouteTableV2Properties{
		// 													AttachedConnections: []*string{
		// 														to.Ptr("All_Vnets")},
		// 														ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 														Routes: []*armnetwork.VirtualHubRouteV2{
		// 															{
		// 																DestinationType: to.Ptr("CIDR"),
		// 																Destinations: []*string{
		// 																	to.Ptr("20.10.0.0/16"),
		// 																	to.Ptr("20.20.0.0/16")},
		// 																	NextHopType: to.Ptr("IPAddress"),
		// 																	NextHops: []*string{
		// 																		to.Ptr("10.0.0.68")},
		// 																	},
		// 																	{
		// 																		DestinationType: to.Ptr("CIDR"),
		// 																		Destinations: []*string{
		// 																			to.Ptr("0.0.0.0/0")},
		// 																			NextHopType: to.Ptr("IPAddress"),
		// 																			NextHops: []*string{
		// 																				to.Ptr("10.0.0.68")},
		// 																		}},
		// 																	},
		// 															}},
		// 															VirtualRouterAsn: to.Ptr[int64](65515),
		// 															VirtualRouterAutoScaleConfiguration: &armnetwork.VirtualRouterAutoScaleConfiguration{
		// 																MinCapacity: to.Ptr[int32](2),
		// 															},
		// 															VirtualRouterIPs: []*string{
		// 																to.Ptr("10.10.1.12"),
		// 																to.Ptr("10.10.1.13")},
		// 																VirtualWan: &armnetwork.SubResource{
		// 																	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWans/virtualWan1"),
		// 																},
		// 															},
		// 													}},
		// 												}
	}
}
