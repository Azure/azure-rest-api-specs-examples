package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/CreateFusionAlertRuleWithFusionScenarioExclusion.json
func ExampleAlertRulesClient_CreateOrUpdate_createsOrUpdatesAFusionAlertRuleWithScenarioExclusionPattern() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertRulesClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "myFirstFusionRule", &armsecurityinsights.FusionAlertRule{
		Etag: to.Ptr("3d00c3ca-0000-0100-0000-5d42d5010000"),
		Kind: to.Ptr(armsecurityinsights.AlertRuleKindFusion),
		Properties: &armsecurityinsights.FusionAlertRuleProperties{
			AlertRuleTemplateName: to.Ptr("f71aba3d-28fb-450b-b192-4e76a83015c8"),
			Enabled:               to.Ptr(true),
			SourceSettings: []*armsecurityinsights.FusionSourceSettings{
				{
					Enabled:    to.Ptr(true),
					SourceName: to.Ptr("Anomalies"),
				},
				{
					Enabled:    to.Ptr(true),
					SourceName: to.Ptr("Alert providers"),
					SourceSubTypes: []*armsecurityinsights.FusionSourceSubTypeSetting{
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Azure Active Directory Identity Protection"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Azure Defender"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Azure Defender for IoT"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Microsoft 365 Defender"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Microsoft Cloud App Security"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Microsoft Defender for Endpoint"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Microsoft Defender for Identity"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Microsoft Defender for Office 365"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Azure Sentinel scheduled analytics rules"),
						}},
				},
				{
					Enabled:    to.Ptr(true),
					SourceName: to.Ptr("Raw logs from other sources"),
					SourceSubTypes: []*armsecurityinsights.FusionSourceSubTypeSetting{
						{
							Enabled:           to.Ptr(true),
							SeverityFilters:   &armsecurityinsights.FusionSubTypeSeverityFilter{},
							SourceSubTypeName: to.Ptr("Palo Alto Networks"),
						}},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.AlertRulesClientCreateOrUpdateResponse{
	// 	                            AlertRuleClassification: &armsecurityinsights.FusionAlertRule{
	// 		Name: to.Ptr("myFirstFusionRule"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/alertRules"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/myFirstFusionRule"),
	// 		Etag: to.Ptr("\"260090e2-0000-0d00-0000-5d6fb8670000\""),
	// 		Kind: to.Ptr(armsecurityinsights.AlertRuleKindFusion),
	// 		Properties: &armsecurityinsights.FusionAlertRuleProperties{
	// 			Description: to.Ptr("Using Fusion technology based on machine learning, Azure Sentinel automatically detects multistage attacks by identifying combinations of anomalous behaviors and suspicious activities observed at various stages of the kill chain. On the basis of these discoveries, Azure Sentinel generates incidents that would otherwise be very difficult to catch. By design, these incidents are low-volume, high-fidelity, and high-severity, which is why this detection is turned ON by default.\n\nThere are a total of 122 Fusion incident types detected by Azure Sentinel.\n\nTo detect these multistage attacks, the following data connectors must be configured:\n- Azure Active Directory Identity Protection.\n- Microsoft Cloud App Security.\n- Microsoft Defender for Endpoint.\n- Azure Defender.\n- Palo Alto Networks.\n- Scheduled Analytics Rules supported by Fusion\n\nFor a full list and description of each scenario that is supported for these multistage attacks, go to https://aka.ms/SentinelFusion."),
	// 			AlertRuleTemplateName: to.Ptr("f71aba3d-28fb-450b-b192-4e76a83015c8"),
	// 			DisplayName: to.Ptr("Advanced Multi-Stage Attack Detection"),
	// 			Enabled: to.Ptr(true),
	// 			LastModifiedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-04T13:13:11.534Z"); return t}()),
	// 			Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
	// 			SourceSettings: []*armsecurityinsights.FusionSourceSettings{
	// 				{
	// 					Enabled: to.Ptr(true),
	// 					SourceName: to.Ptr("Anomalies"),
	// 				},
	// 				{
	// 					Enabled: to.Ptr(true),
	// 					SourceName: to.Ptr("Alert providers"),
	// 					SourceSubTypes: []*armsecurityinsights.FusionSourceSubTypeSetting{
	// 						{
	// 							Enabled: to.Ptr(true),
	// 							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
	// 								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
	// 								}},
	// 								IsSupported: to.Ptr(true),
	// 							},
	// 							SourceSubTypeName: to.Ptr("Azure Active Directory Identity Protection"),
	// 						},
	// 						{
	// 							Enabled: to.Ptr(true),
	// 							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
	// 								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
	// 								}},
	// 								IsSupported: to.Ptr(true),
	// 							},
	// 							SourceSubTypeName: to.Ptr("Azure Defender"),
	// 						},
	// 						{
	// 							Enabled: to.Ptr(true),
	// 							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
	// 								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
	// 								}},
	// 								IsSupported: to.Ptr(true),
	// 							},
	// 							SourceSubTypeName: to.Ptr("Azure Defender for IoT"),
	// 						},
	// 						{
	// 							Enabled: to.Ptr(true),
	// 							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
	// 								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
	// 								}},
	// 								IsSupported: to.Ptr(true),
	// 							},
	// 							SourceSubTypeName: to.Ptr("Microsoft 365 Defender"),
	// 						},
	// 						{
	// 							Enabled: to.Ptr(true),
	// 							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
	// 								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
	// 								}},
	// 								IsSupported: to.Ptr(true),
	// 							},
	// 							SourceSubTypeName: to.Ptr("Microsoft Cloud App Security"),
	// 						},
	// 						{
	// 							Enabled: to.Ptr(true),
	// 							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
	// 								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
	// 								}},
	// 								IsSupported: to.Ptr(true),
	// 							},
	// 							SourceSubTypeName: to.Ptr("Microsoft Defender for Endpoint"),
	// 						},
	// 						{
	// 							Enabled: to.Ptr(true),
	// 							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
	// 								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
	// 								}},
	// 								IsSupported: to.Ptr(true),
	// 							},
	// 							SourceSubTypeName: to.Ptr("Microsoft Defender for Identity"),
	// 						},
	// 						{
	// 							Enabled: to.Ptr(true),
	// 							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
	// 								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
	// 								}},
	// 								IsSupported: to.Ptr(true),
	// 							},
	// 							SourceSubTypeName: to.Ptr("Microsoft Defender for Office 365"),
	// 						},
	// 						{
	// 							Enabled: to.Ptr(true),
	// 							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
	// 								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
	// 									},
	// 									{
	// 										Enabled: to.Ptr(true),
	// 										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
	// 								}},
	// 								IsSupported: to.Ptr(true),
	// 							},
	// 							SourceSubTypeName: to.Ptr("Azure Sentinel scheduled analytics rules"),
	// 					}},
	// 				},
	// 				{
	// 					Enabled: to.Ptr(true),
	// 					SourceName: to.Ptr("Raw logs from other sources"),
	// 					SourceSubTypes: []*armsecurityinsights.FusionSourceSubTypeSetting{
	// 						{
	// 							Enabled: to.Ptr(true),
	// 							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
	// 								IsSupported: to.Ptr(false),
	// 							},
	// 							SourceSubTypeName: to.Ptr("Palo Alto Networks"),
	// 					}},
	// 			}},
	// 			Tactics: []*armsecurityinsights.AttackTactic{
	// 				to.Ptr(armsecurityinsights.AttackTacticCollection),
	// 				to.Ptr(armsecurityinsights.AttackTacticCommandAndControl),
	// 				to.Ptr(armsecurityinsights.AttackTacticCredentialAccess),
	// 				to.Ptr(armsecurityinsights.AttackTacticDefenseEvasion),
	// 				to.Ptr(armsecurityinsights.AttackTacticDiscovery),
	// 				to.Ptr(armsecurityinsights.AttackTacticExecution),
	// 				to.Ptr(armsecurityinsights.AttackTacticExfiltration),
	// 				to.Ptr(armsecurityinsights.AttackTacticImpact),
	// 				to.Ptr(armsecurityinsights.AttackTacticInitialAccess),
	// 				to.Ptr(armsecurityinsights.AttackTacticLateralMovement),
	// 				to.Ptr(armsecurityinsights.AttackTacticPersistence),
	// 				to.Ptr(armsecurityinsights.AttackTacticPrivilegeEscalation)},
	// 			},
	// 		},
	// 		                        }
}
