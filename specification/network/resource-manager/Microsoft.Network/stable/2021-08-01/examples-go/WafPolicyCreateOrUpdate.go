package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/WafPolicyCreateOrUpdate.json
func ExampleWebApplicationFirewallPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWebApplicationFirewallPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg1",
		"Policy1",
		armnetwork.WebApplicationFirewallPolicy{
			Location: to.Ptr("WestUs"),
			Properties: &armnetwork.WebApplicationFirewallPolicyPropertiesFormat{
				CustomRules: []*armnetwork.WebApplicationFirewallCustomRule{
					{
						Name:   to.Ptr("Rule1"),
						Action: to.Ptr(armnetwork.WebApplicationFirewallActionBlock),
						MatchConditions: []*armnetwork.MatchCondition{
							{
								MatchValues: []*string{
									to.Ptr("192.168.1.0/24"),
									to.Ptr("10.0.0.0/24")},
								MatchVariables: []*armnetwork.MatchVariable{
									{
										VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRemoteAddr),
									}},
								Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorIPMatch),
							}},
						Priority: to.Ptr[int32](1),
						RuleType: to.Ptr(armnetwork.WebApplicationFirewallRuleTypeMatchRule),
					},
					{
						Name:   to.Ptr("Rule2"),
						Action: to.Ptr(armnetwork.WebApplicationFirewallActionBlock),
						MatchConditions: []*armnetwork.MatchCondition{
							{
								MatchValues: []*string{
									to.Ptr("192.168.1.0/24")},
								MatchVariables: []*armnetwork.MatchVariable{
									{
										VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRemoteAddr),
									}},
								Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorIPMatch),
							},
							{
								MatchValues: []*string{
									to.Ptr("Windows")},
								MatchVariables: []*armnetwork.MatchVariable{
									{
										Selector:     to.Ptr("UserAgent"),
										VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRequestHeaders),
									}},
								Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorContains),
							}},
						Priority: to.Ptr[int32](2),
						RuleType: to.Ptr(armnetwork.WebApplicationFirewallRuleTypeMatchRule),
					}},
				ManagedRules: &armnetwork.ManagedRulesDefinition{
					Exclusions: []*armnetwork.OwaspCrsExclusionEntry{
						{
							ExclusionManagedRuleSets: []*armnetwork.ExclusionManagedRuleSet{
								{
									RuleGroups: []*armnetwork.ExclusionManagedRuleGroup{
										{
											RuleGroupName: to.Ptr("REQUEST-930-APPLICATION-ATTACK-LFI"),
											Rules: []*armnetwork.ExclusionManagedRule{
												{
													RuleID: to.Ptr("930120"),
												}},
										},
										{
											RuleGroupName: to.Ptr("REQUEST-932-APPLICATION-ATTACK-RCE"),
										}},
									RuleSetType:    to.Ptr("OWASP"),
									RuleSetVersion: to.Ptr("3.2"),
								}},
							MatchVariable:         to.Ptr(armnetwork.OwaspCrsExclusionEntryMatchVariableRequestArgNames),
							Selector:              to.Ptr("hello"),
							SelectorMatchOperator: to.Ptr(armnetwork.OwaspCrsExclusionEntrySelectorMatchOperatorStartsWith),
						},
						{
							ExclusionManagedRuleSets: []*armnetwork.ExclusionManagedRuleSet{
								{
									RuleGroups:     []*armnetwork.ExclusionManagedRuleGroup{},
									RuleSetType:    to.Ptr("OWASP"),
									RuleSetVersion: to.Ptr("3.1"),
								}},
							MatchVariable:         to.Ptr(armnetwork.OwaspCrsExclusionEntryMatchVariableRequestArgNames),
							Selector:              to.Ptr("hello"),
							SelectorMatchOperator: to.Ptr(armnetwork.OwaspCrsExclusionEntrySelectorMatchOperatorEndsWith),
						},
						{
							MatchVariable:         to.Ptr(armnetwork.OwaspCrsExclusionEntryMatchVariableRequestArgNames),
							Selector:              to.Ptr("test"),
							SelectorMatchOperator: to.Ptr(armnetwork.OwaspCrsExclusionEntrySelectorMatchOperatorStartsWith),
						},
						{
							MatchVariable:         to.Ptr(armnetwork.OwaspCrsExclusionEntryMatchVariableRequestArgValues),
							Selector:              to.Ptr("test"),
							SelectorMatchOperator: to.Ptr(armnetwork.OwaspCrsExclusionEntrySelectorMatchOperatorStartsWith),
						}},
					ManagedRuleSets: []*armnetwork.ManagedRuleSet{
						{
							RuleSetType:    to.Ptr("OWASP"),
							RuleSetVersion: to.Ptr("3.2"),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
