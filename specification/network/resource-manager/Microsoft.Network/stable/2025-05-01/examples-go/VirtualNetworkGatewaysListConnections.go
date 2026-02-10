package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkGatewaysListConnections.json
func ExampleVirtualNetworkGatewaysClient_NewListConnectionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkGatewaysClient().NewListConnectionsPager("testrg", "test-vpn-gateway-1", nil)
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
		// page.VirtualNetworkGatewayListConnectionsResult = armnetwork.VirtualNetworkGatewayListConnectionsResult{
		// 	Value: []*armnetwork.VirtualNetworkGatewayConnectionListEntity{
		// 		{
		// 			Name: to.Ptr("test-vpn-connection"),
		// 			Type: to.Ptr("Microsoft.Network/connections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/connections/test-vpn-connection"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\\\"00000000-0000-0000-0000-000000000000\\\""),
		// 			Properties: &armnetwork.VirtualNetworkGatewayConnectionListEntityPropertiesFormat{
		// 				ConnectionType: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionTypeVnet2Vnet),
		// 				EgressBytesTransferred: to.Ptr[int64](0),
		// 				EnableBgp: to.Ptr(true),
		// 				IngressBytesTransferred: to.Ptr[int64](0),
		// 				IPSecPolicies: []*armnetwork.IPSecPolicy{
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				RoutingWeight: to.Ptr[int32](22),
		// 				TrafficSelectorPolicies: []*armnetwork.TrafficSelectorPolicy{
		// 				},
		// 				UsePolicyBasedTrafficSelectors: to.Ptr(false),
		// 				VirtualNetworkGateway1: &armnetwork.VirtualNetworkConnectionGatewayReference{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworkGateways/test-vpn-gateway-1"),
		// 				},
		// 				VirtualNetworkGateway2: &armnetwork.VirtualNetworkConnectionGatewayReference{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg-2/providers/Microsoft.Network/virtualNetworkGateways/test-vpn-gateway-2"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
