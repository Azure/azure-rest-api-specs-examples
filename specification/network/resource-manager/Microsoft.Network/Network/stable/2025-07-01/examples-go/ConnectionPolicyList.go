package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ConnectionPolicyList.json
func ExampleConnectionPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectionPoliciesClient().NewListPager("rg1", "TestHub", nil)
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
		// page = armnetwork.ConnectionPoliciesClientListResponse{
		// 	ListConnectionPoliciesResult: armnetwork.ListConnectionPoliciesResult{
		// 		Value: []*armnetwork.ConnectionPolicy{
		// 			{
		// 				Name: to.Ptr("testpolicy"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/TestHub/connectionPolicies/testpolicy"),
		// 				Etag: to.Ptr("W/\"2b03e4fa-c9ed-403c-815f-d8bd6d40a37b\""),
		// 				Properties: &armnetwork.ConnectionPolicyProperties{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					EnableInternetSecurity: to.Ptr(true),
		// 					AssociatedConnections: []*string{
		// 					},
		// 					RoutingConfiguration: &armnetwork.RoutingConfiguration{
		// 						AssociatedRouteTable: &armnetwork.SubResource{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/TestHub/hubRouteTables/defaultRouteTable"),
		// 						},
		// 						PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
		// 							Labels: []*string{
		// 								to.Ptr("default"),
		// 							},
		// 							IDs: []*armnetwork.SubResource{
		// 							},
		// 						},
		// 					},
		// 				},
		// 				Type: to.Ptr("Microsoft.Network/virtualHubs/connectionPolicies"),
		// 			},
		// 			{
		// 				Name: to.Ptr("testpolicy2"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/TestHub/connectionPolicies/testpolicy2"),
		// 				Etag: to.Ptr("W/\"2b03e4fa-c9ed-403c-815f-d8bd6d40a37b\""),
		// 				Properties: &armnetwork.ConnectionPolicyProperties{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					EnableInternetSecurity: to.Ptr(true),
		// 					AssociatedConnections: []*string{
		// 					},
		// 					RoutingConfiguration: &armnetwork.RoutingConfiguration{
		// 						AssociatedRouteTable: &armnetwork.SubResource{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/TestHub/hubRouteTables/defaultRouteTable"),
		// 						},
		// 						PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
		// 							Labels: []*string{
		// 								to.Ptr("default"),
		// 							},
		// 							IDs: []*armnetwork.SubResource{
		// 							},
		// 						},
		// 						InboundRouteMap: &armnetwork.SubResource{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/TestHub/routeMaps/TestRouteMap"),
		// 						},
		// 						OutboundRouteMap: &armnetwork.SubResource{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/TestHub/routeMaps/TestRouteMap"),
		// 						},
		// 					},
		// 				},
		// 				Type: to.Ptr("Microsoft.Network/virtualHubs/connectionPolicies"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
