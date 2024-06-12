package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRuleTemplates/GetAlertRuleTemplateById.json
func ExampleAlertRuleTemplatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertRuleTemplatesClient().Get(ctx, "myRg", "myWorkspace", "65360bb0-8986-4ade-a89d-af3cf44d28aa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.AlertRuleTemplatesClientGetResponse{
	// 	                            AlertRuleTemplateClassification: &armsecurityinsights.ScheduledAlertRuleTemplate{
	// 		Name: to.Ptr("65360bb0-8986-4ade-a89d-af3cf44d28aa"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/AlertRuleTemplates"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRuleTemplates/65360bb0-8986-4ade-a89d-af3cf44d28aa"),
	// 		Kind: to.Ptr(armsecurityinsights.AlertRuleKindScheduled),
	// 		Properties: &armsecurityinsights.ScheduledAlertRuleTemplateProperties{
	// 			Description: to.Ptr("This alert monitors changes to Amazon VPC (Virtual Private Cloud) settings such as new ACL entries and routes in route tables.\nMore information: https://medium.com/@GorillaStack/the-most-important-aws-cloudtrail-security-events-to-track-a5b9873f8255 \nand https://aws.amazon.com/vpc/"),
	// 			AlertRulesCreatedByTemplateCount: to.Ptr[int32](0),
	// 			CreatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-02-27T00:00:00.000Z"); return t}()),
	// 			DisplayName: to.Ptr("Changes to Amazon VPC settings"),
	// 			EventGroupingSettings: &armsecurityinsights.EventGroupingSettings{
	// 				AggregationKind: to.Ptr(armsecurityinsights.EventGroupingAggregationKindAlertPerResult),
	// 			},
	// 			LastUpdatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-27T00:00:00.000Z"); return t}()),
	// 			Query: to.Ptr("let timeframe = 1d;\nAWSCloudTrail\n| where TimeGenerated >= ago(timeframe)\n| where EventName == \"CreateNetworkAclEntry\"\n    or EventName == \"CreateRoute\"\n| project TimeGenerated, EventName, EventTypeName, UserIdentityAccountId, UserIdentityPrincipalid, UserAgent, UserIdentityUserName, SessionMfaAuthenticated, SourceIpAddress, AWSRegion, EventSource, AdditionalEventData, ResponseElements\n| extend AccountCustomEntity = UserIdentityUserName, IPCustomEntity = SourceIpAddress"),
	// 			QueryFrequency: to.Ptr("P1D"),
	// 			QueryPeriod: to.Ptr("P1D"),
	// 			RequiredDataConnectors: []*armsecurityinsights.AlertRuleTemplateDataSource{
	// 				{
	// 					ConnectorID: to.Ptr("AWS"),
	// 					DataTypes: []*string{
	// 						to.Ptr("AWSCloudTrail")},
	// 				}},
	// 				Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
	// 				Status: to.Ptr(armsecurityinsights.TemplateStatusAvailable),
	// 				Tactics: []*armsecurityinsights.AttackTactic{
	// 					to.Ptr(armsecurityinsights.AttackTacticPrivilegeEscalation),
	// 					to.Ptr(armsecurityinsights.AttackTacticLateralMovement)},
	// 					Techniques: []*string{
	// 						to.Ptr("T1037"),
	// 						to.Ptr("T1021")},
	// 						TriggerOperator: to.Ptr(armsecurityinsights.TriggerOperatorGreaterThan),
	// 						TriggerThreshold: to.Ptr[int32](0),
	// 						Version: to.Ptr("1.0.2"),
	// 					},
	// 				},
	// 				                        }
}
