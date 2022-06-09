```go
package armfrontdoor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-11-01/examples/WafPolicyCreateOrUpdate.json
func ExamplePoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"Policy1",
		armfrontdoor.WebApplicationFirewallPolicy{
			Properties: &armfrontdoor.WebApplicationFirewallPolicyProperties{
				CustomRules: &armfrontdoor.CustomRuleList{
					Rules: []*armfrontdoor.CustomRule{
						{
							Name:   to.Ptr("Rule1"),
							Action: to.Ptr(armfrontdoor.ActionTypeBlock),
							MatchConditions: []*armfrontdoor.MatchCondition{
								{
									MatchValue: []*string{
										to.Ptr("192.168.1.0/24"),
										to.Ptr("10.0.0.0/24")},
									MatchVariable: to.Ptr(armfrontdoor.MatchVariableRemoteAddr),
									Operator:      to.Ptr(armfrontdoor.OperatorIPMatch),
								}},
							Priority:           to.Ptr[int32](1),
							RateLimitThreshold: to.Ptr[int32](1000),
							RuleType:           to.Ptr(armfrontdoor.RuleTypeRateLimitRule),
						},
						{
							Name:   to.Ptr("Rule2"),
							Action: to.Ptr(armfrontdoor.ActionTypeBlock),
							MatchConditions: []*armfrontdoor.MatchCondition{
								{
									MatchValue: []*string{
										to.Ptr("CH")},
									MatchVariable: to.Ptr(armfrontdoor.MatchVariableRemoteAddr),
									Operator:      to.Ptr(armfrontdoor.OperatorGeoMatch),
								},
								{
									MatchValue: []*string{
										to.Ptr("windows")},
									MatchVariable: to.Ptr(armfrontdoor.MatchVariableRequestHeader),
									Operator:      to.Ptr(armfrontdoor.OperatorContains),
									Selector:      to.Ptr("UserAgent"),
									Transforms: []*armfrontdoor.TransformType{
										to.Ptr(armfrontdoor.TransformTypeLowercase)},
								}},
							Priority: to.Ptr[int32](2),
							RuleType: to.Ptr(armfrontdoor.RuleTypeMatchRule),
						}},
				},
				ManagedRules: &armfrontdoor.ManagedRuleSetList{
					ManagedRuleSets: []*armfrontdoor.ManagedRuleSet{
						{
							Exclusions: []*armfrontdoor.ManagedRuleExclusion{
								{
									MatchVariable:         to.Ptr(armfrontdoor.ManagedRuleExclusionMatchVariableRequestHeaderNames),
									Selector:              to.Ptr("User-Agent"),
									SelectorMatchOperator: to.Ptr(armfrontdoor.ManagedRuleExclusionSelectorMatchOperatorEquals),
								}},
							RuleGroupOverrides: []*armfrontdoor.ManagedRuleGroupOverride{
								{
									Exclusions: []*armfrontdoor.ManagedRuleExclusion{
										{
											MatchVariable:         to.Ptr(armfrontdoor.ManagedRuleExclusionMatchVariableRequestCookieNames),
											Selector:              to.Ptr("token"),
											SelectorMatchOperator: to.Ptr(armfrontdoor.ManagedRuleExclusionSelectorMatchOperatorStartsWith),
										}},
									RuleGroupName: to.Ptr("SQLI"),
									Rules: []*armfrontdoor.ManagedRuleOverride{
										{
											Action:       to.Ptr(armfrontdoor.ActionTypeRedirect),
											EnabledState: to.Ptr(armfrontdoor.ManagedRuleEnabledStateEnabled),
											Exclusions: []*armfrontdoor.ManagedRuleExclusion{
												{
													MatchVariable:         to.Ptr(armfrontdoor.ManagedRuleExclusionMatchVariableQueryStringArgNames),
													Selector:              to.Ptr("query"),
													SelectorMatchOperator: to.Ptr(armfrontdoor.ManagedRuleExclusionSelectorMatchOperatorEquals),
												}},
											RuleID: to.Ptr("942100"),
										},
										{
											EnabledState: to.Ptr(armfrontdoor.ManagedRuleEnabledStateDisabled),
											RuleID:       to.Ptr("942110"),
										}},
								}},
							RuleSetAction:  to.Ptr(armfrontdoor.ManagedRuleSetActionTypeBlock),
							RuleSetType:    to.Ptr("DefaultRuleSet"),
							RuleSetVersion: to.Ptr("1.0"),
						}},
				},
				PolicySettings: &armfrontdoor.PolicySettings{
					CustomBlockResponseBody:       to.Ptr("PGh0bWw+CjxoZWFkZXI+PHRpdGxlPkhlbGxvPC90aXRsZT48L2hlYWRlcj4KPGJvZHk+CkhlbGxvIHdvcmxkCjwvYm9keT4KPC9odG1sPg=="),
					CustomBlockResponseStatusCode: to.Ptr[int32](499),
					EnabledState:                  to.Ptr(armfrontdoor.PolicyEnabledStateEnabled),
					Mode:                          to.Ptr(armfrontdoor.PolicyModePrevention),
					RedirectURL:                   to.Ptr("http://www.bing.com"),
					RequestBodyCheck:              to.Ptr(armfrontdoor.PolicyRequestBodyCheckDisabled),
				},
			},
			SKU: &armfrontdoor.SKU{
				Name: to.Ptr(armfrontdoor.SKUNameClassicAzureFrontDoor),
			},
		},
		nil)
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ffrontdoor%2Farmfrontdoor%2Fv1.0.0/sdk/resourcemanager/frontdoor/armfrontdoor/README.md) on how to add the SDK to your project and authenticate.
