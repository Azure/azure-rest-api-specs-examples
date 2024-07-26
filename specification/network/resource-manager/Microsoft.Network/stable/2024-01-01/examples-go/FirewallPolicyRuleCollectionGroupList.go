package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/FirewallPolicyRuleCollectionGroupList.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_NewListPager_listAllFirewallPolicyRuleCollectionGroupsForAGivenFirewallPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirewallPolicyRuleCollectionGroupsClient().NewListPager("rg1", "firewallPolicy", nil)
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
		// page.FirewallPolicyRuleCollectionGroupListResult = armnetwork.FirewallPolicyRuleCollectionGroupListResult{
		// 	Value: []*armnetwork.FirewallPolicyRuleCollectionGroup{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/ruleCollectionGroups/ruleCollectionGroup1"),
		// 			Name: to.Ptr("ruleCollectionGroup1"),
		// 			Etag: to.Ptr("W/\"72090554-7e3b-43f2-80ad-99a9020dcb11\""),
		// 			Properties: &armnetwork.FirewallPolicyRuleCollectionGroupProperties{
		// 				Priority: to.Ptr[int32](110),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
		// 					&armnetwork.FirewallPolicyFilterRuleCollection{
		// 						Name: to.Ptr("Example-Filter-Rule-Collection"),
		// 						Priority: to.Ptr[int32](120),
		// 						RuleCollectionType: to.Ptr(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyFilterRuleCollection),
		// 						Action: &armnetwork.FirewallPolicyFilterRuleCollectionAction{
		// 							Type: to.Ptr(armnetwork.FirewallPolicyFilterRuleCollectionActionTypeDeny),
		// 						},
		// 						Rules: []armnetwork.FirewallPolicyRuleClassification{
		// 							&armnetwork.Rule{
		// 								Name: to.Ptr("network-rule-1"),
		// 								Description: to.Ptr("Network rule"),
		// 								RuleType: to.Ptr(armnetwork.FirewallPolicyRuleTypeNetworkRule),
		// 								DestinationAddresses: []*string{
		// 									to.Ptr("*")},
		// 									DestinationPorts: []*string{
		// 										to.Ptr("*")},
		// 										IPProtocols: []*armnetwork.FirewallPolicyRuleNetworkProtocol{
		// 											to.Ptr(armnetwork.FirewallPolicyRuleNetworkProtocolTCP)},
		// 											SourceAddresses: []*string{
		// 												to.Ptr("10.1.25.0/24")},
		// 										}},
		// 								}},
		// 							},
		// 					}},
		// 				}
	}
}
