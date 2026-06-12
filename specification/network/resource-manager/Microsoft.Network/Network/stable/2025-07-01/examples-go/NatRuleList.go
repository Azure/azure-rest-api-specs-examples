package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NatRuleList.json
func ExampleNatRulesClient_NewListByVPNGatewayPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNatRulesClient().NewListByVPNGatewayPager("rg1", "gateway1", nil)
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
		// page = armnetwork.NatRulesClientListByVPNGatewayResponse{
		// 	ListVPNGatewayNatRulesResult: armnetwork.ListVPNGatewayNatRulesResult{
		// 		Value: []*armnetwork.VPNGatewayNatRule{
		// 			{
		// 				Name: to.Ptr("natRule1"),
		// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/natRules/natRule1"),
		// 				Properties: &armnetwork.VPNGatewayNatRuleProperties{
		// 					Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
		// 					EgressVPNSiteLinkConnections: []*armnetwork.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/vpnConnection1/vpnLinkConnections/vpnLinkConnection1"),
		// 						},
		// 					},
		// 					ExternalMappings: []*armnetwork.VPNNatRuleMapping{
		// 						{
		// 							AddressSpace: to.Ptr("192.168.21.0/24"),
		// 						},
		// 					},
		// 					IngressVPNSiteLinkConnections: []*armnetwork.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/vpnConnection1/vpnLinkConnections/vpnLinkConnection2"),
		// 						},
		// 					},
		// 					InternalMappings: []*armnetwork.VPNNatRuleMapping{
		// 						{
		// 							AddressSpace: to.Ptr("10.4.0.0/24"),
		// 						},
		// 					},
		// 					IPConfigurationID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/cloudnet1-VNG/ipConfigurations/default"),
		// 					Mode: to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("natRule2"),
		// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/natRules/natRule2"),
		// 				Properties: &armnetwork.VPNGatewayNatRuleProperties{
		// 					Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
		// 					EgressVPNSiteLinkConnections: []*armnetwork.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/vpnConnection1/vpnLinkConnections/vpnLinkConnection1"),
		// 						},
		// 					},
		// 					ExternalMappings: []*armnetwork.VPNNatRuleMapping{
		// 						{
		// 							AddressSpace: to.Ptr("192.168.21.0/24"),
		// 						},
		// 					},
		// 					IngressVPNSiteLinkConnections: []*armnetwork.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/vpnConnection1/vpnLinkConnections/vpnLinkConnection2"),
		// 						},
		// 					},
		// 					InternalMappings: []*armnetwork.VPNNatRuleMapping{
		// 						{
		// 							AddressSpace: to.Ptr("10.4.0.0/24"),
		// 						},
		// 					},
		// 					IPConfigurationID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/cloudnet1-VNG1/ipConfigurations/default"),
		// 					Mode: to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
