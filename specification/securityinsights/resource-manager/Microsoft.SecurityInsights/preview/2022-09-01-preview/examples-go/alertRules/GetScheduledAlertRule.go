package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/GetScheduledAlertRule.json
func ExampleAlertRulesClient_Get_getAScheduledAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertRulesClient().Get(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.AlertRulesClientGetResponse{
	// 	                            AlertRuleClassification: &armsecurityinsights.ScheduledAlertRule{
	// 		Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/alertRules"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.AlertRuleKindScheduled),
	// 		Properties: &armsecurityinsights.ScheduledAlertRuleProperties{
	// 			AlertDetailsOverride: &armsecurityinsights.AlertDetailsOverride{
	// 				AlertDescriptionFormat: to.Ptr("Suspicious activity was made by {{ComputerIP}}"),
	// 				AlertDisplayNameFormat: to.Ptr("Alert from {{Computer}}"),
	// 			},
	// 			CustomDetails: map[string]*string{
	// 				"OperatingSystemName": to.Ptr("OSName"),
	// 				"OperatingSystemType": to.Ptr("OSType"),
	// 			},
	// 			EntityMappings: []*armsecurityinsights.EntityMapping{
	// 				{
	// 					EntityType: to.Ptr(armsecurityinsights.EntityMappingTypeHost),
	// 					FieldMappings: []*armsecurityinsights.FieldMapping{
	// 						{
	// 							ColumnName: to.Ptr("Computer"),
	// 							Identifier: to.Ptr("FullName"),
	// 					}},
	// 				},
	// 				{
	// 					EntityType: to.Ptr(armsecurityinsights.EntityMappingTypeIP),
	// 					FieldMappings: []*armsecurityinsights.FieldMapping{
	// 						{
	// 							ColumnName: to.Ptr("ComputerIP"),
	// 							Identifier: to.Ptr("Address"),
	// 					}},
	// 			}},
	// 			EventGroupingSettings: &armsecurityinsights.EventGroupingSettings{
	// 				AggregationKind: to.Ptr(armsecurityinsights.EventGroupingAggregationKindAlertPerResult),
	// 			},
	// 			Query: to.Ptr("Heartbeat"),
	// 			QueryFrequency: to.Ptr("PT1H"),
	// 			QueryPeriod: to.Ptr("P2DT1H30M"),
	// 			Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
	// 			TriggerOperator: to.Ptr(armsecurityinsights.TriggerOperatorGreaterThan),
	// 			TriggerThreshold: to.Ptr[int32](0),
	// 			Description: to.Ptr("An example for a scheduled rule"),
	// 			DisplayName: to.Ptr("My scheduled rule"),
	// 			Enabled: to.Ptr(true),
	// 			IncidentConfiguration: &armsecurityinsights.IncidentConfiguration{
	// 				CreateIncident: to.Ptr(true),
	// 				GroupingConfiguration: &armsecurityinsights.GroupingConfiguration{
	// 					Enabled: to.Ptr(true),
	// 					GroupByAlertDetails: []*armsecurityinsights.AlertDetail{
	// 						to.Ptr(armsecurityinsights.AlertDetailDisplayName)},
	// 						GroupByCustomDetails: []*string{
	// 							to.Ptr("OperatingSystemType"),
	// 							to.Ptr("OperatingSystemName")},
	// 							GroupByEntities: []*armsecurityinsights.EntityMappingType{
	// 								to.Ptr(armsecurityinsights.EntityMappingTypeHost)},
	// 								LookbackDuration: to.Ptr("PT5H"),
	// 								MatchingMethod: to.Ptr(armsecurityinsights.MatchingMethodSelected),
	// 								ReopenClosedIncident: to.Ptr(false),
	// 							},
	// 						},
	// 						LastModifiedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-01T13:17:30.000Z"); return t}()),
	// 						SuppressionDuration: to.Ptr("PT1H"),
	// 						SuppressionEnabled: to.Ptr(false),
	// 						Tactics: []*armsecurityinsights.AttackTactic{
	// 							to.Ptr(armsecurityinsights.AttackTacticPersistence),
	// 							to.Ptr(armsecurityinsights.AttackTacticLateralMovement)},
	// 							Techniques: []*string{
	// 								to.Ptr("T1037"),
	// 								to.Ptr("T1021")},
	// 							},
	// 						},
	// 						                        }
}
