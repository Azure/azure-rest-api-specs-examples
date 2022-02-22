Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/WafPolicyCreateOrUpdate.json
func ExampleWebApplicationFirewallPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewWebApplicationFirewallPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<policy-name>",
		armnetwork.WebApplicationFirewallPolicy{
			Location: to.StringPtr("<location>"),
			Properties: &armnetwork.WebApplicationFirewallPolicyPropertiesFormat{
				CustomRules: []*armnetwork.WebApplicationFirewallCustomRule{
					{
						Name:   to.StringPtr("<name>"),
						Action: armnetwork.WebApplicationFirewallAction("Block").ToPtr(),
						MatchConditions: []*armnetwork.MatchCondition{
							{
								MatchValues: []*string{
									to.StringPtr("192.168.1.0/24"),
									to.StringPtr("10.0.0.0/24")},
								MatchVariables: []*armnetwork.MatchVariable{
									{
										VariableName: armnetwork.WebApplicationFirewallMatchVariable("RemoteAddr").ToPtr(),
									}},
								Operator: armnetwork.WebApplicationFirewallOperator("IPMatch").ToPtr(),
							}},
						Priority: to.Int32Ptr(1),
						RuleType: armnetwork.WebApplicationFirewallRuleType("MatchRule").ToPtr(),
					},
					{
						Name:   to.StringPtr("<name>"),
						Action: armnetwork.WebApplicationFirewallAction("Block").ToPtr(),
						MatchConditions: []*armnetwork.MatchCondition{
							{
								MatchValues: []*string{
									to.StringPtr("192.168.1.0/24")},
								MatchVariables: []*armnetwork.MatchVariable{
									{
										VariableName: armnetwork.WebApplicationFirewallMatchVariable("RemoteAddr").ToPtr(),
									}},
								Operator: armnetwork.WebApplicationFirewallOperator("IPMatch").ToPtr(),
							},
							{
								MatchValues: []*string{
									to.StringPtr("Windows")},
								MatchVariables: []*armnetwork.MatchVariable{
									{
										Selector:     to.StringPtr("<selector>"),
										VariableName: armnetwork.WebApplicationFirewallMatchVariable("RequestHeaders").ToPtr(),
									}},
								Operator: armnetwork.WebApplicationFirewallOperator("Contains").ToPtr(),
							}},
						Priority: to.Int32Ptr(2),
						RuleType: armnetwork.WebApplicationFirewallRuleType("MatchRule").ToPtr(),
					}},
				ManagedRules: &armnetwork.ManagedRulesDefinition{
					Exclusions: []*armnetwork.OwaspCrsExclusionEntry{
						{
							ExclusionManagedRuleSets: []*armnetwork.ExclusionManagedRuleSet{
								{
									RuleGroups: []*armnetwork.ExclusionManagedRuleGroup{
										{
											RuleGroupName: to.StringPtr("<rule-group-name>"),
											Rules: []*armnetwork.ExclusionManagedRule{
												{
													RuleID: to.StringPtr("<rule-id>"),
												}},
										},
										{
											RuleGroupName: to.StringPtr("<rule-group-name>"),
										}},
									RuleSetType:    to.StringPtr("<rule-set-type>"),
									RuleSetVersion: to.StringPtr("<rule-set-version>"),
								}},
							MatchVariable:         armnetwork.OwaspCrsExclusionEntryMatchVariable("RequestArgNames").ToPtr(),
							Selector:              to.StringPtr("<selector>"),
							SelectorMatchOperator: armnetwork.OwaspCrsExclusionEntrySelectorMatchOperator("StartsWith").ToPtr(),
						},
						{
							ExclusionManagedRuleSets: []*armnetwork.ExclusionManagedRuleSet{
								{
									RuleGroups:     []*armnetwork.ExclusionManagedRuleGroup{},
									RuleSetType:    to.StringPtr("<rule-set-type>"),
									RuleSetVersion: to.StringPtr("<rule-set-version>"),
								}},
							MatchVariable:         armnetwork.OwaspCrsExclusionEntryMatchVariable("RequestArgNames").ToPtr(),
							Selector:              to.StringPtr("<selector>"),
							SelectorMatchOperator: armnetwork.OwaspCrsExclusionEntrySelectorMatchOperator("EndsWith").ToPtr(),
						},
						{
							MatchVariable:         armnetwork.OwaspCrsExclusionEntryMatchVariable("RequestArgNames").ToPtr(),
							Selector:              to.StringPtr("<selector>"),
							SelectorMatchOperator: armnetwork.OwaspCrsExclusionEntrySelectorMatchOperator("StartsWith").ToPtr(),
						},
						{
							MatchVariable:         armnetwork.OwaspCrsExclusionEntryMatchVariable("RequestArgValues").ToPtr(),
							Selector:              to.StringPtr("<selector>"),
							SelectorMatchOperator: armnetwork.OwaspCrsExclusionEntrySelectorMatchOperator("StartsWith").ToPtr(),
						}},
					ManagedRuleSets: []*armnetwork.ManagedRuleSet{
						{
							RuleSetType:    to.StringPtr("<rule-set-type>"),
							RuleSetVersion: to.StringPtr("<rule-set-version>"),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebApplicationFirewallPoliciesClientCreateOrUpdateResult)
}
```
