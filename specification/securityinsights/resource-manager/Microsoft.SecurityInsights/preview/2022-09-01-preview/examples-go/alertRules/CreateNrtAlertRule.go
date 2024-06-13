package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/CreateNrtAlertRule.json
func ExampleAlertRulesClient_CreateOrUpdate_createsOrUpdatesANrtAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertRulesClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", &armsecurityinsights.NrtAlertRule{
		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		Kind: to.Ptr(armsecurityinsights.AlertRuleKindNRT),
		Properties: &armsecurityinsights.NrtAlertRuleProperties{
			Description: to.Ptr(""),
			DisplayName: to.Ptr("Rule2"),
			Enabled:     to.Ptr(true),
			EventGroupingSettings: &armsecurityinsights.EventGroupingSettings{
				AggregationKind: to.Ptr(armsecurityinsights.EventGroupingAggregationKindAlertPerResult),
			},
			IncidentConfiguration: &armsecurityinsights.IncidentConfiguration{
				CreateIncident: to.Ptr(true),
				GroupingConfiguration: &armsecurityinsights.GroupingConfiguration{
					Enabled: to.Ptr(true),
					GroupByEntities: []*armsecurityinsights.EntityMappingType{
						to.Ptr(armsecurityinsights.EntityMappingTypeHost),
						to.Ptr(armsecurityinsights.EntityMappingTypeAccount)},
					LookbackDuration:     to.Ptr("PT5H"),
					MatchingMethod:       to.Ptr(armsecurityinsights.MatchingMethodSelected),
					ReopenClosedIncident: to.Ptr(false),
				},
			},
			Query:               to.Ptr("ProtectionStatus | extend HostCustomEntity = Computer | extend IPCustomEntity = ComputerIP_Hidden"),
			Severity:            to.Ptr(armsecurityinsights.AlertSeverityHigh),
			SuppressionDuration: to.Ptr("PT1H"),
			SuppressionEnabled:  to.Ptr(false),
			Tactics: []*armsecurityinsights.AttackTactic{
				to.Ptr(armsecurityinsights.AttackTacticPersistence),
				to.Ptr(armsecurityinsights.AttackTacticLateralMovement)},
			Techniques: []*string{
				to.Ptr("T1037"),
				to.Ptr("T1021")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.AlertRulesClientCreateOrUpdateResponse{
	// 	                            AlertRuleClassification: &armsecurityinsights.NrtAlertRule{
	// 		Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/alertRules"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.AlertRuleKindNRT),
	// 		Properties: &armsecurityinsights.NrtAlertRuleProperties{
	// 			Description: to.Ptr(""),
	// 			DisplayName: to.Ptr("Rule2"),
	// 			Enabled: to.Ptr(true),
	// 			EventGroupingSettings: &armsecurityinsights.EventGroupingSettings{
	// 				AggregationKind: to.Ptr(armsecurityinsights.EventGroupingAggregationKindAlertPerResult),
	// 			},
	// 			IncidentConfiguration: &armsecurityinsights.IncidentConfiguration{
	// 				CreateIncident: to.Ptr(true),
	// 				GroupingConfiguration: &armsecurityinsights.GroupingConfiguration{
	// 					Enabled: to.Ptr(true),
	// 					GroupByEntities: []*armsecurityinsights.EntityMappingType{
	// 						to.Ptr(armsecurityinsights.EntityMappingTypeHost),
	// 						to.Ptr(armsecurityinsights.EntityMappingTypeAccount)},
	// 						LookbackDuration: to.Ptr("PT5H"),
	// 						MatchingMethod: to.Ptr(armsecurityinsights.MatchingMethodSelected),
	// 						ReopenClosedIncident: to.Ptr(false),
	// 					},
	// 				},
	// 				LastModifiedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:15:30.000Z"); return t}()),
	// 				Query: to.Ptr("ProtectionStatus | extend HostCustomEntity = Computer | extend IPCustomEntity = ComputerIP_Hidden"),
	// 				Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
	// 				SuppressionDuration: to.Ptr("PT1H"),
	// 				SuppressionEnabled: to.Ptr(false),
	// 				Tactics: []*armsecurityinsights.AttackTactic{
	// 					to.Ptr(armsecurityinsights.AttackTacticPersistence),
	// 					to.Ptr(armsecurityinsights.AttackTacticLateralMovement)},
	// 					Techniques: []*string{
	// 						to.Ptr("T1037"),
	// 						to.Ptr("T1021")},
	// 					},
	// 				},
	// 				                        }
}
