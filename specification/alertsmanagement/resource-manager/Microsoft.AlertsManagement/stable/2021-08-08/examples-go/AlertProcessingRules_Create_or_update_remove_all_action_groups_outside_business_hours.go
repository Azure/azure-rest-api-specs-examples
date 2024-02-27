package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6d2438481021a94793b07b226df06d5f3c61d51d/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_Create_or_update_remove_all_action_groups_outside_business_hours.json
func ExampleAlertProcessingRulesClient_CreateOrUpdate_createOrUpdateARuleThatRemovesAllActionGroupsOutsideBusinessHoursMonFri09001700EasternStandardTime() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertProcessingRulesClient().CreateOrUpdate(ctx, "alertscorrelationrg", "RemoveActionGroupsOutsideBusinessHours", armalertsmanagement.AlertProcessingRule{
		Location: to.Ptr("Global"),
		Tags:     map[string]*string{},
		Properties: &armalertsmanagement.AlertProcessingRuleProperties{
			Description: to.Ptr("Remove all ActionGroups outside business hours"),
			Actions: []armalertsmanagement.ActionClassification{
				&armalertsmanagement.RemoveAllActionGroups{
					ActionType: to.Ptr(armalertsmanagement.ActionTypeRemoveAllActionGroups),
				}},
			Enabled: to.Ptr(true),
			Schedule: &armalertsmanagement.Schedule{
				Recurrences: []armalertsmanagement.RecurrenceClassification{
					&armalertsmanagement.DailyRecurrence{
						EndTime:        to.Ptr("09:00:00"),
						RecurrenceType: to.Ptr(armalertsmanagement.RecurrenceTypeDaily),
						StartTime:      to.Ptr("17:00:00"),
					},
					&armalertsmanagement.WeeklyRecurrence{
						RecurrenceType: to.Ptr(armalertsmanagement.RecurrenceTypeWeekly),
						DaysOfWeek: []*armalertsmanagement.DaysOfWeek{
							to.Ptr(armalertsmanagement.DaysOfWeekSaturday),
							to.Ptr(armalertsmanagement.DaysOfWeekSunday)},
					}},
				TimeZone: to.Ptr("Eastern Standard Time"),
			},
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
	// 	Name: to.Ptr("RemoveActionGroupsOutsideBusinessHours"),
	// 	Type: to.Ptr("Microsoft.AlertsManagement/actionRules"),
	// 	ID: to.Ptr("/subscriptions/subId1/resourceGroups/alertscorrelationrg/providers/Microsoft.AlertsManagement/actionRules/RemoveActionGroupsOutsideBusinessHours"),
	// 	Location: to.Ptr("Global"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armalertsmanagement.AlertProcessingRuleProperties{
	// 		Description: to.Ptr("Remove all ActionGroups outside business hours"),
	// 		Actions: []armalertsmanagement.ActionClassification{
	// 			&armalertsmanagement.RemoveAllActionGroups{
	// 				ActionType: to.Ptr(armalertsmanagement.ActionTypeRemoveAllActionGroups),
	// 		}},
	// 		Enabled: to.Ptr(true),
	// 		Schedule: &armalertsmanagement.Schedule{
	// 			Recurrences: []armalertsmanagement.RecurrenceClassification{
	// 				&armalertsmanagement.DailyRecurrence{
	// 					EndTime: to.Ptr("09:00:00"),
	// 					RecurrenceType: to.Ptr(armalertsmanagement.RecurrenceTypeDaily),
	// 					StartTime: to.Ptr("17:00:00"),
	// 				},
	// 				&armalertsmanagement.WeeklyRecurrence{
	// 					RecurrenceType: to.Ptr(armalertsmanagement.RecurrenceTypeWeekly),
	// 					DaysOfWeek: []*armalertsmanagement.DaysOfWeek{
	// 						to.Ptr(armalertsmanagement.DaysOfWeekSaturday),
	// 						to.Ptr(armalertsmanagement.DaysOfWeekSunday)},
	// 				}},
	// 				TimeZone: to.Ptr("Eastern Standard Time"),
	// 			},
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
