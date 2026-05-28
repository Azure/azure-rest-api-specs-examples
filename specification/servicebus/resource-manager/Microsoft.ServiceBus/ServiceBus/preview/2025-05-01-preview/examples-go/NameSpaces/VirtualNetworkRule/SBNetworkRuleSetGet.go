package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: 2025-05-01-preview/NameSpaces/VirtualNetworkRule/SBNetworkRuleSetGet.json
func ExampleNamespacesClient_GetNetworkRuleSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().GetNetworkRuleSet(ctx, "ResourceGroup", "sdk-Namespace-6019", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicebus.NamespacesClientGetNetworkRuleSetResponse{
	// 	NetworkRuleSet: armservicebus.NetworkRuleSet{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.ServiceBus/Namespaces/NetworkRuleSet"),
	// 		ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-AustraliaEast/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-9659/networkruleset/default"),
	// 		Properties: &armservicebus.NetworkRuleSetProperties{
	// 			DefaultAction: to.Ptr(armservicebus.DefaultActionAllow),
	// 			IPRules: []*armservicebus.NWRuleSetIPRules{
	// 				{
	// 					Action: to.Ptr(armservicebus.NetworkRuleIPActionAllow),
	// 					IPMask: to.Ptr("1.1.1.1"),
	// 				},
	// 				{
	// 					Action: to.Ptr(armservicebus.NetworkRuleIPActionAllow),
	// 					IPMask: to.Ptr("1.1.1.2"),
	// 				},
	// 				{
	// 					Action: to.Ptr(armservicebus.NetworkRuleIPActionAllow),
	// 					IPMask: to.Ptr("1.1.1.3"),
	// 				},
	// 				{
	// 					Action: to.Ptr(armservicebus.NetworkRuleIPActionAllow),
	// 					IPMask: to.Ptr("1.1.1.4"),
	// 				},
	// 				{
	// 					Action: to.Ptr(armservicebus.NetworkRuleIPActionAllow),
	// 					IPMask: to.Ptr("1.1.1.5"),
	// 				},
	// 			},
	// 			PublicNetworkAccess: to.Ptr(armservicebus.PublicNetworkAccessFlagEnabled),
	// 			VirtualNetworkRules: []*armservicebus.NWRuleSetVirtualNetworkRules{
	// 				{
	// 					IgnoreMissingVnetServiceEndpoint: to.Ptr(true),
	// 					Subnet: &armservicebus.Subnet{
	// 						ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourcegroups/alitest/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet2"),
	// 					},
	// 				},
	// 				{
	// 					IgnoreMissingVnetServiceEndpoint: to.Ptr(false),
	// 					Subnet: &armservicebus.Subnet{
	// 						ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourcegroups/alitest/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet3"),
	// 					},
	// 				},
	// 				{
	// 					IgnoreMissingVnetServiceEndpoint: to.Ptr(false),
	// 					Subnet: &armservicebus.Subnet{
	// 						ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourcegroups/alitest/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet6"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
