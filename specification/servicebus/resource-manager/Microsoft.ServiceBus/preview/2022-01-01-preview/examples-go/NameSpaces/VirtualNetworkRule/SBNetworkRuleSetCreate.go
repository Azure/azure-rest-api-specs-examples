package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/NameSpaces/VirtualNetworkRule/SBNetworkRuleSetCreate.json
func ExampleNamespacesClient_CreateOrUpdateNetworkRuleSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewNamespacesClient("Subscription", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdateNetworkRuleSet(ctx,
		"ResourceGroup",
		"sdk-Namespace-6019",
		armservicebus.NetworkRuleSet{
			Properties: &armservicebus.NetworkRuleSetProperties{
				DefaultAction: to.Ptr(armservicebus.DefaultActionDeny),
				IPRules: []*armservicebus.NWRuleSetIPRules{
					{
						Action: to.Ptr(armservicebus.NetworkRuleIPActionAllow),
						IPMask: to.Ptr("1.1.1.1"),
					},
					{
						Action: to.Ptr(armservicebus.NetworkRuleIPActionAllow),
						IPMask: to.Ptr("1.1.1.2"),
					},
					{
						Action: to.Ptr(armservicebus.NetworkRuleIPActionAllow),
						IPMask: to.Ptr("1.1.1.3"),
					},
					{
						Action: to.Ptr(armservicebus.NetworkRuleIPActionAllow),
						IPMask: to.Ptr("1.1.1.4"),
					},
					{
						Action: to.Ptr(armservicebus.NetworkRuleIPActionAllow),
						IPMask: to.Ptr("1.1.1.5"),
					}},
				VirtualNetworkRules: []*armservicebus.NWRuleSetVirtualNetworkRules{
					{
						IgnoreMissingVnetServiceEndpoint: to.Ptr(true),
						Subnet: &armservicebus.Subnet{
							ID: to.Ptr("/subscriptions/854d368f-1828-428f-8f3c-f2affa9b2f7d/resourcegroups/alitest/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet2"),
						},
					},
					{
						IgnoreMissingVnetServiceEndpoint: to.Ptr(false),
						Subnet: &armservicebus.Subnet{
							ID: to.Ptr("/subscriptions/854d368f-1828-428f-8f3c-f2affa9b2f7d/resourcegroups/alitest/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet3"),
						},
					},
					{
						IgnoreMissingVnetServiceEndpoint: to.Ptr(false),
						Subnet: &armservicebus.Subnet{
							ID: to.Ptr("/subscriptions/854d368f-1828-428f-8f3c-f2affa9b2f7d/resourcegroups/alitest/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet6"),
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
