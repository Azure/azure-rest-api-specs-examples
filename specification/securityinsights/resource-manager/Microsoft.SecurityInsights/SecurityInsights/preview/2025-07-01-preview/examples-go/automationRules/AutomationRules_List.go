package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/automationRules/AutomationRules_List.json
func ExampleAutomationRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAutomationRulesClient().NewListPager("myRg", "myWorkspace", nil)
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
		// page = armsecurityinsights.AutomationRulesClientListResponse{
		// 	AutomationRulesList: armsecurityinsights.AutomationRulesList{
		// 		Value: []*armsecurityinsights.AutomationRule{
		// 			{
		// 				Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 				Type: to.Ptr("Microsoft.SecurityInsights/automationRules"),
		// 				Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 				ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/automationRules/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 				Properties: &armsecurityinsights.AutomationRuleProperties{
		// 					Actions: []armsecurityinsights.AutomationRuleActionClassification{
		// 						&armsecurityinsights.AutomationRuleAddIncidentTaskAction{
		// 							ActionConfiguration: &armsecurityinsights.AddIncidentTaskActionProperties{
		// 								Description: to.Ptr("Reset passwords for compromised users."),
		// 								Title: to.Ptr("Reset user passwords"),
		// 							},
		// 							ActionType: to.Ptr(armsecurityinsights.ActionTypeAddIncidentTask),
		// 							Order: to.Ptr[int32](1),
		// 						},
		// 					},
		// 					CreatedBy: &armsecurityinsights.ClientInfo{
		// 						Name: to.Ptr("john doe"),
		// 						Email: to.Ptr("john.doe@contoso.com"),
		// 						ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
		// 						UserPrincipalName: to.Ptr("john@contoso.com"),
		// 					},
		// 					CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:00Z"); return t}()),
		// 					DisplayName: to.Ptr("Suspicious user sign-in events"),
		// 					LastModifiedBy: &armsecurityinsights.ClientInfo{
		// 						Name: to.Ptr("john doe"),
		// 						Email: to.Ptr("john.doe@contoso.com"),
		// 						ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
		// 						UserPrincipalName: to.Ptr("john@contoso.com"),
		// 					},
		// 					LastModifiedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:30Z"); return t}()),
		// 					Order: to.Ptr[int32](1),
		// 					TriggeringLogic: &armsecurityinsights.AutomationRuleTriggeringLogic{
		// 						Conditions: []armsecurityinsights.AutomationRuleConditionClassification{
		// 							&armsecurityinsights.PropertyArrayConditionProperties{
		// 								ConditionProperties: &armsecurityinsights.AutomationRulePropertyArrayValuesCondition{
		// 									ArrayConditionType: to.Ptr(armsecurityinsights.AutomationRulePropertyArrayConditionSupportedArrayConditionTypeAnyItem),
		// 									ArrayType: to.Ptr(armsecurityinsights.AutomationRulePropertyArrayConditionSupportedArrayTypeIncidentLabels),
		// 									ItemConditions: []armsecurityinsights.AutomationRuleConditionClassification{
		// 										&armsecurityinsights.PropertyConditionProperties{
		// 											ConditionProperties: &armsecurityinsights.AutomationRulePropertyValuesCondition{
		// 												Operator: to.Ptr(armsecurityinsights.AutomationRulePropertyConditionSupportedOperatorEquals),
		// 												PropertyName: to.Ptr(armsecurityinsights.AutomationRulePropertyConditionSupportedPropertyIncidentLabel),
		// 												PropertyValues: []*string{
		// 													to.Ptr("myLabel"),
		// 												},
		// 											},
		// 											ConditionType: to.Ptr(armsecurityinsights.ConditionTypeProperty),
		// 										},
		// 									},
		// 								},
		// 								ConditionType: to.Ptr(armsecurityinsights.ConditionTypePropertyArray),
		// 							},
		// 						},
		// 						IsEnabled: to.Ptr(true),
		// 						TriggersOn: to.Ptr(armsecurityinsights.TriggersOnIncidents),
		// 						TriggersWhen: to.Ptr(armsecurityinsights.TriggersWhenCreated),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
