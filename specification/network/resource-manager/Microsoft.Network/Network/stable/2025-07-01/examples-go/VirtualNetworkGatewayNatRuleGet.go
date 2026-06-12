package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/VirtualNetworkGatewayNatRuleGet.json
func ExampleVirtualNetworkGatewayNatRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkGatewayNatRulesClient().Get(ctx, "rg1", "gateway1", "natRule1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.VirtualNetworkGatewayNatRulesClientGetResponse{
	// 	VirtualNetworkGatewayNatRule: armnetwork.VirtualNetworkGatewayNatRule{
	// 		Name: to.Ptr("natRule1"),
	// 		Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/gateway1/natRules/natRule1"),
	// 		Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
	// 			Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
	// 			ExternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 				{
	// 					AddressSpace: to.Ptr("50.4.0.0/24"),
	// 					PortRange: to.Ptr("200-200"),
	// 				},
	// 			},
	// 			InternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 				{
	// 					AddressSpace: to.Ptr("10.4.0.0/24"),
	// 					PortRange: to.Ptr("100-100"),
	// 				},
	// 			},
	// 			IPConfigurationID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/gateway1/ipConfigurations/default"),
	// 			Mode: to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
