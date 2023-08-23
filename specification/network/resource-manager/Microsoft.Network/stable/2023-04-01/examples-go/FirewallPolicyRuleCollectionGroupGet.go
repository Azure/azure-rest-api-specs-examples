package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/FirewallPolicyRuleCollectionGroupGet.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_Get_getFirewallPolicyRuleCollectionGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallPolicyRuleCollectionGroupsClient().Get(ctx, "rg1", "firewallPolicy", "ruleCollectionGroup1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FirewallPolicyRuleCollectionGroup = armnetwork.FirewallPolicyRuleCollectionGroup{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/ruleCollectionGroups/ruleCollectionGroup1"),
	// 	Name: to.Ptr("ruleCollectionGroup1"),
	// 	Etag: to.Ptr("W/\"72090554-7e3b-43f2-80ad-99a9020dcb11\""),
	// 	Properties: &armnetwork.FirewallPolicyRuleCollectionGroupProperties{
	// 		Priority: to.Ptr[int32](110),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
	// 			&armnetwork.FirewallPolicyFilterRuleCollection{
	// 				Name: to.Ptr("Example-Filter-Rule-Collection"),
	// 				Priority: to.Ptr[int32](200),
	// 				RuleCollectionType: to.Ptr(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyFilterRuleCollection),
	// 				Action: &armnetwork.FirewallPolicyFilterRuleCollectionAction{
	// 					Type: to.Ptr(armnetwork.FirewallPolicyFilterRuleCollectionActionTypeDeny),
	// 				},
	// 				Rules: []armnetwork.FirewallPolicyRuleClassification{
	// 					&armnetwork.Rule{
	// 						Name: to.Ptr("network-rule1"),
	// 						RuleType: to.Ptr(armnetwork.FirewallPolicyRuleTypeNetworkRule),
	// 						DestinationAddresses: []*string{
	// 							to.Ptr("*")},
	// 							DestinationPorts: []*string{
	// 								to.Ptr("*")},
	// 								IPProtocols: []*armnetwork.FirewallPolicyRuleNetworkProtocol{
	// 									to.Ptr(armnetwork.FirewallPolicyRuleNetworkProtocolTCP)},
	// 									SourceAddresses: []*string{
	// 										to.Ptr("10.1.25.0/24")},
	// 								}},
	// 						}},
	// 					},
	// 				}
}
