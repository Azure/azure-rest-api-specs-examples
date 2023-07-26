package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9975d3476c05bcc6bd9535ad3dfb564e6a168fa5/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/WafListAllPolicies.json
func ExampleWebApplicationFirewallPoliciesClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebApplicationFirewallPoliciesClient().NewListAllPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.WebApplicationFirewallPolicyListResult = armnetwork.WebApplicationFirewallPolicyListResult{
		// 	Value: []*armnetwork.WebApplicationFirewallPolicy{
		// 		{
		// 			Name: to.Ptr("Policy1"),
		// 			Type: to.Ptr("Microsoft.Network/applicationgatewaywebapplicationfirewallpolicies"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/Policy1"),
		// 			Location: to.Ptr("WestUs"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armnetwork.WebApplicationFirewallPolicyPropertiesFormat{
		// 				CustomRules: []*armnetwork.WebApplicationFirewallCustomRule{
		// 					{
		// 						Name: to.Ptr("Rule1"),
		// 						Action: to.Ptr(armnetwork.WebApplicationFirewallActionBlock),
		// 						MatchConditions: []*armnetwork.MatchCondition{
		// 							{
		// 								MatchValues: []*string{
		// 									to.Ptr("192.168.1.0/24"),
		// 									to.Ptr("10.0.0.0/24")},
		// 									MatchVariables: []*armnetwork.MatchVariable{
		// 										{
		// 											VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRemoteAddr),
		// 									}},
		// 									NegationConditon: to.Ptr(false),
		// 									Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorIPMatch),
		// 									Transforms: []*armnetwork.WebApplicationFirewallTransform{
		// 									},
		// 							}},
		// 							Priority: to.Ptr[int32](1),
		// 							RuleType: to.Ptr(armnetwork.WebApplicationFirewallRuleTypeMatchRule),
		// 							State: to.Ptr(armnetwork.WebApplicationFirewallStateEnabled),
		// 						},
		// 						{
		// 							Name: to.Ptr("Rule2"),
		// 							Action: to.Ptr(armnetwork.WebApplicationFirewallActionBlock),
		// 							MatchConditions: []*armnetwork.MatchCondition{
		// 								{
		// 									MatchValues: []*string{
		// 										to.Ptr("192.168.1.0/24")},
		// 										MatchVariables: []*armnetwork.MatchVariable{
		// 											{
		// 												VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRemoteAddr),
		// 										}},
		// 										NegationConditon: to.Ptr(false),
		// 										Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorIPMatch),
		// 									},
		// 									{
		// 										MatchValues: []*string{
		// 											to.Ptr("Windows")},
		// 											MatchVariables: []*armnetwork.MatchVariable{
		// 												{
		// 													Selector: to.Ptr("UserAgent"),
		// 													VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariable("RequestHeader")),
		// 											}},
		// 											NegationConditon: to.Ptr(false),
		// 											Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorContains),
		// 									}},
		// 									Priority: to.Ptr[int32](2),
		// 									RuleType: to.Ptr(armnetwork.WebApplicationFirewallRuleTypeMatchRule),
		// 									State: to.Ptr(armnetwork.WebApplicationFirewallStateEnabled),
		// 								},
		// 								{
		// 									Name: to.Ptr("RateLimitRule3"),
		// 									Action: to.Ptr(armnetwork.WebApplicationFirewallActionBlock),
		// 									GroupByUserSession: []*armnetwork.GroupByUserSession{
		// 										{
		// 											GroupByVariables: []*armnetwork.GroupByVariable{
		// 												{
		// 													VariableName: to.Ptr(armnetwork.ApplicationGatewayFirewallUserSessionVariableClientAddr),
		// 											}},
		// 									}},
		// 									MatchConditions: []*armnetwork.MatchCondition{
		// 										{
		// 											MatchValues: []*string{
		// 												to.Ptr("192.168.1.0/24"),
		// 												to.Ptr("10.0.0.0/24")},
		// 												MatchVariables: []*armnetwork.MatchVariable{
		// 													{
		// 														VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRemoteAddr),
		// 												}},
		// 												NegationConditon: to.Ptr(true),
		// 												Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorIPMatch),
		// 										}},
		// 										Priority: to.Ptr[int32](3),
		// 										RateLimitDuration: to.Ptr(armnetwork.ApplicationGatewayFirewallRateLimitDurationOneMin),
		// 										RateLimitThreshold: to.Ptr[int32](10),
		// 										RuleType: to.Ptr(armnetwork.WebApplicationFirewallRuleTypeRateLimitRule),
		// 								}},
		// 								ManagedRules: &armnetwork.ManagedRulesDefinition{
		// 									ManagedRuleSets: []*armnetwork.ManagedRuleSet{
		// 										{
		// 											RuleSetType: to.Ptr("OWASP"),
		// 											RuleSetVersion: to.Ptr("3.2"),
		// 									}},
		// 								},
		// 								PolicySettings: &armnetwork.PolicySettings{
		// 									CustomBlockResponseBody: to.Ptr("SGVsbG8="),
		// 									CustomBlockResponseStatusCode: to.Ptr[int32](405),
		// 									FileUploadEnforcement: to.Ptr(true),
		// 									FileUploadLimitInMb: to.Ptr[int32](4000),
		// 									LogScrubbing: &armnetwork.PolicySettingsLogScrubbing{
		// 										ScrubbingRules: []*armnetwork.WebApplicationFirewallScrubbingRules{
		// 											{
		// 												MatchVariable: to.Ptr(armnetwork.ScrubbingRuleEntryMatchVariableRequestArgNames),
		// 												Selector: to.Ptr("test"),
		// 												SelectorMatchOperator: to.Ptr(armnetwork.ScrubbingRuleEntryMatchOperatorEquals),
		// 												State: to.Ptr(armnetwork.ScrubbingRuleEntryStateEnabled),
		// 											},
		// 											{
		// 												MatchVariable: to.Ptr(armnetwork.ScrubbingRuleEntryMatchVariableRequestIPAddress),
		// 												Selector: to.Ptr("*"),
		// 												SelectorMatchOperator: to.Ptr(armnetwork.ScrubbingRuleEntryMatchOperatorEqualsAny),
		// 												State: to.Ptr(armnetwork.ScrubbingRuleEntryStateEnabled),
		// 										}},
		// 										State: to.Ptr(armnetwork.WebApplicationFirewallScrubbingStateEnabled),
		// 									},
		// 									MaxRequestBodySizeInKb: to.Ptr[int32](2000),
		// 									Mode: to.Ptr(armnetwork.WebApplicationFirewallModePrevention),
		// 									RequestBodyCheck: to.Ptr(true),
		// 									RequestBodyEnforcement: to.Ptr(true),
		// 									RequestBodyInspectLimitInKB: to.Ptr[int32](2000),
		// 									State: to.Ptr(armnetwork.WebApplicationFirewallEnabledStateEnabled),
		// 								},
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								ResourceState: to.Ptr(armnetwork.WebApplicationFirewallPolicyResourceStateEnabled),
		// 							},
		// 					}},
		// 				}
	}
}
