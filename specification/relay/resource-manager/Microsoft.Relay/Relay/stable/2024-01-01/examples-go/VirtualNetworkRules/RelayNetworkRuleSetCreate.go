package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay/v2"
)

// Generated from example definition: 2024-01-01/VirtualNetworkRules/RelayNetworkRuleSetCreate.json
func ExampleNamespacesClient_CreateOrUpdateNetworkRuleSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("Subscription", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().CreateOrUpdateNetworkRuleSet(ctx, "ResourceGroup", "example-RelayNamespace-6019", armrelay.NetworkRuleSet{
		Properties: &armrelay.NetworkRuleSetProperties{
			DefaultAction: to.Ptr(armrelay.DefaultActionDeny),
			IPRules: []*armrelay.NWRuleSetIPRules{
				{
					Action: to.Ptr(armrelay.NetworkRuleIPActionAllow),
					IPMask: to.Ptr("1.1.1.1"),
				},
				{
					Action: to.Ptr(armrelay.NetworkRuleIPActionAllow),
					IPMask: to.Ptr("1.1.1.2"),
				},
				{
					Action: to.Ptr(armrelay.NetworkRuleIPActionAllow),
					IPMask: to.Ptr("1.1.1.3"),
				},
				{
					Action: to.Ptr(armrelay.NetworkRuleIPActionAllow),
					IPMask: to.Ptr("1.1.1.4"),
				},
				{
					Action: to.Ptr(armrelay.NetworkRuleIPActionAllow),
					IPMask: to.Ptr("1.1.1.5"),
				},
			},
			TrustedServiceAccessEnabled: to.Ptr(false),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrelay.NamespacesClientCreateOrUpdateNetworkRuleSetResponse{
	// 	NetworkRuleSet: &armrelay.NetworkRuleSet{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.Relay/Namespaces/NetworkRuleSet"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroupid/providers/Microsoft.Relay/namespaces/example-RelayNamespace-9659/networkruleset/default"),
	// 		Properties: &armrelay.NetworkRuleSetProperties{
	// 			DefaultAction: to.Ptr(armrelay.DefaultActionDeny),
	// 			IPRules: []*armrelay.NWRuleSetIPRules{
	// 				{
	// 					Action: to.Ptr(armrelay.NetworkRuleIPActionAllow),
	// 					IPMask: to.Ptr("1.1.1.1"),
	// 				},
	// 				{
	// 					Action: to.Ptr(armrelay.NetworkRuleIPActionAllow),
	// 					IPMask: to.Ptr("1.1.1.2"),
	// 				},
	// 				{
	// 					Action: to.Ptr(armrelay.NetworkRuleIPActionAllow),
	// 					IPMask: to.Ptr("1.1.1.3"),
	// 				},
	// 				{
	// 					Action: to.Ptr(armrelay.NetworkRuleIPActionAllow),
	// 					IPMask: to.Ptr("1.1.1.4"),
	// 				},
	// 				{
	// 					Action: to.Ptr(armrelay.NetworkRuleIPActionAllow),
	// 					IPMask: to.Ptr("1.1.1.5"),
	// 				},
	// 			},
	// 			TrustedServiceAccessEnabled: to.Ptr(false),
	// 		},
	// 	},
	// }
}
