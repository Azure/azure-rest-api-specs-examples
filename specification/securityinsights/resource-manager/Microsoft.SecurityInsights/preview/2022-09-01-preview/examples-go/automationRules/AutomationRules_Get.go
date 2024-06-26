package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/automationRules/AutomationRules_Get.json
func ExampleAutomationRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAutomationRulesClient().Get(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AutomationRule = armsecurityinsights.AutomationRule{
	// 	Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/automationRules"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/automationRules/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 	Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 	Properties: &armsecurityinsights.AutomationRuleProperties{
	// 		Actions: []armsecurityinsights.AutomationRuleActionClassification{
	// 			&armsecurityinsights.AutomationRuleRunPlaybookAction{
	// 				ActionType: to.Ptr(armsecurityinsights.ActionTypeRunPlaybook),
	// 				Order: to.Ptr[int32](1),
	// 				ActionConfiguration: &armsecurityinsights.PlaybookActionProperties{
	// 					LogicAppResourceID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.Logic/workflows/IncidentPlaybook"),
	// 					TenantID: to.Ptr("d23e3eef-eed0-428f-a2d5-bc48c268e31d"),
	// 				},
	// 		}},
	// 		CreatedBy: &armsecurityinsights.ClientInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john.doe@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 			UserPrincipalName: to.Ptr("john@contoso.com"),
	// 		},
	// 		CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:00.000Z"); return t}()),
	// 		DisplayName: to.Ptr("Suspicious alerts in workspace"),
	// 		LastModifiedBy: &armsecurityinsights.ClientInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john.doe@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 			UserPrincipalName: to.Ptr("john@contoso.com"),
	// 		},
	// 		LastModifiedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:30.000Z"); return t}()),
	// 		Order: to.Ptr[int32](1),
	// 		TriggeringLogic: &armsecurityinsights.AutomationRuleTriggeringLogic{
	// 			Conditions: []armsecurityinsights.AutomationRuleConditionClassification{
	// 				&armsecurityinsights.BooleanConditionProperties{
	// 					ConditionType: to.Ptr(armsecurityinsights.ConditionTypeBoolean),
	// 					ConditionProperties: &armsecurityinsights.AutomationRuleBooleanCondition{
	// 						InnerConditions: []armsecurityinsights.AutomationRuleConditionClassification{
	// 							&armsecurityinsights.PropertyConditionProperties{
	// 								ConditionType: to.Ptr(armsecurityinsights.ConditionTypeProperty),
	// 								ConditionProperties: &armsecurityinsights.AutomationRulePropertyValuesCondition{
	// 									Operator: to.Ptr(armsecurityinsights.AutomationRulePropertyConditionSupportedOperatorEquals),
	// 									PropertyName: to.Ptr(armsecurityinsights.AutomationRulePropertyConditionSupportedPropertyAccountName),
	// 									PropertyValues: []*string{
	// 										to.Ptr("Administrator")},
	// 									},
	// 								},
	// 								&armsecurityinsights.PropertyConditionProperties{
	// 									ConditionType: to.Ptr(armsecurityinsights.ConditionTypeProperty),
	// 									ConditionProperties: &armsecurityinsights.AutomationRulePropertyValuesCondition{
	// 										Operator: to.Ptr(armsecurityinsights.AutomationRulePropertyConditionSupportedOperatorEquals),
	// 										PropertyName: to.Ptr(armsecurityinsights.AutomationRulePropertyConditionSupportedPropertyHostName),
	// 										PropertyValues: []*string{
	// 											to.Ptr("MainServer")},
	// 										},
	// 								}},
	// 								Operator: to.Ptr(armsecurityinsights.AutomationRuleBooleanConditionSupportedOperatorOr),
	// 							},
	// 						},
	// 						&armsecurityinsights.PropertyArrayConditionProperties{
	// 							ConditionType: to.Ptr(armsecurityinsights.ConditionTypePropertyArray),
	// 							ConditionProperties: &armsecurityinsights.AutomationRulePropertyArrayValuesCondition{
	// 								ArrayConditionType: to.Ptr(armsecurityinsights.AutomationRulePropertyArrayConditionSupportedArrayConditionTypeAnyItem),
	// 								ArrayType: to.Ptr(armsecurityinsights.AutomationRulePropertyArrayConditionSupportedArrayTypeCustomDetails),
	// 								ItemConditions: []armsecurityinsights.AutomationRuleConditionClassification{
	// 									&armsecurityinsights.PropertyConditionProperties{
	// 										ConditionType: to.Ptr(armsecurityinsights.ConditionTypeProperty),
	// 										ConditionProperties: &armsecurityinsights.AutomationRulePropertyValuesCondition{
	// 											Operator: to.Ptr(armsecurityinsights.AutomationRulePropertyConditionSupportedOperatorEquals),
	// 											PropertyName: to.Ptr(armsecurityinsights.AutomationRulePropertyConditionSupportedPropertyIncidentCustomDetailsKey),
	// 											PropertyValues: []*string{
	// 												to.Ptr("AlertTags")},
	// 											},
	// 										},
	// 										&armsecurityinsights.PropertyArrayConditionProperties{
	// 											ConditionType: to.Ptr(armsecurityinsights.ConditionTypePropertyArray),
	// 											ConditionProperties: &armsecurityinsights.AutomationRulePropertyArrayValuesCondition{
	// 												ArrayConditionType: to.Ptr(armsecurityinsights.AutomationRulePropertyArrayConditionSupportedArrayConditionTypeAnyItem),
	// 												ArrayType: to.Ptr(armsecurityinsights.AutomationRulePropertyArrayConditionSupportedArrayTypeCustomDetailValues),
	// 												ItemConditions: []armsecurityinsights.AutomationRuleConditionClassification{
	// 													&armsecurityinsights.PropertyConditionProperties{
	// 														ConditionType: to.Ptr(armsecurityinsights.ConditionTypeProperty),
	// 														ConditionProperties: &armsecurityinsights.AutomationRulePropertyValuesCondition{
	// 															Operator: to.Ptr(armsecurityinsights.AutomationRulePropertyConditionSupportedOperatorEquals),
	// 															PropertyName: to.Ptr(armsecurityinsights.AutomationRulePropertyConditionSupportedPropertyIncidentCustomDetailsValue),
	// 															PropertyValues: []*string{
	// 																to.Ptr("HighPriority")},
	// 															},
	// 													}},
	// 												},
	// 										}},
	// 									},
	// 							}},
	// 							IsEnabled: to.Ptr(true),
	// 							TriggersOn: to.Ptr(armsecurityinsights.TriggersOnIncidents),
	// 							TriggersWhen: to.Ptr(armsecurityinsights.TriggersWhenCreated),
	// 						},
	// 					},
	// 				}
}
