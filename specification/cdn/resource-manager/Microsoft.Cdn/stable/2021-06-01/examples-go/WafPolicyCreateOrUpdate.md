Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcdn%2Farmcdn%2Fv0.5.0/sdk/resourcemanager/cdn/armcdn/README.md) on how to add the SDK to your project and authenticate.

```go
package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPolicyCreateOrUpdate.json
func ExamplePoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcdn.NewPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<policy-name>",
		armcdn.WebApplicationFirewallPolicy{
			Location: to.Ptr("<location>"),
			Properties: &armcdn.WebApplicationFirewallPolicyProperties{
				CustomRules: &armcdn.CustomRuleList{
					Rules: []*armcdn.CustomRule{
						{
							Name:         to.Ptr("<name>"),
							Action:       to.Ptr(armcdn.ActionTypeBlock),
							EnabledState: to.Ptr(armcdn.CustomRuleEnabledStateEnabled),
							MatchConditions: []*armcdn.MatchCondition{
								{
									MatchValue: []*string{
										to.Ptr("CH")},
									MatchVariable:   to.Ptr(armcdn.WafMatchVariableRemoteAddr),
									NegateCondition: to.Ptr(false),
									Operator:        to.Ptr(armcdn.OperatorGeoMatch),
									Transforms:      []*armcdn.TransformType{},
								},
								{
									MatchValue: []*string{
										to.Ptr("windows")},
									MatchVariable:   to.Ptr(armcdn.WafMatchVariableRequestHeader),
									NegateCondition: to.Ptr(false),
									Operator:        to.Ptr(armcdn.OperatorContains),
									Selector:        to.Ptr("<selector>"),
									Transforms:      []*armcdn.TransformType{},
								},
								{
									MatchValue: []*string{
										to.Ptr("<?php"),
										to.Ptr("?>")},
									MatchVariable:   to.Ptr(armcdn.WafMatchVariableQueryString),
									NegateCondition: to.Ptr(false),
									Operator:        to.Ptr(armcdn.OperatorContains),
									Selector:        to.Ptr("<selector>"),
									Transforms: []*armcdn.TransformType{
										to.Ptr(armcdn.TransformTypeURLDecode),
										to.Ptr(armcdn.TransformTypeLowercase)},
								}},
							Priority: to.Ptr[int32](2),
						}},
				},
				ManagedRules: &armcdn.ManagedRuleSetList{
					ManagedRuleSets: []*armcdn.ManagedRuleSet{
						{
							RuleGroupOverrides: []*armcdn.ManagedRuleGroupOverride{
								{
									RuleGroupName: to.Ptr("<rule-group-name>"),
									Rules: []*armcdn.ManagedRuleOverride{
										{
											Action:       to.Ptr(armcdn.ActionTypeRedirect),
											EnabledState: to.Ptr(armcdn.ManagedRuleEnabledStateEnabled),
											RuleID:       to.Ptr("<rule-id>"),
										},
										{
											EnabledState: to.Ptr(armcdn.ManagedRuleEnabledStateDisabled),
											RuleID:       to.Ptr("<rule-id>"),
										}},
								}},
							RuleSetType:    to.Ptr("<rule-set-type>"),
							RuleSetVersion: to.Ptr("<rule-set-version>"),
						}},
				},
				PolicySettings: &armcdn.PolicySettings{
					DefaultCustomBlockResponseBody:       to.Ptr("<default-custom-block-response-body>"),
					DefaultCustomBlockResponseStatusCode: to.Ptr(armcdn.PolicySettingsDefaultCustomBlockResponseStatusCode(200)),
					DefaultRedirectURL:                   to.Ptr("<default-redirect-url>"),
				},
				RateLimitRules: &armcdn.RateLimitRuleList{
					Rules: []*armcdn.RateLimitRule{
						{
							Name:         to.Ptr("<name>"),
							Action:       to.Ptr(armcdn.ActionTypeBlock),
							EnabledState: to.Ptr(armcdn.CustomRuleEnabledStateEnabled),
							MatchConditions: []*armcdn.MatchCondition{
								{
									MatchValue: []*string{
										to.Ptr("192.168.1.0/24"),
										to.Ptr("10.0.0.0/24")},
									MatchVariable:   to.Ptr(armcdn.WafMatchVariableRemoteAddr),
									NegateCondition: to.Ptr(false),
									Operator:        to.Ptr(armcdn.OperatorIPMatch),
									Transforms:      []*armcdn.TransformType{},
								}},
							Priority:                   to.Ptr[int32](1),
							RateLimitDurationInMinutes: to.Ptr[int32](0),
							RateLimitThreshold:         to.Ptr[int32](1000),
						}},
				},
			},
			SKU: &armcdn.SKU{
				Name: to.Ptr(armcdn.SKUNameStandardMicrosoft),
			},
		},
		&armcdn.PoliciesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
