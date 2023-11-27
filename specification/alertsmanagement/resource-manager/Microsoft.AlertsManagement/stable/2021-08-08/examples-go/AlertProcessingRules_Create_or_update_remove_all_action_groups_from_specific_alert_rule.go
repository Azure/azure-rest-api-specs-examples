package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a436672b07fb1fe276c203b086b3f0e0d0c4aa24/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_Create_or_update_remove_all_action_groups_from_specific_alert_rule.json
func ExampleAlertProcessingRulesClient_CreateOrUpdate_createOrUpdateARuleThatRemovesAllActionGroupsFromAllAlertsInASubscriptionComingFromASpecificAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertProcessingRulesClient().CreateOrUpdate(ctx, "alertscorrelationrg", "RemoveActionGroupsSpecificAlertRule", armalertsmanagement.AlertProcessingRule{
		Location: to.Ptr("Global"),
		Tags:     map[string]*string{},
		Properties: &armalertsmanagement.AlertProcessingRuleProperties{
			Description: to.Ptr("Removes all ActionGroups from all Alerts that fire on above AlertRule"),
			Actions: []armalertsmanagement.ActionClassification{
				&armalertsmanagement.RemoveAllActionGroups{
					ActionType: to.Ptr(armalertsmanagement.ActionTypeRemoveAllActionGroups),
				}},
			Conditions: []*armalertsmanagement.Condition{
				{
					Field:    to.Ptr(armalertsmanagement.FieldAlertRuleID),
					Operator: to.Ptr(armalertsmanagement.OperatorEquals),
					Values: []*string{
						to.Ptr("/subscriptions/suubId1/resourceGroups/Rgid2/providers/microsoft.insights/activityLogAlerts/RuleName")},
				}},
			Enabled: to.Ptr(true),
			Scopes: []*string{
				to.Ptr("/subscriptions/subId1")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AlertProcessingRule = armalertsmanagement.AlertProcessingRule{
	// 	Name: to.Ptr("RemoveActionGroupsSpecificAlertRule"),
	// 	Type: to.Ptr("Microsoft.AlertsManagement/actionRules"),
	// 	ID: to.Ptr("/subscriptions/subId1/resourceGroups/alertscorrelationrg/providers/Microsoft.AlertsManagement/actionRules/RemoveActionGroupsSpecificAlertRule"),
	// 	Location: to.Ptr("Global"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armalertsmanagement.AlertProcessingRuleProperties{
	// 		Description: to.Ptr("Removes all ActionGroups from all Alerts that fire on above AlertRule"),
	// 		Actions: []armalertsmanagement.ActionClassification{
	// 			&armalertsmanagement.RemoveAllActionGroups{
	// 				ActionType: to.Ptr(armalertsmanagement.ActionTypeRemoveAllActionGroups),
	// 		}},
	// 		Conditions: []*armalertsmanagement.Condition{
	// 			{
	// 				Field: to.Ptr(armalertsmanagement.FieldAlertRuleID),
	// 				Operator: to.Ptr(armalertsmanagement.OperatorEquals),
	// 				Values: []*string{
	// 					to.Ptr("/subscriptions/suubId1/resourceGroups/Rgid2/providers/microsoft.insights/activityLogAlerts/RuleName")},
	// 			}},
	// 			Enabled: to.Ptr(true),
	// 			Scopes: []*string{
	// 				to.Ptr("/subscriptions/subId1")},
	// 			},
	// 			SystemData: &armalertsmanagement.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T22:05:09.000Z"); return t}()),
	// 				CreatedBy: to.Ptr("abc@microsoft.com"),
	// 				CreatedByType: to.Ptr(armalertsmanagement.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09.000Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("xyz@microsoft.com"),
	// 				LastModifiedByType: to.Ptr(armalertsmanagement.CreatedByTypeUser),
	// 			},
	// 		}
}
