package armalertprocessingrules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertprocessingrules/armalertprocessingrules"
)

// Generated from example definition: 2021-08-08/AlertProcessingRules_Create_or_update_remove_all_action_groups_outside_business_hours.json
func ExampleClient_CreateOrUpdate_createOrUpdateARuleThatRemovesAllActionGroupsOutsideBusinessHoursMonFri09001700EasternStandardTime() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertprocessingrules.NewClientFactory("subId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().CreateOrUpdate(ctx, "alertscorrelationrg", "RemoveActionGroupsOutsideBusinessHours", armalertprocessingrules.AlertProcessingRule{
		Location: to.Ptr("Global"),
		Properties: &armalertprocessingrules.AlertProcessingRuleProperties{
			Description: to.Ptr("Remove all ActionGroups outside business hours"),
			Actions: []armalertprocessingrules.ActionClassification{
				&armalertprocessingrules.RemoveAllActionGroups{
					ActionType: to.Ptr(armalertprocessingrules.ActionTypeRemoveAllActionGroups),
				},
			},
			Enabled: to.Ptr(true),
			Schedule: &armalertprocessingrules.Schedule{
				Recurrences: []armalertprocessingrules.RecurrenceClassification{
					&armalertprocessingrules.DailyRecurrence{
						EndTime:        to.Ptr("09:00:00"),
						RecurrenceType: to.Ptr(armalertprocessingrules.RecurrenceTypeDaily),
						StartTime:      to.Ptr("17:00:00"),
					},
					&armalertprocessingrules.WeeklyRecurrence{
						DaysOfWeek: []*armalertprocessingrules.DaysOfWeek{
							to.Ptr(armalertprocessingrules.DaysOfWeekSaturday),
							to.Ptr(armalertprocessingrules.DaysOfWeekSunday),
						},
						RecurrenceType: to.Ptr(armalertprocessingrules.RecurrenceTypeWeekly),
					},
				},
				TimeZone: to.Ptr("Eastern Standard Time"),
			},
			Scopes: []*string{
				to.Ptr("/subscriptions/subId1"),
			},
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armalertprocessingrules.ClientCreateOrUpdateResponse{
	// 	AlertProcessingRule: &armalertprocessingrules.AlertProcessingRule{
	// 		Name: to.Ptr("RemoveActionGroupsOutsideBusinessHours"),
	// 		Type: to.Ptr("Microsoft.AlertsManagement/actionRules"),
	// 		ID: to.Ptr("/subscriptions/subId1/resourceGroups/alertscorrelationrg/providers/Microsoft.AlertsManagement/actionRules/RemoveActionGroupsOutsideBusinessHours"),
	// 		Location: to.Ptr("Global"),
	// 		Properties: &armalertprocessingrules.AlertProcessingRuleProperties{
	// 			Description: to.Ptr("Remove all ActionGroups outside business hours"),
	// 			Actions: []armalertprocessingrules.ActionClassification{
	// 				&armalertprocessingrules.RemoveAllActionGroups{
	// 					ActionType: to.Ptr(armalertprocessingrules.ActionTypeRemoveAllActionGroups),
	// 				},
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			Schedule: &armalertprocessingrules.Schedule{
	// 				Recurrences: []armalertprocessingrules.RecurrenceClassification{
	// 					&armalertprocessingrules.DailyRecurrence{
	// 						EndTime: to.Ptr("09:00:00"),
	// 						RecurrenceType: to.Ptr(armalertprocessingrules.RecurrenceTypeDaily),
	// 						StartTime: to.Ptr("17:00:00"),
	// 					},
	// 					&armalertprocessingrules.WeeklyRecurrence{
	// 						DaysOfWeek: []*armalertprocessingrules.DaysOfWeek{
	// 							to.Ptr(armalertprocessingrules.DaysOfWeekSaturday),
	// 							to.Ptr(armalertprocessingrules.DaysOfWeekSunday),
	// 						},
	// 						RecurrenceType: to.Ptr(armalertprocessingrules.RecurrenceTypeWeekly),
	// 					},
	// 				},
	// 				TimeZone: to.Ptr("Eastern Standard Time"),
	// 			},
	// 			Scopes: []*string{
	// 				to.Ptr("/subscriptions/subId1"),
	// 			},
	// 		},
	// 		SystemData: &armalertprocessingrules.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T22:05:09Z"); return t}()),
	// 			CreatedBy: to.Ptr("abc@microsoft.com"),
	// 			CreatedByType: to.Ptr(armalertprocessingrules.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("xyz@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armalertprocessingrules.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
