package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a436672b07fb1fe276c203b086b3f0e0d0c4aa24/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_Patch.json
func ExampleAlertProcessingRulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertProcessingRulesClient().Update(ctx, "alertscorrelationrg", "WeeklySuppression", armalertsmanagement.PatchObject{
		Properties: &armalertsmanagement.PatchProperties{
			Enabled: to.Ptr(false),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AlertProcessingRule = armalertsmanagement.AlertProcessingRule{
	// 	Name: to.Ptr("WeeklySuppression"),
	// 	Type: to.Ptr("Microsoft.AlertsManagement/actionRules"),
	// 	ID: to.Ptr("/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/resourceGroups/alertscorrelationrg/providers/Microsoft.AlertsManagement/actionRules/WeeklySuppression"),
	// 	Location: to.Ptr("Global"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armalertsmanagement.AlertProcessingRuleProperties{
	// 		Description: to.Ptr("Alert processing rule on resource group for weekly suppression"),
	// 		Actions: []armalertsmanagement.ActionClassification{
	// 			&armalertsmanagement.RemoveAllActionGroups{
	// 				ActionType: to.Ptr(armalertsmanagement.ActionTypeRemoveAllActionGroups),
	// 		}},
	// 		Conditions: []*armalertsmanagement.Condition{
	// 			{
	// 				Field: to.Ptr(armalertsmanagement.FieldSeverity),
	// 				Operator: to.Ptr(armalertsmanagement.OperatorEquals),
	// 				Values: []*string{
	// 					to.Ptr("Sev0"),
	// 					to.Ptr("Sev2")},
	// 				},
	// 				{
	// 					Field: to.Ptr(armalertsmanagement.FieldMonitorService),
	// 					Operator: to.Ptr(armalertsmanagement.OperatorEquals),
	// 					Values: []*string{
	// 						to.Ptr("Platform"),
	// 						to.Ptr("Application Insights")},
	// 					},
	// 					{
	// 						Field: to.Ptr(armalertsmanagement.FieldTargetResourceType),
	// 						Operator: to.Ptr(armalertsmanagement.OperatorNotEquals),
	// 						Values: []*string{
	// 							to.Ptr("Microsoft.Compute/VirtualMachines")},
	// 					}},
	// 					Enabled: to.Ptr(false),
	// 					Schedule: &armalertsmanagement.Schedule{
	// 						EffectiveFrom: to.Ptr("2018-12-10T22:05:09"),
	// 						EffectiveUntil: to.Ptr("2018-12-05T22:05:09"),
	// 						Recurrences: []armalertsmanagement.RecurrenceClassification{
	// 							&armalertsmanagement.WeeklyRecurrence{
	// 								EndTime: to.Ptr("14:00:00"),
	// 								RecurrenceType: to.Ptr(armalertsmanagement.RecurrenceTypeWeekly),
	// 								StartTime: to.Ptr("06:00:00"),
	// 								DaysOfWeek: []*armalertsmanagement.DaysOfWeek{
	// 									to.Ptr(armalertsmanagement.DaysOfWeekTuesday),
	// 									to.Ptr(armalertsmanagement.DaysOfWeekFriday)},
	// 							}},
	// 							TimeZone: to.Ptr("Pacific Standard Time"),
	// 						},
	// 						Scopes: []*string{
	// 							to.Ptr("/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/resourceGroups/alertscorrelationrg")},
	// 						},
	// 						SystemData: &armalertsmanagement.SystemData{
	// 							CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09.000Z"); return t}()),
	// 							CreatedBy: to.Ptr("abc@microsoft.com"),
	// 							CreatedByType: to.Ptr(armalertsmanagement.CreatedByTypeUser),
	// 							LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09.000Z"); return t}()),
	// 							LastModifiedBy: to.Ptr("xyz@microsoft.com"),
	// 							LastModifiedByType: to.Ptr(armalertsmanagement.CreatedByTypeUser),
	// 						},
	// 					}
}
