Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ffrontdoor%2Farmfrontdoor%2Fv0.1.0/sdk/resourcemanager/frontdoor/armfrontdoor/README.md) on how to add the SDK to your project and authenticate.

```go
package armfrontdoor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-11-01/examples/WafPolicyCreateOrUpdate.json
func ExamplePoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewPoliciesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<policy-name>",
		armfrontdoor.WebApplicationFirewallPolicy{
			Properties: &armfrontdoor.WebApplicationFirewallPolicyProperties{
				CustomRules: &armfrontdoor.CustomRuleList{
					Rules: []*armfrontdoor.CustomRule{
						{
							Name:   to.StringPtr("<name>"),
							Action: armfrontdoor.ActionTypeBlock.ToPtr(),
							MatchConditions: []*armfrontdoor.MatchCondition{
								{
									MatchValue: []*string{
										to.StringPtr("192.168.1.0/24"),
										to.StringPtr("10.0.0.0/24")},
									MatchVariable: armfrontdoor.MatchVariableRemoteAddr.ToPtr(),
									Operator:      armfrontdoor.OperatorIPMatch.ToPtr(),
								}},
							Priority:           to.Int32Ptr(1),
							RateLimitThreshold: to.Int32Ptr(1000),
							RuleType:           armfrontdoor.RuleTypeRateLimitRule.ToPtr(),
						},
						{
							Name:   to.StringPtr("<name>"),
							Action: armfrontdoor.ActionTypeBlock.ToPtr(),
							MatchConditions: []*armfrontdoor.MatchCondition{
								{
									MatchValue: []*string{
										to.StringPtr("CH")},
									MatchVariable: armfrontdoor.MatchVariableRemoteAddr.ToPtr(),
									Operator:      armfrontdoor.OperatorGeoMatch.ToPtr(),
								},
								{
									MatchValue: []*string{
										to.StringPtr("windows")},
									MatchVariable: armfrontdoor.MatchVariableRequestHeader.ToPtr(),
									Operator:      armfrontdoor.OperatorContains.ToPtr(),
									Selector:      to.StringPtr("<selector>"),
									Transforms: []*armfrontdoor.TransformType{
										armfrontdoor.TransformTypeLowercase.ToPtr()},
								}},
							Priority: to.Int32Ptr(2),
							RuleType: armfrontdoor.RuleTypeMatchRule.ToPtr(),
						}},
				},
				ManagedRules: &armfrontdoor.ManagedRuleSetList{
					ManagedRuleSets: []*armfrontdoor.ManagedRuleSet{
						{
							Exclusions: []*armfrontdoor.ManagedRuleExclusion{
								{
									MatchVariable:         armfrontdoor.ManagedRuleExclusionMatchVariableRequestHeaderNames.ToPtr(),
									Selector:              to.StringPtr("<selector>"),
									SelectorMatchOperator: armfrontdoor.ManagedRuleExclusionSelectorMatchOperatorEquals.ToPtr(),
								}},
							RuleGroupOverrides: []*armfrontdoor.ManagedRuleGroupOverride{
								{
									Exclusions: []*armfrontdoor.ManagedRuleExclusion{
										{
											MatchVariable:         armfrontdoor.ManagedRuleExclusionMatchVariableRequestCookieNames.ToPtr(),
											Selector:              to.StringPtr("<selector>"),
											SelectorMatchOperator: armfrontdoor.ManagedRuleExclusionSelectorMatchOperatorStartsWith.ToPtr(),
										}},
									RuleGroupName: to.StringPtr("<rule-group-name>"),
									Rules: []*armfrontdoor.ManagedRuleOverride{
										{
											Action:       armfrontdoor.ActionTypeRedirect.ToPtr(),
											EnabledState: armfrontdoor.ManagedRuleEnabledStateEnabled.ToPtr(),
											Exclusions: []*armfrontdoor.ManagedRuleExclusion{
												{
													MatchVariable:         armfrontdoor.ManagedRuleExclusionMatchVariableQueryStringArgNames.ToPtr(),
													Selector:              to.StringPtr("<selector>"),
													SelectorMatchOperator: armfrontdoor.ManagedRuleExclusionSelectorMatchOperatorEquals.ToPtr(),
												}},
											RuleID: to.StringPtr("<rule-id>"),
										},
										{
											EnabledState: armfrontdoor.ManagedRuleEnabledStateDisabled.ToPtr(),
											RuleID:       to.StringPtr("<rule-id>"),
										}},
								}},
							RuleSetAction:  armfrontdoor.ManagedRuleSetActionTypeBlock.ToPtr(),
							RuleSetType:    to.StringPtr("<rule-set-type>"),
							RuleSetVersion: to.StringPtr("<rule-set-version>"),
						}},
				},
				PolicySettings: &armfrontdoor.PolicySettings{
					CustomBlockResponseBody:       to.StringPtr("<custom-block-response-body>"),
					CustomBlockResponseStatusCode: to.Int32Ptr(499),
					EnabledState:                  armfrontdoor.PolicyEnabledStateEnabled.ToPtr(),
					Mode:                          armfrontdoor.PolicyModePrevention.ToPtr(),
					RedirectURL:                   to.StringPtr("<redirect-url>"),
					RequestBodyCheck:              armfrontdoor.PolicyRequestBodyCheckDisabled.ToPtr(),
				},
			},
			SKU: &armfrontdoor.SKU{
				Name: armfrontdoor.SKUNameClassicAzureFrontDoor.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WebApplicationFirewallPolicy.ID: %s\n", *res.ID)
}
```
