package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/GetAllAlertRules.json
func ExampleAlertRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertRulesClient().NewListPager("myRg", "myWorkspace", nil)
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
		// page.AlertRulesList = armsecurityinsights.AlertRulesList{
		// 	Value: []armsecurityinsights.AlertRuleClassification{
		// 		&armsecurityinsights.ScheduledAlertRule{
		// 			Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/alertRules"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.AlertRuleKindScheduled),
		// 			Properties: &armsecurityinsights.ScheduledAlertRuleProperties{
		// 				AlertDetailsOverride: &armsecurityinsights.AlertDetailsOverride{
		// 					AlertDescriptionFormat: to.Ptr("Suspicious activity was made by {{ComputerIP}}"),
		// 					AlertDisplayNameFormat: to.Ptr("Alert from {{Computer}}"),
		// 				},
		// 				CustomDetails: map[string]*string{
		// 					"OperatingSystemName": to.Ptr("OSName"),
		// 					"OperatingSystemType": to.Ptr("OSType"),
		// 				},
		// 				EntityMappings: []*armsecurityinsights.EntityMapping{
		// 					{
		// 						EntityType: to.Ptr(armsecurityinsights.EntityMappingTypeHost),
		// 						FieldMappings: []*armsecurityinsights.FieldMapping{
		// 							{
		// 								ColumnName: to.Ptr("Computer"),
		// 								Identifier: to.Ptr("FullName"),
		// 						}},
		// 					},
		// 					{
		// 						EntityType: to.Ptr(armsecurityinsights.EntityMappingTypeIP),
		// 						FieldMappings: []*armsecurityinsights.FieldMapping{
		// 							{
		// 								ColumnName: to.Ptr("ComputerIP"),
		// 								Identifier: to.Ptr("Address"),
		// 						}},
		// 				}},
		// 				EventGroupingSettings: &armsecurityinsights.EventGroupingSettings{
		// 					AggregationKind: to.Ptr(armsecurityinsights.EventGroupingAggregationKindAlertPerResult),
		// 				},
		// 				Query: to.Ptr("Heartbeat"),
		// 				QueryFrequency: to.Ptr("PT1H"),
		// 				QueryPeriod: to.Ptr("P2DT1H30M"),
		// 				Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
		// 				TriggerOperator: to.Ptr(armsecurityinsights.TriggerOperatorGreaterThan),
		// 				TriggerThreshold: to.Ptr[int32](0),
		// 				Description: to.Ptr("An example for a scheduled rule"),
		// 				DisplayName: to.Ptr("My scheduled rule"),
		// 				Enabled: to.Ptr(true),
		// 				IncidentConfiguration: &armsecurityinsights.IncidentConfiguration{
		// 					CreateIncident: to.Ptr(true),
		// 					GroupingConfiguration: &armsecurityinsights.GroupingConfiguration{
		// 						Enabled: to.Ptr(true),
		// 						GroupByAlertDetails: []*armsecurityinsights.AlertDetail{
		// 							to.Ptr(armsecurityinsights.AlertDetailDisplayName)},
		// 							GroupByCustomDetails: []*string{
		// 								to.Ptr("OperatingSystemType"),
		// 								to.Ptr("OperatingSystemName")},
		// 								GroupByEntities: []*armsecurityinsights.EntityMappingType{
		// 									to.Ptr(armsecurityinsights.EntityMappingTypeHost)},
		// 									LookbackDuration: to.Ptr("PT5H"),
		// 									MatchingMethod: to.Ptr(armsecurityinsights.MatchingMethodSelected),
		// 									ReopenClosedIncident: to.Ptr(false),
		// 								},
		// 							},
		// 							LastModifiedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-01T13:17:30.000Z"); return t}()),
		// 							SuppressionDuration: to.Ptr("PT1H"),
		// 							SuppressionEnabled: to.Ptr(false),
		// 							Tactics: []*armsecurityinsights.AttackTactic{
		// 								to.Ptr(armsecurityinsights.AttackTacticPersistence),
		// 								to.Ptr(armsecurityinsights.AttackTacticLateralMovement)},
		// 								Techniques: []*string{
		// 									to.Ptr("T1037"),
		// 									to.Ptr("T1021")},
		// 								},
		// 							},
		// 							&armsecurityinsights.MicrosoftSecurityIncidentCreationAlertRule{
		// 								Name: to.Ptr("microsoftSecurityIncidentCreationRuleExample"),
		// 								Type: to.Ptr("Microsoft.SecurityInsights/alertRules"),
		// 								ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/microsoftSecurityIncidentCreationRuleExample"),
		// 								Etag: to.Ptr("\"260097e0-0000-0d00-0000-5d6fa88f0000\""),
		// 								Kind: to.Ptr(armsecurityinsights.AlertRuleKindMicrosoftSecurityIncidentCreation),
		// 								Properties: &armsecurityinsights.MicrosoftSecurityIncidentCreationAlertRuleProperties{
		// 									ProductFilter: to.Ptr(armsecurityinsights.MicrosoftSecurityProductNameMicrosoftCloudAppSecurity),
		// 									DisplayName: to.Ptr("testing displayname"),
		// 									Enabled: to.Ptr(true),
		// 									LastModifiedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-04T12:05:35.729Z"); return t}()),
		// 								},
		// 							},
		// 							&armsecurityinsights.FusionAlertRule{
		// 								Name: to.Ptr("myFirstFusionRule"),
		// 								Type: to.Ptr("Microsoft.SecurityInsights/alertRules"),
		// 								ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/myFirstFusionRule"),
		// 								Etag: to.Ptr("\"25005c11-0000-0d00-0000-5d6cc0e20000\""),
		// 								Kind: to.Ptr(armsecurityinsights.AlertRuleKindFusion),
		// 								Properties: &armsecurityinsights.FusionAlertRuleProperties{
		// 									Description: to.Ptr("Microsoft Sentinel uses Fusion, a correlation engine based on scalable machine learning algorithms, to automatically detect multistage attacks by identifying combinations of anomalous behaviors and suspicious activities that are observed at various stages of the kill chain. On the basis of these discoveries, Azure Sentinel generates incidents that would otherwise be very difficult to catch. By design, these incidents are low-volume, high-fidelity, and high-severity, which is why this detection is turned ON by default.\n\nSince Fusion correlates multiple signals from various products to detect advanced multistage attacks, successful Fusion detections are presented as Fusion incidents on the Microsoft Sentinel Incidents page. This rule covers the following detections:\n- Fusion for emerging threats\n- Fusion for ransomware\n- Scenario-based Fusion detections (122 scenarios)\n\nTo enable these detections, we recommend you configure the following data connectors for best results:\n- Out-of-the-box anomaly detections\n- Azure Active Directory Identity Protection\n- Azure Defender\n- Azure Defender for IoT\n- Microsoft 365 Defender\n- Microsoft Cloud App Security    \n- Microsoft Defender for Endpoint\n- Microsoft Defender for Identity\n- Microsoft Defender for Office 365\n- Palo Alto Networks\n- Scheduled analytics rules, both built-in and those created by your security analysts. Analytics rules must contain kill-chain (tactics) and entity mapping information in order to be used by Fusion.\n\nFor the full description of each detection that is supported by Fusion, go to https://aka.ms/SentinelFusion."),
		// 									AlertRuleTemplateName: to.Ptr("f71aba3d-28fb-450b-b192-4e76a83015c8"),
		// 									DisplayName: to.Ptr("Advanced Multi-Stage Attack Detection"),
		// 									Enabled: to.Ptr(true),
		// 									LastModifiedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-22T07:12:34.906Z"); return t}()),
		// 									Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
		// 									SourceSettings: []*armsecurityinsights.FusionSourceSettings{
		// 										{
		// 											Enabled: to.Ptr(true),
		// 											SourceName: to.Ptr("Anomalies"),
		// 										},
		// 										{
		// 											Enabled: to.Ptr(true),
		// 											SourceName: to.Ptr("Alert providers"),
		// 											SourceSubTypes: []*armsecurityinsights.FusionSourceSubTypeSetting{
		// 												{
		// 													Enabled: to.Ptr(true),
		// 													SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
		// 														Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 														}},
		// 														IsSupported: to.Ptr(true),
		// 													},
		// 													SourceSubTypeDisplayName: to.Ptr("Azure Active Directory Identity Protection"),
		// 													SourceSubTypeName: to.Ptr("Azure Active Directory Identity Protection"),
		// 												},
		// 												{
		// 													Enabled: to.Ptr(true),
		// 													SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
		// 														Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 														}},
		// 														IsSupported: to.Ptr(true),
		// 													},
		// 													SourceSubTypeDisplayName: to.Ptr("Microsoft Defender for Cloud"),
		// 													SourceSubTypeName: to.Ptr("Azure Defender"),
		// 												},
		// 												{
		// 													Enabled: to.Ptr(true),
		// 													SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
		// 														Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 														}},
		// 														IsSupported: to.Ptr(true),
		// 													},
		// 													SourceSubTypeDisplayName: to.Ptr("Microsoft Defender for IoT"),
		// 													SourceSubTypeName: to.Ptr("Azure Defender for IoT"),
		// 												},
		// 												{
		// 													Enabled: to.Ptr(true),
		// 													SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
		// 														Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 														}},
		// 														IsSupported: to.Ptr(true),
		// 													},
		// 													SourceSubTypeDisplayName: to.Ptr("Microsoft 365 Defender"),
		// 													SourceSubTypeName: to.Ptr("Microsoft 365 Defender"),
		// 												},
		// 												{
		// 													Enabled: to.Ptr(true),
		// 													SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
		// 														Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 														}},
		// 														IsSupported: to.Ptr(true),
		// 													},
		// 													SourceSubTypeDisplayName: to.Ptr("Microsoft Cloud App Security"),
		// 													SourceSubTypeName: to.Ptr("Microsoft Cloud App Security"),
		// 												},
		// 												{
		// 													Enabled: to.Ptr(true),
		// 													SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
		// 														Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 														}},
		// 														IsSupported: to.Ptr(true),
		// 													},
		// 													SourceSubTypeDisplayName: to.Ptr("Microsoft Defender for Endpoint"),
		// 													SourceSubTypeName: to.Ptr("Microsoft Defender for Endpoint"),
		// 												},
		// 												{
		// 													Enabled: to.Ptr(true),
		// 													SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
		// 														Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 														}},
		// 														IsSupported: to.Ptr(true),
		// 													},
		// 													SourceSubTypeDisplayName: to.Ptr("Microsoft Defender for Identity"),
		// 													SourceSubTypeName: to.Ptr("Microsoft Defender for Identity"),
		// 												},
		// 												{
		// 													Enabled: to.Ptr(true),
		// 													SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
		// 														Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 														}},
		// 														IsSupported: to.Ptr(true),
		// 													},
		// 													SourceSubTypeDisplayName: to.Ptr("Microsoft Defender for Office 365"),
		// 													SourceSubTypeName: to.Ptr("Microsoft Defender for Office 365"),
		// 												},
		// 												{
		// 													Enabled: to.Ptr(true),
		// 													SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
		// 														Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 															},
		// 															{
		// 																Enabled: to.Ptr(true),
		// 																Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 														}},
		// 														IsSupported: to.Ptr(true),
		// 													},
		// 													SourceSubTypeDisplayName: to.Ptr("Azure Sentinel scheduled analytics rules"),
		// 													SourceSubTypeName: to.Ptr("Azure Sentinel scheduled analytics rules"),
		// 											}},
		// 										},
		// 										{
		// 											Enabled: to.Ptr(true),
		// 											SourceName: to.Ptr("Raw logs from other sources"),
		// 											SourceSubTypes: []*armsecurityinsights.FusionSourceSubTypeSetting{
		// 												{
		// 													Enabled: to.Ptr(true),
		// 													SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
		// 														IsSupported: to.Ptr(false),
		// 													},
		// 													SourceSubTypeDisplayName: to.Ptr("Palo Alto Networks"),
		// 													SourceSubTypeName: to.Ptr("Palo Alto Networks"),
		// 											}},
		// 									}},
		// 									Tactics: []*armsecurityinsights.AttackTactic{
		// 										to.Ptr(armsecurityinsights.AttackTacticCollection),
		// 										to.Ptr(armsecurityinsights.AttackTacticCommandAndControl),
		// 										to.Ptr(armsecurityinsights.AttackTacticCredentialAccess),
		// 										to.Ptr(armsecurityinsights.AttackTacticDefenseEvasion),
		// 										to.Ptr(armsecurityinsights.AttackTacticDiscovery),
		// 										to.Ptr(armsecurityinsights.AttackTacticExecution),
		// 										to.Ptr(armsecurityinsights.AttackTacticExfiltration),
		// 										to.Ptr(armsecurityinsights.AttackTacticImpact),
		// 										to.Ptr(armsecurityinsights.AttackTacticInitialAccess),
		// 										to.Ptr(armsecurityinsights.AttackTacticLateralMovement),
		// 										to.Ptr(armsecurityinsights.AttackTacticPersistence),
		// 										to.Ptr(armsecurityinsights.AttackTacticPrivilegeEscalation)},
		// 									},
		// 							}},
		// 						}
	}
}
