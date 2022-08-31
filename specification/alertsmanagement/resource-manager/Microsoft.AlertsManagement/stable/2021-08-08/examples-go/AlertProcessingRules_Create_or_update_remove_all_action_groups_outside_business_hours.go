package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_Create_or_update_remove_all_action_groups_outside_business_hours.json
func ExampleAlertProcessingRulesClient_CreateOrUpdate_createOrUpdateARuleThatRemovesAllActionGroupsOutsideBusinessHoursMonFri09001700EasternStandardTime() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armalertsmanagement.NewAlertProcessingRulesClient("subId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "alertscorrelationrg", "RemoveActionGroupsOutsideBusinessHours", armalertsmanagement.AlertProcessingRule{
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
	// TODO: use response item
	_ = res
}
