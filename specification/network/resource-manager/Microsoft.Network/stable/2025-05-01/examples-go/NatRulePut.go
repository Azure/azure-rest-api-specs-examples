package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NatRulePut.json
func ExampleNatRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNatRulesClient().BeginCreateOrUpdate(ctx, "rg1", "gateway1", "natRule1", armnetwork.VPNGatewayNatRule{
		Properties: &armnetwork.VPNGatewayNatRuleProperties{
			Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
			ExternalMappings: []*armnetwork.VPNNatRuleMapping{
				{
					AddressSpace: to.Ptr("192.168.21.0/24"),
				}},
			InternalMappings: []*armnetwork.VPNNatRuleMapping{
				{
					AddressSpace: to.Ptr("10.4.0.0/24"),
				}},
			IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/cloudnet1-VNG/ipConfigurations/default"),
			Mode:              to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
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
	// res.VPNGatewayNatRule = armnetwork.VPNGatewayNatRule{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/natRules/natRule1"),
	// 	Name: to.Ptr("natRule1"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VPNGatewayNatRuleProperties{
	// 		Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
	// 		EgressVPNSiteLinkConnections: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/vpnConnection1/vpnLinkConnections/vpnLinkConnection1"),
	// 		}},
	// 		ExternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 			{
	// 				AddressSpace: to.Ptr("192.168.21.0/24"),
	// 		}},
	// 		IngressVPNSiteLinkConnections: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/vpnConnection1/vpnLinkConnections/vpnLinkConnection2"),
	// 		}},
	// 		InternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 			{
	// 				AddressSpace: to.Ptr("10.4.0.0/24"),
	// 		}},
	// 		IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/cloudnet1-VNG/ipConfigurations/default"),
	// 		Mode: to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
