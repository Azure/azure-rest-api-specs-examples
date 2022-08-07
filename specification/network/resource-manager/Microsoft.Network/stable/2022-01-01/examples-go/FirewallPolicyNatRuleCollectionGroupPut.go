package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/FirewallPolicyNatRuleCollectionGroupPut.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_BeginCreateOrUpdate_firewallPolicyNatRuleCollectionGroupPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPolicyRuleCollectionGroupsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "firewallPolicy", "ruleCollectionGroup1", armnetwork.FirewallPolicyRuleCollectionGroup{
		Properties: &armnetwork.FirewallPolicyRuleCollectionGroupProperties{
			Priority: to.Ptr[int32](100),
			RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
				&armnetwork.FirewallPolicyNatRuleCollection{
					Name:               to.Ptr("Example-Nat-Rule-Collection"),
					Priority:           to.Ptr[int32](100),
					RuleCollectionType: to.Ptr(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyNatRuleCollection),
					Action: &armnetwork.FirewallPolicyNatRuleCollectionAction{
						Type: to.Ptr(armnetwork.FirewallPolicyNatRuleCollectionActionTypeDNAT),
					},
					Rules: []armnetwork.FirewallPolicyRuleClassification{
						&armnetwork.NatRule{
							Name:     to.Ptr("nat-rule1"),
							RuleType: to.Ptr(armnetwork.FirewallPolicyRuleTypeNatRule),
							DestinationAddresses: []*string{
								to.Ptr("152.23.32.23")},
							DestinationPorts: []*string{
								to.Ptr("8080")},
							IPProtocols: []*armnetwork.FirewallPolicyRuleNetworkProtocol{
								to.Ptr(armnetwork.FirewallPolicyRuleNetworkProtocolTCP),
								to.Ptr(armnetwork.FirewallPolicyRuleNetworkProtocolUDP)},
							SourceAddresses: []*string{
								to.Ptr("2.2.2.2")},
							SourceIPGroups: []*string{},
							TranslatedFqdn: to.Ptr("internalhttp.server.net"),
							TranslatedPort: to.Ptr("8080"),
						}},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
