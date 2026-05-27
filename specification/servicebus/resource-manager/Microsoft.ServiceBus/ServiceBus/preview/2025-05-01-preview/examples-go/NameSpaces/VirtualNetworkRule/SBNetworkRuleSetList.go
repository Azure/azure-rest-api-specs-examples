package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: 2025-05-01-preview/NameSpaces/VirtualNetworkRule/SBNetworkRuleSetList.json
func ExampleNamespacesClient_NewListNetworkRuleSetsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNamespacesClient().NewListNetworkRuleSetsPager("ResourceGroup", "sdk-Namespace-6019", nil)
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
		// page = armservicebus.NamespacesClientListNetworkRuleSetsResponse{
		// 	NetworkRuleSetListResult: armservicebus.NetworkRuleSetListResult{
		// 		Value: []*armservicebus.NetworkRuleSet{
		// 			{
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces/NetworkRuleSet"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/resourcegroupid/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-9659/networkrulesets/default"),
		// 				Properties: &armservicebus.NetworkRuleSetProperties{
		// 					DefaultAction: to.Ptr(armservicebus.DefaultActionDeny),
		// 					IPRules: []*armservicebus.NWRuleSetIPRules{
		// 						{
		// 							Action: to.Ptr(armservicebus.NetworkRuleIPActionAllow),
		// 							IPMask: to.Ptr("1.1.1.1"),
		// 						},
		// 						{
		// 							Action: to.Ptr(armservicebus.NetworkRuleIPActionAllow),
		// 							IPMask: to.Ptr("1.1.1.2"),
		// 						},
		// 						{
		// 							Action: to.Ptr(armservicebus.NetworkRuleIPActionAllow),
		// 							IPMask: to.Ptr("1.1.1.3"),
		// 						},
		// 						{
		// 							Action: to.Ptr(armservicebus.NetworkRuleIPActionAllow),
		// 							IPMask: to.Ptr("1.1.1.4"),
		// 						},
		// 						{
		// 							Action: to.Ptr(armservicebus.NetworkRuleIPActionAllow),
		// 							IPMask: to.Ptr("1.1.1.5"),
		// 						},
		// 					},
		// 					VirtualNetworkRules: []*armservicebus.NWRuleSetVirtualNetworkRules{
		// 						{
		// 							IgnoreMissingVnetServiceEndpoint: to.Ptr(true),
		// 							Subnet: &armservicebus.Subnet{
		// 								ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet2"),
		// 							},
		// 						},
		// 						{
		// 							IgnoreMissingVnetServiceEndpoint: to.Ptr(false),
		// 							Subnet: &armservicebus.Subnet{
		// 								ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet3"),
		// 							},
		// 						},
		// 						{
		// 							IgnoreMissingVnetServiceEndpoint: to.Ptr(false),
		// 							Subnet: &armservicebus.Subnet{
		// 								ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet6"),
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
