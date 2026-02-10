package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/FirewallPolicyNatRuleCollectionGroupGet.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_Get_getFirewallPolicyNatRuleCollectionGroup() {
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
	// 		Priority: to.Ptr[int32](100),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
	// 			&armnetwork.FirewallPolicyNatRuleCollection{
	// 				Name: to.Ptr("NatRC"),
	// 				Priority: to.Ptr[int32](100),
	// 				RuleCollectionType: to.Ptr(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyNatRuleCollection),
	// 				Action: &armnetwork.FirewallPolicyNatRuleCollectionAction{
	// 					Type: to.Ptr(armnetwork.FirewallPolicyNatRuleCollectionActionTypeDNAT),
	// 				},
	// 				Rules: []armnetwork.FirewallPolicyRuleClassification{
	// 					&armnetwork.NatRule{
	// 						Name: to.Ptr("NatRule1"),
	// 						RuleType: to.Ptr(armnetwork.FirewallPolicyRuleTypeNatRule),
	// 						DestinationAddresses: []*string{
	// 							to.Ptr("152.23.32.23")},
	// 							DestinationPorts: []*string{
	// 								to.Ptr("8080")},
	// 								IPProtocols: []*armnetwork.FirewallPolicyRuleNetworkProtocol{
	// 									to.Ptr(armnetwork.FirewallPolicyRuleNetworkProtocolTCP),
	// 									to.Ptr(armnetwork.FirewallPolicyRuleNetworkProtocolUDP)},
	// 									SourceAddresses: []*string{
	// 										to.Ptr("2.2.2.2")},
	// 										SourceIPGroups: []*string{
	// 										},
	// 										TranslatedFqdn: to.Ptr("internalhttpserver"),
	// 										TranslatedPort: to.Ptr("8080"),
	// 								}},
	// 						}},
	// 					},
	// 				}
}
