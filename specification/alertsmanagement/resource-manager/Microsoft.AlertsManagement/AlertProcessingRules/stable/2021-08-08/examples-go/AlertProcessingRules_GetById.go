package armalertprocessingrules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertprocessingrules/armalertprocessingrules"
)

// Generated from example definition: 2021-08-08/AlertProcessingRules_GetById.json
func ExampleClient_GetByName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertprocessingrules.NewClientFactory("1e3ff1c0-771a-4119-a03b-be82a51e232d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().GetByName(ctx, "alertscorrelationrg", "DailySuppression", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armalertprocessingrules.ClientGetByNameResponse{
	// 	AlertProcessingRule: &armalertprocessingrules.AlertProcessingRule{
	// 		Name: to.Ptr("DailySuppression"),
	// 		Type: to.Ptr("Microsoft.AlertsManagement/actionRules"),
	// 		ID: to.Ptr("/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/resourceGroups/alertscorrelationrg/providers/Microsoft.AlertsManagement/actionRules/DailySuppression"),
	// 		Location: to.Ptr("Global"),
	// 		Properties: &armalertprocessingrules.AlertProcessingRuleProperties{
	// 			Description: to.Ptr("Alert processing rule on resource group for daily and weekly scheduling"),
	// 			Actions: []armalertprocessingrules.ActionClassification{
	// 				&armalertprocessingrules.AddActionGroups{
	// 					ActionGroupIDs: []*string{
	// 						to.Ptr("actiongGroup1"),
	// 						to.Ptr("actiongGroup2"),
	// 					},
	// 					ActionType: to.Ptr(armalertprocessingrules.ActionTypeAddActionGroups),
	// 				},
	// 			},
	// 			Conditions: []*armalertprocessingrules.Condition{
	// 				{
	// 					Field: to.Ptr(armalertprocessingrules.FieldSeverity),
	// 					Operator: to.Ptr(armalertprocessingrules.OperatorEquals),
	// 					Values: []*string{
	// 						to.Ptr("Sev0"),
	// 						to.Ptr("Sev2"),
	// 					},
	// 				},
	// 				{
	// 					Field: to.Ptr(armalertprocessingrules.FieldMonitorService),
	// 					Operator: to.Ptr(armalertprocessingrules.OperatorEquals),
	// 					Values: []*string{
	// 						to.Ptr("Platform"),
	// 						to.Ptr("Application Insights"),
	// 					},
	// 				},
	// 				{
	// 					Field: to.Ptr(armalertprocessingrules.FieldMonitorCondition),
	// 					Operator: to.Ptr(armalertprocessingrules.OperatorEquals),
	// 					Values: []*string{
	// 						to.Ptr("Fired"),
	// 					},
	// 				},
	// 				{
	// 					Field: to.Ptr(armalertprocessingrules.FieldTargetResourceType),
	// 					Operator: to.Ptr(armalertprocessingrules.OperatorNotEquals),
	// 					Values: []*string{
	// 						to.Ptr("Microsoft.Compute/VirtualMachines"),
	// 					},
	// 				},
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			Schedule: &armalertprocessingrules.Schedule{
	// 				EffectiveFrom: to.Ptr("2018-01-10T22:05:09"),
	// 				EffectiveUntil: to.Ptr("2018-12-10T22:05:09"),
	// 				Recurrences: []armalertprocessingrules.RecurrenceClassification{
	// 					&armalertprocessingrules.DailyRecurrence{
	// 						EndTime: to.Ptr("14:00:00"),
	// 						RecurrenceType: to.Ptr(armalertprocessingrules.RecurrenceTypeDaily),
	// 						StartTime: to.Ptr("06:00:00"),
	// 					},
	// 					&armalertprocessingrules.WeeklyRecurrence{
	// 						DaysOfWeek: []*armalertprocessingrules.DaysOfWeek{
	// 							to.Ptr(armalertprocessingrules.DaysOfWeekSaturday),
	// 							to.Ptr(armalertprocessingrules.DaysOfWeekSunday),
	// 						},
	// 						EndTime: to.Ptr("20:00:00"),
	// 						RecurrenceType: to.Ptr(armalertprocessingrules.RecurrenceTypeWeekly),
	// 						StartTime: to.Ptr("10:00:00"),
	// 					},
	// 				},
	// 				TimeZone: to.Ptr("Pacific Standard Time"),
	// 			},
	// 			Scopes: []*string{
	// 				to.Ptr("/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/resourceGroups/alertscorrelationrg"),
	// 			},
	// 		},
	// 		SystemData: &armalertprocessingrules.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09Z"); return t}()),
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
