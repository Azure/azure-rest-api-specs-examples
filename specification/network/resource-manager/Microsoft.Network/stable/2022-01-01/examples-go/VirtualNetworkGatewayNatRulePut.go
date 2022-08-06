package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualNetworkGatewayNatRulePut.json
func ExampleVirtualNetworkGatewayNatRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewVirtualNetworkGatewayNatRulesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "gateway1", "natRule1", armnetwork.VirtualNetworkGatewayNatRule{
		Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
			Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
			ExternalMappings: []*armnetwork.VPNNatRuleMapping{
				{
					AddressSpace: to.Ptr("192.168.21.0/24"),
					PortRange:    to.Ptr("300-400"),
				}},
			InternalMappings: []*armnetwork.VPNNatRuleMapping{
				{
					AddressSpace: to.Ptr("10.4.0.0/24"),
					PortRange:    to.Ptr("200-300"),
				}},
			IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/gateway1/ipConfigurations/default"),
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
	// TODO: use response item
	_ = res
}
