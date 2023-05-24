package armfrontdoor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/frontdoor/resource-manager/Microsoft.Network/stable/2022-05-01/examples/WafPolicyGet.json
func ExamplePoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfrontdoor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoliciesClient().Get(ctx, "rg1", "Policy1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WebApplicationFirewallPolicy = armfrontdoor.WebApplicationFirewallPolicy{
	// 	Name: to.Ptr("Policy1"),
	// 	Type: to.Ptr("Microsoft.Network/frontdoorwebapplicationfirewallpolicies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/FrontDoorWebApplicationFirewallPolicies/Policy1"),
	// 	Location: to.Ptr("WestUs"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armfrontdoor.WebApplicationFirewallPolicyProperties{
	// 		CustomRules: &armfrontdoor.CustomRuleList{
	// 			Rules: []*armfrontdoor.CustomRule{
	// 				{
	// 					Name: to.Ptr("Rule1"),
	// 					Action: to.Ptr(armfrontdoor.ActionTypeBlock),
	// 					EnabledState: to.Ptr(armfrontdoor.CustomRuleEnabledStateEnabled),
	// 					MatchConditions: []*armfrontdoor.MatchCondition{
	// 						{
	// 							MatchValue: []*string{
	// 								to.Ptr("192.168.1.0/24"),
	// 								to.Ptr("10.0.0.0/24")},
	// 								MatchVariable: to.Ptr(armfrontdoor.MatchVariableRemoteAddr),
	// 								NegateCondition: to.Ptr(false),
	// 								Operator: to.Ptr(armfrontdoor.OperatorIPMatch),
	// 								Transforms: []*armfrontdoor.TransformType{
	// 								},
	// 						}},
	// 						Priority: to.Ptr[int32](1),
	// 						RateLimitDurationInMinutes: to.Ptr[int32](0),
	// 						RateLimitThreshold: to.Ptr[int32](1000),
	// 						RuleType: to.Ptr(armfrontdoor.RuleTypeRateLimitRule),
	// 					},
	// 					{
	// 						Name: to.Ptr("Rule2"),
	// 						Action: to.Ptr(armfrontdoor.ActionTypeBlock),
	// 						EnabledState: to.Ptr(armfrontdoor.CustomRuleEnabledStateEnabled),
	// 						MatchConditions: []*armfrontdoor.MatchCondition{
	// 							{
	// 								MatchValue: []*string{
	// 									to.Ptr("CH")},
	// 									MatchVariable: to.Ptr(armfrontdoor.MatchVariableRemoteAddr),
	// 									NegateCondition: to.Ptr(false),
	// 									Operator: to.Ptr(armfrontdoor.OperatorGeoMatch),
	// 								},
	// 								{
	// 									MatchValue: []*string{
	// 										to.Ptr("windows")},
	// 										MatchVariable: to.Ptr(armfrontdoor.MatchVariableRequestHeader),
	// 										NegateCondition: to.Ptr(false),
	// 										Operator: to.Ptr(armfrontdoor.OperatorContains),
	// 										Selector: to.Ptr("UserAgent"),
	// 										Transforms: []*armfrontdoor.TransformType{
	// 											to.Ptr(armfrontdoor.TransformTypeLowercase)},
	// 									}},
	// 									Priority: to.Ptr[int32](2),
	// 									RateLimitDurationInMinutes: to.Ptr[int32](0),
	// 									RateLimitThreshold: to.Ptr[int32](0),
	// 									RuleType: to.Ptr(armfrontdoor.RuleTypeMatchRule),
	// 							}},
	// 						},
	// 						FrontendEndpointLinks: []*armfrontdoor.FrontendEndpointLink{
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontdoors/fd1/frontendendpoints/fd1-azurefd-net"),
	// 						}},
	// 						ManagedRules: &armfrontdoor.ManagedRuleSetList{
	// 							ManagedRuleSets: []*armfrontdoor.ManagedRuleSet{
	// 								{
	// 									Exclusions: []*armfrontdoor.ManagedRuleExclusion{
	// 										{
	// 											MatchVariable: to.Ptr(armfrontdoor.ManagedRuleExclusionMatchVariableRequestHeaderNames),
	// 											Selector: to.Ptr("User-Agent"),
	// 											SelectorMatchOperator: to.Ptr(armfrontdoor.ManagedRuleExclusionSelectorMatchOperatorEquals),
	// 									}},
	// 									RuleGroupOverrides: []*armfrontdoor.ManagedRuleGroupOverride{
	// 										{
	// 											Exclusions: []*armfrontdoor.ManagedRuleExclusion{
	// 											},
	// 											RuleGroupName: to.Ptr("SQLI"),
	// 											Rules: []*armfrontdoor.ManagedRuleOverride{
	// 												{
	// 													Action: to.Ptr(armfrontdoor.ActionTypeRedirect),
	// 													EnabledState: to.Ptr(armfrontdoor.ManagedRuleEnabledStateEnabled),
	// 													Exclusions: []*armfrontdoor.ManagedRuleExclusion{
	// 													},
	// 													RuleID: to.Ptr("942100"),
	// 												},
	// 												{
	// 													EnabledState: to.Ptr(armfrontdoor.ManagedRuleEnabledStateDisabled),
	// 													RuleID: to.Ptr("942110"),
	// 											}},
	// 									}},
	// 									RuleSetAction: to.Ptr(armfrontdoor.ManagedRuleSetActionTypeBlock),
	// 									RuleSetType: to.Ptr("DefaultRuleSet"),
	// 									RuleSetVersion: to.Ptr("1.0"),
	// 							}},
	// 						},
	// 						PolicySettings: &armfrontdoor.PolicySettings{
	// 							CustomBlockResponseBody: to.Ptr("PGh0bWw+CjxoZWFkZXI+PHRpdGxlPkhlbGxvPC90aXRsZT48L2hlYWRlcj4KPGJvZHk+CkhlbGxvIHdvcmxkCjwvYm9keT4KPC9odG1sPg=="),
	// 							CustomBlockResponseStatusCode: to.Ptr[int32](499),
	// 							EnabledState: to.Ptr(armfrontdoor.PolicyEnabledStateEnabled),
	// 							Mode: to.Ptr(armfrontdoor.PolicyModePrevention),
	// 							RedirectURL: to.Ptr("http://www.bing.com"),
	// 							RequestBodyCheck: to.Ptr(armfrontdoor.PolicyRequestBodyCheckDisabled),
	// 						},
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 						ResourceState: to.Ptr(armfrontdoor.PolicyResourceStateEnabled),
	// 						SecurityPolicyLinks: []*armfrontdoor.SecurityPolicyLink{
	// 						},
	// 					},
	// 					SKU: &armfrontdoor.SKU{
	// 						Name: to.Ptr(armfrontdoor.SKUNameClassicAzureFrontDoor),
	// 					},
	// 				}
}
