package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/WafPolicyCreateOrUpdate.json
func ExampleWebApplicationFirewallPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebApplicationFirewallPoliciesClient().CreateOrUpdate(ctx, "rg1", "Policy1", armnetwork.WebApplicationFirewallPolicy{
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
				},
				{
					Name:   to.Ptr("RateLimitRule3"),
					Action: to.Ptr(armnetwork.WebApplicationFirewallActionBlock),
					GroupByUserSession: []*armnetwork.GroupByUserSession{
						{
							GroupByVariables: []*armnetwork.GroupByVariable{
								{
									VariableName: to.Ptr(armnetwork.ApplicationGatewayFirewallUserSessionVariableClientAddr),
								}},
						}},
					MatchConditions: []*armnetwork.MatchCondition{
						{
							MatchValues: []*string{
								to.Ptr("192.168.1.0/24"),
								to.Ptr("10.0.0.0/24")},
							MatchVariables: []*armnetwork.MatchVariable{
								{
									VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRemoteAddr),
								}},
							NegationConditon: to.Ptr(true),
							Operator:         to.Ptr(armnetwork.WebApplicationFirewallOperatorIPMatch),
						}},
					Priority:           to.Ptr[int32](3),
					RateLimitDuration:  to.Ptr(armnetwork.ApplicationGatewayFirewallRateLimitDurationOneMin),
					RateLimitThreshold: to.Ptr[int32](10),
					RuleType:           to.Ptr(armnetwork.WebApplicationFirewallRuleTypeRateLimitRule),
				},
				{
					Name:   to.Ptr("Rule4"),
					Action: to.Ptr(armnetwork.WebApplicationFirewallActionJSChallenge),
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
								to.Ptr("Bot")},
							MatchVariables: []*armnetwork.MatchVariable{
								{
									Selector:     to.Ptr("UserAgent"),
									VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRequestHeaders),
								}},
							Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorContains),
						}},
					Priority: to.Ptr[int32](4),
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
						RuleGroupOverrides: []*armnetwork.ManagedRuleGroupOverride{
							{
								RuleGroupName: to.Ptr("REQUEST-931-APPLICATION-ATTACK-RFI"),
								Rules: []*armnetwork.ManagedRuleOverride{
									{
										Action: to.Ptr(armnetwork.ActionTypeLog),
										RuleID: to.Ptr("931120"),
										State:  to.Ptr(armnetwork.ManagedRuleEnabledStateEnabled),
									},
									{
										Action: to.Ptr(armnetwork.ActionTypeAnomalyScoring),
										RuleID: to.Ptr("931130"),
										State:  to.Ptr(armnetwork.ManagedRuleEnabledStateDisabled),
									}},
							}},
						RuleSetType:    to.Ptr("OWASP"),
						RuleSetVersion: to.Ptr("3.2"),
					},
					{
						RuleGroupOverrides: []*armnetwork.ManagedRuleGroupOverride{
							{
								RuleGroupName: to.Ptr("UnknownBots"),
								Rules: []*armnetwork.ManagedRuleOverride{
									{
										Action: to.Ptr(armnetwork.ActionTypeJSChallenge),
										RuleID: to.Ptr("300700"),
										State:  to.Ptr(armnetwork.ManagedRuleEnabledStateEnabled),
									}},
							}},
						RuleSetType:    to.Ptr("Microsoft_BotManagerRuleSet"),
						RuleSetVersion: to.Ptr("1.0"),
					}},
			},
			PolicySettings: &armnetwork.PolicySettings{
				JsChallengeCookieExpirationInMins: to.Ptr[int32](100),
				LogScrubbing: &armnetwork.PolicySettingsLogScrubbing{
					ScrubbingRules: []*armnetwork.WebApplicationFirewallScrubbingRules{
						{
							MatchVariable:         to.Ptr(armnetwork.ScrubbingRuleEntryMatchVariableRequestArgNames),
							Selector:              to.Ptr("test"),
							SelectorMatchOperator: to.Ptr(armnetwork.ScrubbingRuleEntryMatchOperatorEquals),
							State:                 to.Ptr(armnetwork.ScrubbingRuleEntryStateEnabled),
						},
						{
							MatchVariable:         to.Ptr(armnetwork.ScrubbingRuleEntryMatchVariableRequestIPAddress),
							SelectorMatchOperator: to.Ptr(armnetwork.ScrubbingRuleEntryMatchOperatorEqualsAny),
							State:                 to.Ptr(armnetwork.ScrubbingRuleEntryStateEnabled),
						}},
					State: to.Ptr(armnetwork.WebApplicationFirewallScrubbingStateEnabled),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WebApplicationFirewallPolicy = armnetwork.WebApplicationFirewallPolicy{
	// 	Name: to.Ptr("Policy1"),
	// 	Type: to.Ptr("Microsoft.Network/applicationgatewaywebapplicationfirewallpolicies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/Policy1"),
	// 	Location: to.Ptr("WestUs"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armnetwork.WebApplicationFirewallPolicyPropertiesFormat{
	// 		CustomRules: []*armnetwork.WebApplicationFirewallCustomRule{
	// 			{
	// 				Name: to.Ptr("Rule1"),
	// 				Action: to.Ptr(armnetwork.WebApplicationFirewallActionBlock),
	// 				MatchConditions: []*armnetwork.MatchCondition{
	// 					{
	// 						MatchValues: []*string{
	// 							to.Ptr("192.168.1.0/24"),
	// 							to.Ptr("10.0.0.0/24")},
	// 							MatchVariables: []*armnetwork.MatchVariable{
	// 								{
	// 									VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRemoteAddr),
	// 							}},
	// 							NegationConditon: to.Ptr(false),
	// 							Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorIPMatch),
	// 							Transforms: []*armnetwork.WebApplicationFirewallTransform{
	// 							},
	// 					}},
	// 					Priority: to.Ptr[int32](1),
	// 					RuleType: to.Ptr(armnetwork.WebApplicationFirewallRuleTypeMatchRule),
	// 					State: to.Ptr(armnetwork.WebApplicationFirewallStateEnabled),
	// 				},
	// 				{
	// 					Name: to.Ptr("Rule2"),
	// 					Action: to.Ptr(armnetwork.WebApplicationFirewallActionBlock),
	// 					MatchConditions: []*armnetwork.MatchCondition{
	// 						{
	// 							MatchValues: []*string{
	// 								to.Ptr("192.168.1.0/24")},
	// 								MatchVariables: []*armnetwork.MatchVariable{
	// 									{
	// 										VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRemoteAddr),
	// 								}},
	// 								NegationConditon: to.Ptr(false),
	// 								Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorIPMatch),
	// 							},
	// 							{
	// 								MatchValues: []*string{
	// 									to.Ptr("Windows")},
	// 									MatchVariables: []*armnetwork.MatchVariable{
	// 										{
	// 											Selector: to.Ptr("UserAgent"),
	// 											VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariable("RequestHeader")),
	// 									}},
	// 									NegationConditon: to.Ptr(false),
	// 									Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorContains),
	// 							}},
	// 							Priority: to.Ptr[int32](2),
	// 							RuleType: to.Ptr(armnetwork.WebApplicationFirewallRuleTypeMatchRule),
	// 							State: to.Ptr(armnetwork.WebApplicationFirewallStateEnabled),
	// 						},
	// 						{
	// 							Name: to.Ptr("RateLimitRule3"),
	// 							Action: to.Ptr(armnetwork.WebApplicationFirewallActionBlock),
	// 							GroupByUserSession: []*armnetwork.GroupByUserSession{
	// 								{
	// 									GroupByVariables: []*armnetwork.GroupByVariable{
	// 										{
	// 											VariableName: to.Ptr(armnetwork.ApplicationGatewayFirewallUserSessionVariableClientAddr),
	// 									}},
	// 							}},
	// 							MatchConditions: []*armnetwork.MatchCondition{
	// 								{
	// 									MatchValues: []*string{
	// 										to.Ptr("192.168.1.0/24"),
	// 										to.Ptr("10.0.0.0/24")},
	// 										MatchVariables: []*armnetwork.MatchVariable{
	// 											{
	// 												VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRemoteAddr),
	// 										}},
	// 										NegationConditon: to.Ptr(true),
	// 										Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorIPMatch),
	// 								}},
	// 								Priority: to.Ptr[int32](3),
	// 								RateLimitDuration: to.Ptr(armnetwork.ApplicationGatewayFirewallRateLimitDurationOneMin),
	// 								RateLimitThreshold: to.Ptr[int32](10),
	// 								RuleType: to.Ptr(armnetwork.WebApplicationFirewallRuleTypeRateLimitRule),
	// 							},
	// 							{
	// 								Name: to.Ptr("Rule4"),
	// 								Action: to.Ptr(armnetwork.WebApplicationFirewallActionJSChallenge),
	// 								MatchConditions: []*armnetwork.MatchCondition{
	// 									{
	// 										MatchValues: []*string{
	// 											to.Ptr("192.168.1.0/24")},
	// 											MatchVariables: []*armnetwork.MatchVariable{
	// 												{
	// 													VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRemoteAddr),
	// 											}},
	// 											NegationConditon: to.Ptr(false),
	// 											Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorIPMatch),
	// 										},
	// 										{
	// 											MatchValues: []*string{
	// 												to.Ptr("Bot")},
	// 												MatchVariables: []*armnetwork.MatchVariable{
	// 													{
	// 														Selector: to.Ptr("UserAgent"),
	// 														VariableName: to.Ptr(armnetwork.WebApplicationFirewallMatchVariableRequestHeaders),
	// 												}},
	// 												NegationConditon: to.Ptr(false),
	// 												Operator: to.Ptr(armnetwork.WebApplicationFirewallOperatorContains),
	// 										}},
	// 										Priority: to.Ptr[int32](4),
	// 										RuleType: to.Ptr(armnetwork.WebApplicationFirewallRuleTypeMatchRule),
	// 										State: to.Ptr(armnetwork.WebApplicationFirewallStateEnabled),
	// 								}},
	// 								ManagedRules: &armnetwork.ManagedRulesDefinition{
	// 									Exclusions: []*armnetwork.OwaspCrsExclusionEntry{
	// 										{
	// 											ExclusionManagedRuleSets: []*armnetwork.ExclusionManagedRuleSet{
	// 												{
	// 													RuleGroups: []*armnetwork.ExclusionManagedRuleGroup{
	// 														{
	// 															RuleGroupName: to.Ptr("REQUEST-930-APPLICATION-ATTACK-LFI"),
	// 															Rules: []*armnetwork.ExclusionManagedRule{
	// 																{
	// 																	RuleID: to.Ptr("930120"),
	// 															}},
	// 														},
	// 														{
	// 															RuleGroupName: to.Ptr("REQUEST-932-APPLICATION-ATTACK-RCE"),
	// 													}},
	// 													RuleSetType: to.Ptr("OWASP"),
	// 													RuleSetVersion: to.Ptr("3.2"),
	// 											}},
	// 											MatchVariable: to.Ptr(armnetwork.OwaspCrsExclusionEntryMatchVariableRequestArgNames),
	// 											Selector: to.Ptr("hello"),
	// 											SelectorMatchOperator: to.Ptr(armnetwork.OwaspCrsExclusionEntrySelectorMatchOperatorStartsWith),
	// 										},
	// 										{
	// 											ExclusionManagedRuleSets: []*armnetwork.ExclusionManagedRuleSet{
	// 												{
	// 													RuleGroups: []*armnetwork.ExclusionManagedRuleGroup{
	// 													},
	// 													RuleSetType: to.Ptr("OWASP"),
	// 													RuleSetVersion: to.Ptr("3.1"),
	// 											}},
	// 											MatchVariable: to.Ptr(armnetwork.OwaspCrsExclusionEntryMatchVariableRequestArgNames),
	// 											Selector: to.Ptr("hello"),
	// 											SelectorMatchOperator: to.Ptr(armnetwork.OwaspCrsExclusionEntrySelectorMatchOperatorEndsWith),
	// 										},
	// 										{
	// 											MatchVariable: to.Ptr(armnetwork.OwaspCrsExclusionEntryMatchVariableRequestArgNames),
	// 											Selector: to.Ptr("test"),
	// 											SelectorMatchOperator: to.Ptr(armnetwork.OwaspCrsExclusionEntrySelectorMatchOperatorStartsWith),
	// 										},
	// 										{
	// 											MatchVariable: to.Ptr(armnetwork.OwaspCrsExclusionEntryMatchVariableRequestArgValues),
	// 											Selector: to.Ptr("test"),
	// 											SelectorMatchOperator: to.Ptr(armnetwork.OwaspCrsExclusionEntrySelectorMatchOperatorStartsWith),
	// 									}},
	// 									ManagedRuleSets: []*armnetwork.ManagedRuleSet{
	// 										{
	// 											RuleGroupOverrides: []*armnetwork.ManagedRuleGroupOverride{
	// 												{
	// 													RuleGroupName: to.Ptr("REQUEST-931-APPLICATION-ATTACK-RFI"),
	// 													Rules: []*armnetwork.ManagedRuleOverride{
	// 														{
	// 															Action: to.Ptr(armnetwork.ActionTypeLog),
	// 															RuleID: to.Ptr("931120"),
	// 															State: to.Ptr(armnetwork.ManagedRuleEnabledStateEnabled),
	// 														},
	// 														{
	// 															Action: to.Ptr(armnetwork.ActionTypeAnomalyScoring),
	// 															RuleID: to.Ptr("931130"),
	// 															State: to.Ptr(armnetwork.ManagedRuleEnabledStateDisabled),
	// 													}},
	// 											}},
	// 											RuleSetType: to.Ptr("OWASP"),
	// 											RuleSetVersion: to.Ptr("3.2"),
	// 										},
	// 										{
	// 											RuleGroupOverrides: []*armnetwork.ManagedRuleGroupOverride{
	// 												{
	// 													RuleGroupName: to.Ptr("UnknownBots"),
	// 													Rules: []*armnetwork.ManagedRuleOverride{
	// 														{
	// 															Action: to.Ptr(armnetwork.ActionTypeJSChallenge),
	// 															RuleID: to.Ptr("300700"),
	// 															State: to.Ptr(armnetwork.ManagedRuleEnabledStateEnabled),
	// 													}},
	// 											}},
	// 											RuleSetType: to.Ptr("Microsoft_BotManagerRuleSet"),
	// 											RuleSetVersion: to.Ptr("1.0"),
	// 									}},
	// 								},
	// 								PolicySettings: &armnetwork.PolicySettings{
	// 									CustomBlockResponseBody: to.Ptr("SGVsbG8="),
	// 									CustomBlockResponseStatusCode: to.Ptr[int32](405),
	// 									FileUploadEnforcement: to.Ptr(true),
	// 									FileUploadLimitInMb: to.Ptr[int32](4000),
	// 									JsChallengeCookieExpirationInMins: to.Ptr[int32](100),
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
	// 												SelectorMatchOperator: to.Ptr(armnetwork.ScrubbingRuleEntryMatchOperatorEqualsAny),
	// 												State: to.Ptr(armnetwork.ScrubbingRuleEntryStateEnabled),
	// 										}},
	// 										State: to.Ptr(armnetwork.WebApplicationFirewallScrubbingStateEnabled),
	// 									},
	// 									MaxRequestBodySizeInKb: to.Ptr[int32](2000),
	// 									Mode: to.Ptr(armnetwork.WebApplicationFirewallModeDetection),
	// 									RequestBodyCheck: to.Ptr(true),
	// 									RequestBodyEnforcement: to.Ptr(true),
	// 									RequestBodyInspectLimitInKB: to.Ptr[int32](2000),
	// 									State: to.Ptr(armnetwork.WebApplicationFirewallEnabledStateEnabled),
	// 								},
	// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 								ResourceState: to.Ptr(armnetwork.WebApplicationFirewallPolicyResourceStateEnabled),
	// 							},
	// 						}
}
