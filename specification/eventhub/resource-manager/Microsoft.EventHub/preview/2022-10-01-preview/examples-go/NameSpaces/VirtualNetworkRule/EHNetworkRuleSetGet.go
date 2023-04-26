package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/NameSpaces/VirtualNetworkRule/EHNetworkRuleSetGet.json
func ExampleNamespacesClient_GetNetworkRuleSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.NetworkRuleSet = armeventhub.NetworkRuleSet{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.EventHub/Namespaces/NetworkRuleSet"),
	// 	ID: to.Ptr("/subscriptions/subscriptionid/resourceGroups/resourcegroupid/providers/Microsoft.EventHub/namespaces/sdk-Namespace-9659/networkruleset/default"),
	// 	Properties: &armeventhub.NetworkRuleSetProperties{
	// 		DefaultAction: to.Ptr(armeventhub.DefaultActionDeny),
	// 		IPRules: []*armeventhub.NWRuleSetIPRules{
	// 			{
	// 				Action: to.Ptr(armeventhub.NetworkRuleIPActionAllow),
	// 				IPMask: to.Ptr("1.1.1.1"),
	// 			},
	// 			{
	// 				Action: to.Ptr(armeventhub.NetworkRuleIPActionAllow),
	// 				IPMask: to.Ptr("1.1.1.2"),
	// 			},
	// 			{
	// 				Action: to.Ptr(armeventhub.NetworkRuleIPActionAllow),
	// 				IPMask: to.Ptr("1.1.1.3"),
	// 			},
	// 			{
	// 				Action: to.Ptr(armeventhub.NetworkRuleIPActionAllow),
	// 				IPMask: to.Ptr("1.1.1.4"),
	// 			},
	// 			{
	// 				Action: to.Ptr(armeventhub.NetworkRuleIPActionAllow),
	// 				IPMask: to.Ptr("1.1.1.5"),
	// 		}},
	// 		VirtualNetworkRules: []*armeventhub.NWRuleSetVirtualNetworkRules{
	// 			{
	// 				IgnoreMissingVnetServiceEndpoint: to.Ptr(true),
	// 				Subnet: &armeventhub.Subnet{
	// 					ID: to.Ptr("/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet2"),
	// 				},
	// 			},
	// 			{
	// 				IgnoreMissingVnetServiceEndpoint: to.Ptr(false),
	// 				Subnet: &armeventhub.Subnet{
	// 					ID: to.Ptr("/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet3"),
	// 				},
	// 			},
	// 			{
	// 				IgnoreMissingVnetServiceEndpoint: to.Ptr(false),
	// 				Subnet: &armeventhub.Subnet{
	// 					ID: to.Ptr("/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet6"),
	// 				},
	// 		}},
	// 	},
	// }
}
