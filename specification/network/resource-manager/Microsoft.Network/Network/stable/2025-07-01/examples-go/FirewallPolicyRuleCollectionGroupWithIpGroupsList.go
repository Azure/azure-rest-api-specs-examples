package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/FirewallPolicyRuleCollectionGroupWithIpGroupsList.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_NewListPager_listAllFirewallPolicyRuleCollectionGroupsWithIPGroupsForAGivenFirewallPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
		// page = armnetwork.FirewallPolicyRuleCollectionGroupsClientListResponse{
		// 	FirewallPolicyRuleCollectionGroupListResult: armnetwork.FirewallPolicyRuleCollectionGroupListResult{
		// 		Value: []*armnetwork.FirewallPolicyRuleCollectionGroup{
		// 			{
		// 				Name: to.Ptr("ruleCollectionGroup1"),
		// 				Etag: to.Ptr("W/\"72090554-7e3b-43f2-80ad-99a9020dcb11\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/ruleCollectionGroups/ruleCollectionGroup1"),
		// 				Properties: &armnetwork.FirewallPolicyRuleCollectionGroupProperties{
		// 					Priority: to.Ptr[int32](110),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
		// 						&armnetwork.FirewallPolicyFilterRuleCollection{
		// 							Name: to.Ptr("Example-Filter-Rule-Collection"),
		// 							Action: &armnetwork.FirewallPolicyFilterRuleCollectionAction{
		// 								Type: to.Ptr(armnetwork.FirewallPolicyFilterRuleCollectionActionTypeDeny),
		// 							},
		// 							Priority: to.Ptr[int32](120),
		// 							RuleCollectionType: to.Ptr(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyFilterRuleCollection),
		// 							Rules: []armnetwork.FirewallPolicyRuleClassification{
		// 								&armnetwork.Rule{
		// 									Name: to.Ptr("network-rule-1"),
		// 									Description: to.Ptr("Network rule"),
		// 									DestinationIPGroups: []*string{
		// 										to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/resourceGroups/rg1/ipGroups/ipGroups2"),
		// 									},
		// 									DestinationPorts: []*string{
		// 										to.Ptr("*"),
		// 									},
		// 									IPProtocols: []*armnetwork.FirewallPolicyRuleNetworkProtocol{
		// 										to.Ptr(armnetwork.FirewallPolicyRuleNetworkProtocolTCP),
		// 									},
		// 									RuleType: to.Ptr(armnetwork.FirewallPolicyRuleTypeNetworkRule),
		// 									SourceIPGroups: []*string{
		// 										to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/resourceGroups/rg1/ipGroups/ipGroups1"),
		// 									},
		// 								},
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
