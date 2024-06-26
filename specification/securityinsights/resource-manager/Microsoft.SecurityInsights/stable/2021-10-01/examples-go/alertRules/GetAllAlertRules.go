package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/alertRules/GetAllAlertRules.json
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
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
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
		// 							},
		// 						},
		// 						&armsecurityinsights.MicrosoftSecurityIncidentCreationAlertRule{
		// 							Name: to.Ptr("microsoftSecurityIncidentCreationRuleExample"),
		// 							Type: to.Ptr("Microsoft.SecurityInsights/alertRules"),
		// 							ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/microsoftSecurityIncidentCreationRuleExample"),
		// 							Etag: to.Ptr("\"260097e0-0000-0d00-0000-5d6fa88f0000\""),
		// 							Kind: to.Ptr(armsecurityinsights.AlertRuleKindMicrosoftSecurityIncidentCreation),
		// 							Properties: &armsecurityinsights.MicrosoftSecurityIncidentCreationAlertRuleProperties{
		// 								ProductFilter: to.Ptr(armsecurityinsights.MicrosoftSecurityProductNameMicrosoftCloudAppSecurity),
		// 								DisplayName: to.Ptr("testing displayname"),
		// 								Enabled: to.Ptr(true),
		// 								LastModifiedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-04T12:05:35.729Z"); return t}()),
		// 							},
		// 						},
		// 						&armsecurityinsights.FusionAlertRule{
		// 							Name: to.Ptr("myFirstFusionRule"),
		// 							Type: to.Ptr("Microsoft.SecurityInsights/alertRules"),
		// 							ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/myFirstFusionRule"),
		// 							Etag: to.Ptr("\"25005c11-0000-0d00-0000-5d6cc0e20000\""),
		// 							Kind: to.Ptr(armsecurityinsights.AlertRuleKindFusion),
		// 							Properties: &armsecurityinsights.FusionAlertRuleProperties{
		// 								Description: to.Ptr("In this mode, Sentinel combines low fidelity alerts, which themselves may not be actionable, and events across multiple products, into high fidelity security interesting incidents. The system looks at multiple products to produce actionable incidents. Custom tailored to each tenant, Fusion not only reduces false positive rates but also can detect attacks with limited or missing information. \nIncidents generated by Fusion system will encase two or more alerts. By design, Fusion incidents are low volume, high fidelity and will be high severity, which is why Fusion is turned ON by default in Azure Sentinel.\n\nFor Fusion to work, please configure the following data sources in Data Connectors tab:\nRequired - Azure Active Directory Identity Protection\nRequired - Microsoft Cloud App Security\nIf Available - Palo Alto Network\n\nFor full list of scenarios covered by Fusion, and detail instructions on how to configure the required data sources, go to aka.ms/SentinelFusion"),
		// 								AlertRuleTemplateName: to.Ptr("f71aba3d-28fb-450b-b192-4e76a83015c8"),
		// 								DisplayName: to.Ptr("Advanced Multi-Stage Attack Detection"),
		// 								Enabled: to.Ptr(false),
		// 								LastModifiedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-02T07:12:34.906Z"); return t}()),
		// 								Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
		// 								Tactics: []*armsecurityinsights.AttackTactic{
		// 									to.Ptr(armsecurityinsights.AttackTacticPersistence),
		// 									to.Ptr(armsecurityinsights.AttackTacticLateralMovement),
		// 									to.Ptr(armsecurityinsights.AttackTacticExfiltration),
		// 									to.Ptr(armsecurityinsights.AttackTacticCommandAndControl)},
		// 								},
		// 						}},
		// 					}
	}
}
