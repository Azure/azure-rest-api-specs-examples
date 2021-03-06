package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewayNatRulePut.json
func ExampleVirtualNetworkGatewayNatRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewVirtualNetworkGatewayNatRulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<virtual-network-gateway-name>",
		"<nat-rule-name>",
		armnetwork.VirtualNetworkGatewayNatRule{
			Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
				Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
				ExternalMappings: []*armnetwork.VPNNatRuleMapping{
					{
						AddressSpace: to.Ptr("<address-space>"),
						PortRange:    to.Ptr("<port-range>"),
					}},
				InternalMappings: []*armnetwork.VPNNatRuleMapping{
					{
						AddressSpace: to.Ptr("<address-space>"),
						PortRange:    to.Ptr("<port-range>"),
					}},
				IPConfigurationID: to.Ptr("<ipconfiguration-id>"),
				Mode:              to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
			},
		},
		&armnetwork.VirtualNetworkGatewayNatRulesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
