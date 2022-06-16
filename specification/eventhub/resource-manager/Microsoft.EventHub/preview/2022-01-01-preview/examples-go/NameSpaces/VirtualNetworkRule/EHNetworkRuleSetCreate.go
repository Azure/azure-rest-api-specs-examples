package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-01-01-preview/examples/NameSpaces/VirtualNetworkRule/EHNetworkRuleSetCreate.json
func ExampleNamespacesClient_CreateOrUpdateNetworkRuleSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeventhub.NewNamespacesClient("Subscription", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdateNetworkRuleSet(ctx,
		"ResourceGroup",
		"sdk-Namespace-6019",
		armeventhub.NetworkRuleSet{
			Properties: &armeventhub.NetworkRuleSetProperties{
				DefaultAction: to.Ptr(armeventhub.DefaultActionDeny),
				IPRules: []*armeventhub.NWRuleSetIPRules{
					{
						Action: to.Ptr(armeventhub.NetworkRuleIPActionAllow),
						IPMask: to.Ptr("1.1.1.1"),
					},
					{
						Action: to.Ptr(armeventhub.NetworkRuleIPActionAllow),
						IPMask: to.Ptr("1.1.1.2"),
					},
					{
						Action: to.Ptr(armeventhub.NetworkRuleIPActionAllow),
						IPMask: to.Ptr("1.1.1.3"),
					},
					{
						Action: to.Ptr(armeventhub.NetworkRuleIPActionAllow),
						IPMask: to.Ptr("1.1.1.4"),
					},
					{
						Action: to.Ptr(armeventhub.NetworkRuleIPActionAllow),
						IPMask: to.Ptr("1.1.1.5"),
					}},
				VirtualNetworkRules: []*armeventhub.NWRuleSetVirtualNetworkRules{
					{
						IgnoreMissingVnetServiceEndpoint: to.Ptr(true),
						Subnet: &armeventhub.Subnet{
							ID: to.Ptr("/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet2"),
						},
					},
					{
						IgnoreMissingVnetServiceEndpoint: to.Ptr(false),
						Subnet: &armeventhub.Subnet{
							ID: to.Ptr("/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet3"),
						},
					},
					{
						IgnoreMissingVnetServiceEndpoint: to.Ptr(false),
						Subnet: &armeventhub.Subnet{
							ID: to.Ptr("/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet6"),
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
