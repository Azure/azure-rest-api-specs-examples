package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7b551033155a63739b6d28f79b9c07569f6179b8/specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/WafPatchPolicy.json
func ExamplePoliciesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPoliciesClient().BeginUpdate(ctx, "rg1", "MicrosoftCdnWafPolicy", armcdn.WebApplicationFirewallPolicyPatchParameters{
		Tags: map[string]*string{
			"foo": to.Ptr("bar"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WebApplicationFirewallPolicy = armcdn.WebApplicationFirewallPolicy{
	// 	Name: to.Ptr("MicrosoftCdnWafPolicy"),
	// 	Type: to.Ptr("Microsoft.Cdn/cdnwebapplicationfirewallpolicies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cdn/CdnWebApplicationFirewallPolicies/MicrosoftCdnWafPolicy"),
	// 	Location: to.Ptr("WestUs"),
	// 	Tags: map[string]*string{
	// 		"foo": to.Ptr("bar"),
	// 	},
	// 	Properties: &armcdn.WebApplicationFirewallPolicyProperties{
	// 		CustomRules: &armcdn.CustomRuleList{
	// 			Rules: []*armcdn.CustomRule{
	// 				{
	// 					Name: to.Ptr("CustomRule1"),
	// 					Action: to.Ptr(armcdn.ActionTypeBlock),
	// 					EnabledState: to.Ptr(armcdn.CustomRuleEnabledStateEnabled),
	// 					MatchConditions: []*armcdn.MatchCondition{
	// 						{
	// 							MatchValue: []*string{
	// 								to.Ptr("CH")},
	// 								MatchVariable: to.Ptr(armcdn.WafMatchVariableRemoteAddr),
	// 								NegateCondition: to.Ptr(false),
	// 								Operator: to.Ptr(armcdn.OperatorGeoMatch),
	// 								Transforms: []*armcdn.TransformType{
	// 								},
	// 							},
	// 							{
	// 								MatchValue: []*string{
	// 									to.Ptr("windows")},
	// 									MatchVariable: to.Ptr(armcdn.WafMatchVariableRequestHeader),
	// 									NegateCondition: to.Ptr(false),
	// 									Operator: to.Ptr(armcdn.OperatorContains),
	// 									Selector: to.Ptr("UserAgent"),
	// 									Transforms: []*armcdn.TransformType{
	// 									},
	// 								},
	// 								{
	// 									MatchValue: []*string{
	// 										to.Ptr("<?php"),
	// 										to.Ptr("?>")},
	// 										MatchVariable: to.Ptr(armcdn.WafMatchVariableQueryString),
	// 										NegateCondition: to.Ptr(false),
	// 										Operator: to.Ptr(armcdn.OperatorContains),
	// 										Selector: to.Ptr("search"),
	// 										Transforms: []*armcdn.TransformType{
	// 											to.Ptr(armcdn.TransformTypeURLDecode),
	// 											to.Ptr(armcdn.TransformTypeLowercase)},
	// 									}},
	// 									Priority: to.Ptr[int32](2),
	// 							}},
	// 						},
	// 						EndpointLinks: []*armcdn.LinkedEndpoint{
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cdn/profiles/profile1/endpoints/testEndpoint1"),
	// 							},
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cdn/profiles/profile1/endpoints/testEndpoint2"),
	// 						}},
	// 						ManagedRules: &armcdn.ManagedRuleSetList{
	// 							ManagedRuleSets: []*armcdn.ManagedRuleSet{
	// 								{
	// 									RuleGroupOverrides: []*armcdn.ManagedRuleGroupOverride{
	// 										{
	// 											RuleGroupName: to.Ptr("Group1"),
	// 											Rules: []*armcdn.ManagedRuleOverride{
	// 												{
	// 													Action: to.Ptr(armcdn.ActionTypeRedirect),
	// 													EnabledState: to.Ptr(armcdn.ManagedRuleEnabledStateEnabled),
	// 													RuleID: to.Ptr("GROUP1-0001"),
	// 												},
	// 												{
	// 													EnabledState: to.Ptr(armcdn.ManagedRuleEnabledStateDisabled),
	// 													RuleID: to.Ptr("GROUP1-0002"),
	// 											}},
	// 									}},
	// 									RuleSetType: to.Ptr("DefaultRuleSet"),
	// 									RuleSetVersion: to.Ptr("preview-1.0"),
	// 							}},
	// 						},
	// 						PolicySettings: &armcdn.PolicySettings{
	// 							DefaultCustomBlockResponseBody: to.Ptr("PGh0bWw+CjxoZWFkZXI+PHRpdGxlPkhlbGxvPC90aXRsZT48L2hlYWRlcj4KPGJvZHk+CkhlbGxvIHdvcmxkCjwvYm9keT4KPC9odG1sPg=="),
	// 							DefaultCustomBlockResponseStatusCode: to.Ptr(armcdn.PolicySettingsDefaultCustomBlockResponseStatusCode(403)),
	// 							DefaultRedirectURL: to.Ptr("http://www.bing.com"),
	// 							EnabledState: to.Ptr(armcdn.PolicyEnabledStateEnabled),
	// 							Mode: to.Ptr(armcdn.PolicyModePrevention),
	// 						},
	// 						ProvisioningState: to.Ptr(armcdn.ProvisioningStateSucceeded),
	// 						RateLimitRules: &armcdn.RateLimitRuleList{
	// 							Rules: []*armcdn.RateLimitRule{
	// 								{
	// 									Name: to.Ptr("RateLimitRule1"),
	// 									Action: to.Ptr(armcdn.ActionTypeBlock),
	// 									EnabledState: to.Ptr(armcdn.CustomRuleEnabledStateEnabled),
	// 									MatchConditions: []*armcdn.MatchCondition{
	// 										{
	// 											MatchValue: []*string{
	// 												to.Ptr("192.168.1.0/24"),
	// 												to.Ptr("10.0.0.0/24")},
	// 												MatchVariable: to.Ptr(armcdn.WafMatchVariableRemoteAddr),
	// 												NegateCondition: to.Ptr(false),
	// 												Operator: to.Ptr(armcdn.OperatorIPMatch),
	// 												Transforms: []*armcdn.TransformType{
	// 												},
	// 										}},
	// 										Priority: to.Ptr[int32](1),
	// 										RateLimitDurationInMinutes: to.Ptr[int32](0),
	// 										RateLimitThreshold: to.Ptr[int32](1000),
	// 								}},
	// 							},
	// 							ResourceState: to.Ptr(armcdn.PolicyResourceStateEnabled),
	// 						},
	// 						SKU: &armcdn.SKU{
	// 							Name: to.Ptr(armcdn.SKUNameStandardMicrosoft),
	// 						},
	// 					}
}
