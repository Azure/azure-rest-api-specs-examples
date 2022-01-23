Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcdn%2Farmcdn%2Fv0.3.0/sdk/resourcemanager/cdn/armcdn/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPolicyCreateOrUpdate.json
func ExamplePoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewPoliciesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<policy-name>",
		armcdn.WebApplicationFirewallPolicy{
			Location: to.StringPtr("<location>"),
			Properties: &armcdn.WebApplicationFirewallPolicyProperties{
				CustomRules: &armcdn.CustomRuleList{
					Rules: []*armcdn.CustomRule{
						{
							Name:         to.StringPtr("<name>"),
							Action:       armcdn.ActionType("Block").ToPtr(),
							EnabledState: armcdn.CustomRuleEnabledState("Enabled").ToPtr(),
							MatchConditions: []*armcdn.MatchCondition{
								{
									MatchValue: []*string{
										to.StringPtr("CH")},
									MatchVariable:   armcdn.WafMatchVariable("RemoteAddr").ToPtr(),
									NegateCondition: to.BoolPtr(false),
									Operator:        armcdn.Operator("GeoMatch").ToPtr(),
									Transforms:      []*armcdn.TransformType{},
								},
								{
									MatchValue: []*string{
										to.StringPtr("windows")},
									MatchVariable:   armcdn.WafMatchVariable("RequestHeader").ToPtr(),
									NegateCondition: to.BoolPtr(false),
									Operator:        armcdn.Operator("Contains").ToPtr(),
									Selector:        to.StringPtr("<selector>"),
									Transforms:      []*armcdn.TransformType{},
								},
								{
									MatchValue: []*string{
										to.StringPtr("<?php"),
										to.StringPtr("?>")},
									MatchVariable:   armcdn.WafMatchVariable("QueryString").ToPtr(),
									NegateCondition: to.BoolPtr(false),
									Operator:        armcdn.Operator("Contains").ToPtr(),
									Selector:        to.StringPtr("<selector>"),
									Transforms: []*armcdn.TransformType{
										armcdn.TransformType("UrlDecode").ToPtr(),
										armcdn.TransformType("Lowercase").ToPtr()},
								}},
							Priority: to.Int32Ptr(2),
						}},
				},
				ManagedRules: &armcdn.ManagedRuleSetList{
					ManagedRuleSets: []*armcdn.ManagedRuleSet{
						{
							RuleGroupOverrides: []*armcdn.ManagedRuleGroupOverride{
								{
									RuleGroupName: to.StringPtr("<rule-group-name>"),
									Rules: []*armcdn.ManagedRuleOverride{
										{
											Action:       armcdn.ActionType("Redirect").ToPtr(),
											EnabledState: armcdn.ManagedRuleEnabledState("Enabled").ToPtr(),
											RuleID:       to.StringPtr("<rule-id>"),
										},
										{
											EnabledState: armcdn.ManagedRuleEnabledState("Disabled").ToPtr(),
											RuleID:       to.StringPtr("<rule-id>"),
										}},
								}},
							RuleSetType:    to.StringPtr("<rule-set-type>"),
							RuleSetVersion: to.StringPtr("<rule-set-version>"),
						}},
				},
				PolicySettings: &armcdn.PolicySettings{
					DefaultCustomBlockResponseBody:       to.StringPtr("<default-custom-block-response-body>"),
					DefaultCustomBlockResponseStatusCode: armcdn.PolicySettingsDefaultCustomBlockResponseStatusCode(200).ToPtr(),
					DefaultRedirectURL:                   to.StringPtr("<default-redirect-url>"),
				},
				RateLimitRules: &armcdn.RateLimitRuleList{
					Rules: []*armcdn.RateLimitRule{
						{
							Name:         to.StringPtr("<name>"),
							Action:       armcdn.ActionType("Block").ToPtr(),
							EnabledState: armcdn.CustomRuleEnabledState("Enabled").ToPtr(),
							MatchConditions: []*armcdn.MatchCondition{
								{
									MatchValue: []*string{
										to.StringPtr("192.168.1.0/24"),
										to.StringPtr("10.0.0.0/24")},
									MatchVariable:   armcdn.WafMatchVariable("RemoteAddr").ToPtr(),
									NegateCondition: to.BoolPtr(false),
									Operator:        armcdn.Operator("IPMatch").ToPtr(),
									Transforms:      []*armcdn.TransformType{},
								}},
							Priority:                   to.Int32Ptr(1),
							RateLimitDurationInMinutes: to.Int32Ptr(0),
							RateLimitThreshold:         to.Int32Ptr(1000),
						}},
				},
			},
			SKU: &armcdn.SKU{
				Name: armcdn.SKUName("Standard_Microsoft").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.PoliciesClientCreateOrUpdateResult)
}
```
