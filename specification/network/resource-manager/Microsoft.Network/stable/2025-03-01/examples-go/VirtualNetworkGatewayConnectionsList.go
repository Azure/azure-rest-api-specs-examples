package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkGatewayConnectionsList.json
func ExampleVirtualNetworkGatewayConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkGatewayConnectionsClient().NewListPager("rg1", nil)
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
		// page.VirtualNetworkGatewayConnectionListResult = armnetwork.VirtualNetworkGatewayConnectionListResult{
		// 	Value: []*armnetwork.VirtualNetworkGatewayConnection{
		// 		{
		// 			Name: to.Ptr("conn1"),
		// 			Type: to.Ptr("Microsoft.Network/connections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/connections/conn1"),
		// 			Location: to.Ptr("centralus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 			Properties: &armnetwork.VirtualNetworkGatewayConnectionPropertiesFormat{
		// 				ConnectionMode: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionModeDefault),
		// 				ConnectionProtocol: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionProtocolIKEv1),
		// 				ConnectionType: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionTypeIPsec),
		// 				DpdTimeoutSeconds: to.Ptr[int32](30),
		// 				EgressBytesTransferred: to.Ptr[int64](0),
		// 				EgressNatRules: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw1/natRules/natRule2"),
		// 				}},
		// 				EnableBgp: to.Ptr(false),
		// 				GatewayCustomBgpIPAddresses: []*armnetwork.GatewayCustomBgpIPAddressIPConfiguration{
		// 					{
		// 						CustomBgpIPAddress: to.Ptr("169.254.21.1"),
		// 						IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/default"),
		// 					},
		// 					{
		// 						CustomBgpIPAddress: to.Ptr("169.254.21.3"),
		// 						IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/ActiveActive"),
		// 				}},
		// 				IngressBytesTransferred: to.Ptr[int64](0),
		// 				IngressNatRules: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw1/natRules/natRule1"),
		// 				}},
		// 				IPSecPolicies: []*armnetwork.IPSecPolicy{
		// 				},
		// 				LocalNetworkGateway2: &armnetwork.LocalNetworkGateway{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/localNetworkGateways/localgw1"),
		// 					Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				RoutingWeight: to.Ptr[int32](0),
		// 				TrafficSelectorPolicies: []*armnetwork.TrafficSelectorPolicy{
		// 				},
		// 				UseLocalAzureIPAddress: to.Ptr(false),
		// 				UsePolicyBasedTrafficSelectors: to.Ptr(false),
		// 				VirtualNetworkGateway1: &armnetwork.VirtualNetworkGateway{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw1"),
		// 					Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("conn2"),
		// 			Type: to.Ptr("Microsoft.Network/connections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/connections/conn2"),
		// 			Location: to.Ptr("eastus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 			Properties: &armnetwork.VirtualNetworkGatewayConnectionPropertiesFormat{
		// 				ConnectionMode: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionModeDefault),
		// 				ConnectionProtocol: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionProtocolIKEv2),
		// 				ConnectionType: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionTypeIPsec),
		// 				DpdTimeoutSeconds: to.Ptr[int32](20),
		// 				EgressBytesTransferred: to.Ptr[int64](0),
		// 				EgressNatRules: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw2/natRules/natRule2"),
		// 				}},
		// 				EnableBgp: to.Ptr(false),
		// 				GatewayCustomBgpIPAddresses: []*armnetwork.GatewayCustomBgpIPAddressIPConfiguration{
		// 					{
		// 						CustomBgpIPAddress: to.Ptr("169.254.21.4"),
		// 						IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw2/ipConfigurations/default"),
		// 					},
		// 					{
		// 						CustomBgpIPAddress: to.Ptr("169.254.21.6"),
		// 						IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw2/ipConfigurations/ActiveActive"),
		// 				}},
		// 				IngressBytesTransferred: to.Ptr[int64](0),
		// 				IngressNatRules: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw2/natRules/natRule1"),
		// 				}},
		// 				IPSecPolicies: []*armnetwork.IPSecPolicy{
		// 				},
		// 				LocalNetworkGateway2: &armnetwork.LocalNetworkGateway{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/localNetworkGateways/localgw2"),
		// 					Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				RoutingWeight: to.Ptr[int32](0),
		// 				TrafficSelectorPolicies: []*armnetwork.TrafficSelectorPolicy{
		// 				},
		// 				UseLocalAzureIPAddress: to.Ptr(true),
		// 				UsePolicyBasedTrafficSelectors: to.Ptr(false),
		// 				VirtualNetworkGateway1: &armnetwork.VirtualNetworkGateway{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw2"),
		// 					Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
