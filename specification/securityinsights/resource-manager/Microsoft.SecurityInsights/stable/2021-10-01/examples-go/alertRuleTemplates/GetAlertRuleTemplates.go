package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/alertRuleTemplates/GetAlertRuleTemplates.json
func ExampleAlertRuleTemplatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertRuleTemplatesClient().NewListPager("myRg", "myWorkspace", nil)
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
		// page.AlertRuleTemplatesList = armsecurityinsights.AlertRuleTemplatesList{
		// 	Value: []armsecurityinsights.AlertRuleTemplateClassification{
		// 		&armsecurityinsights.ScheduledAlertRuleTemplate{
		// 			Name: to.Ptr("65360bb0-8986-4ade-a89d-af3cf44d28aa"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/AlertRuleTemplates"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/AlertRuleTemplates/65360bb0-8986-4ade-a89d-af3cf44d28aa"),
		// 			Kind: to.Ptr(armsecurityinsights.AlertRuleKindScheduled),
		// 			Properties: &armsecurityinsights.ScheduledAlertRuleTemplateProperties{
		// 				Description: to.Ptr("This alert monitors changes to Amazon VPC (Virtual Private Cloud) settings such as new ACL entries and routes in route tables.\nMore information: https://medium.com/@GorillaStack/the-most-important-aws-cloudtrail-security-events-to-track-a5b9873f8255 \nand https://aws.amazon.com/vpc/"),
		// 				AlertRulesCreatedByTemplateCount: to.Ptr[int32](0),
		// 				CreatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-02-27T00:00:00.000Z"); return t}()),
		// 				DisplayName: to.Ptr("Changes to Amazon VPC settings"),
		// 				LastUpdatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-27T10:00:00.000Z"); return t}()),
		// 				Query: to.Ptr("let timeframe = 1d;\nAWSCloudTrail\n| where TimeGenerated >= ago(timeframe)\n| where EventName == \"CreateNetworkAclEntry\"\n    or EventName == \"CreateRoute\"\n| project TimeGenerated, EventName, EventTypeName, UserIdentityAccountId, UserIdentityPrincipalid, UserAgent, UserIdentityUserName, SessionMfaAuthenticated, SourceIpAddress, AWSRegion, EventSource, AdditionalEventData, ResponseElements\n| extend AccountCustomEntity = UserIdentityUserName, IPCustomEntity = SourceIpAddress"),
		// 				QueryFrequency: to.Ptr("P1D"),
		// 				QueryPeriod: to.Ptr("P1D"),
		// 				RequiredDataConnectors: []*armsecurityinsights.AlertRuleTemplateDataSource{
		// 					{
		// 						ConnectorID: to.Ptr("AWS"),
		// 						DataTypes: []*string{
		// 							to.Ptr("AWSCloudTrail")},
		// 					}},
		// 					Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 					Status: to.Ptr(armsecurityinsights.TemplateStatusAvailable),
		// 					Tactics: []*armsecurityinsights.AttackTactic{
		// 						to.Ptr(armsecurityinsights.AttackTacticPrivilegeEscalation),
		// 						to.Ptr(armsecurityinsights.AttackTacticLateralMovement)},
		// 						TriggerOperator: to.Ptr(armsecurityinsights.TriggerOperatorGreaterThan),
		// 						TriggerThreshold: to.Ptr[int32](0),
		// 						Version: to.Ptr("1.0.1"),
		// 					},
		// 				},
		// 				&armsecurityinsights.FusionAlertRuleTemplate{
		// 					Name: to.Ptr("f71aba3d-28fb-450b-b192-4e76a83015c8"),
		// 					Type: to.Ptr("Microsoft.SecurityInsights/AlertRuleTemplates"),
		// 					ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/AlertRuleTemplates/f71aba3d-28fb-450b-b192-4e76a83015c8"),
		// 					Kind: to.Ptr(armsecurityinsights.AlertRuleKindFusion),
		// 					Properties: &armsecurityinsights.FusionAlertRuleTemplateProperties{
		// 						Description: to.Ptr("Place holder: Fusion uses graph powered machine learning algorithms to correlate between millions of lower fidelity anomalous activities from different products such as Azure AD Identity Protection, and Microsoft Cloud App Security, to combine them into a manageable number of interesting security cases.\n"),
		// 						AlertRulesCreatedByTemplateCount: to.Ptr[int32](0),
		// 						CreatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-25T00:00:00.000Z"); return t}()),
		// 						DisplayName: to.Ptr("Advanced Multi-Stage Attack Detection"),
		// 						LastUpdatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-27T10:00:00.000Z"); return t}()),
		// 						Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
		// 						Status: to.Ptr(armsecurityinsights.TemplateStatusAvailable),
		// 						Tactics: []*armsecurityinsights.AttackTactic{
		// 							to.Ptr(armsecurityinsights.AttackTacticPersistence),
		// 							to.Ptr(armsecurityinsights.AttackTacticLateralMovement),
		// 							to.Ptr(armsecurityinsights.AttackTacticExfiltration),
		// 							to.Ptr(armsecurityinsights.AttackTacticCommandAndControl)},
		// 						},
		// 					},
		// 					&armsecurityinsights.MicrosoftSecurityIncidentCreationAlertRuleTemplate{
		// 						Name: to.Ptr("b3cfc7c0-092c-481c-a55b-34a3979758cb"),
		// 						Type: to.Ptr("Microsoft.SecurityInsights/AlertRuleTemplates"),
		// 						ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/AlertRuleTemplates/b3cfc7c0-092c-481c-a55b-34a3979758cb"),
		// 						Kind: to.Ptr(armsecurityinsights.AlertRuleKindMicrosoftSecurityIncidentCreation),
		// 						Properties: &armsecurityinsights.MicrosoftSecurityIncidentCreationAlertRuleTemplateProperties{
		// 							Description: to.Ptr("Create incidents based on all alerts generated in Microsoft Cloud App Security"),
		// 							AlertRulesCreatedByTemplateCount: to.Ptr[int32](0),
		// 							CreatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-16T00:00:00.000Z"); return t}()),
		// 							DisplayName: to.Ptr("Create incidents based on Microsoft Cloud App Security alerts"),
		// 							LastUpdatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-27T10:00:00.000Z"); return t}()),
		// 							ProductFilter: to.Ptr(armsecurityinsights.MicrosoftSecurityProductNameMicrosoftCloudAppSecurity),
		// 							Status: to.Ptr(armsecurityinsights.TemplateStatusAvailable),
		// 						},
		// 				}},
		// 			}
	}
}
