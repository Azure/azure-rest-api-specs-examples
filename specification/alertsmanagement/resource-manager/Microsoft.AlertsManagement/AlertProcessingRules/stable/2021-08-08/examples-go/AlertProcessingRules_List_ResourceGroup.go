package armalertprocessingrules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertprocessingrules/armalertprocessingrules"
)

// Generated from example definition: 2021-08-08/AlertProcessingRules_List_ResourceGroup.json
func ExampleClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertprocessingrules.NewClientFactory("1e3ff1c0-771a-4119-a03b-be82a51e232d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListByResourceGroupPager("alertscorrelationrg", nil)
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
		// page = armalertprocessingrules.ClientListByResourceGroupResponse{
		// 	List: armalertprocessingrules.List{
		// 		NextLink: to.Ptr("https://management.azure.com:443/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/resourceGroups/alertscorrelationrg/providers/Microsoft.AlertsManagement/actionRules?api-version=2021-08-08&ctoken=%2bRID%3aPlwOAPHEGwB9UwEAAAAgCw%3d%3d%23RT%3a2%23TRC%3a500%23RTD%3aqtQyMDE4LTA2LTEyVDE1OjEyOjE1"),
		// 		Value: []*armalertprocessingrules.AlertProcessingRule{
		// 			{
		// 				Name: to.Ptr("DailySuppression"),
		// 				Type: to.Ptr("Microsoft.AlertsManagement/actionRules"),
		// 				ID: to.Ptr("/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/resourceGroups/alertscorrelationrg/providers/Microsoft.AlertsManagement/actionRules/DailySuppression"),
		// 				Location: to.Ptr("Global"),
		// 				Properties: &armalertprocessingrules.AlertProcessingRuleProperties{
		// 					Description: to.Ptr("Alert processing rule on resource group for daily suppression"),
		// 					Actions: []armalertprocessingrules.ActionClassification{
		// 						&armalertprocessingrules.RemoveAllActionGroups{
		// 							ActionType: to.Ptr(armalertprocessingrules.ActionTypeRemoveAllActionGroups),
		// 						},
		// 					},
		// 					Conditions: []*armalertprocessingrules.Condition{
		// 						{
		// 							Field: to.Ptr(armalertprocessingrules.FieldSeverity),
		// 							Operator: to.Ptr(armalertprocessingrules.OperatorEquals),
		// 							Values: []*string{
		// 								to.Ptr("Sev0"),
		// 								to.Ptr("Sev2"),
		// 							},
		// 						},
		// 						{
		// 							Field: to.Ptr(armalertprocessingrules.FieldMonitorService),
		// 							Operator: to.Ptr(armalertprocessingrules.OperatorEquals),
		// 							Values: []*string{
		// 								to.Ptr("Platform"),
		// 								to.Ptr("Application Insights"),
		// 							},
		// 						},
		// 						{
		// 							Field: to.Ptr(armalertprocessingrules.FieldTargetResourceType),
		// 							Operator: to.Ptr(armalertprocessingrules.OperatorNotEquals),
		// 							Values: []*string{
		// 								to.Ptr("Microsoft.Compute/VirtualMachines"),
		// 							},
		// 						},
		// 					},
		// 					Enabled: to.Ptr(true),
		// 					Schedule: &armalertprocessingrules.Schedule{
		// 						EffectiveFrom: to.Ptr("2018-09-12T06:00:00"),
		// 						EffectiveUntil: to.Ptr("2018-09-20T14:00:00"),
		// 						Recurrences: []armalertprocessingrules.RecurrenceClassification{
		// 							&armalertprocessingrules.DailyRecurrence{
		// 								EndTime: to.Ptr("14:00:00"),
		// 								RecurrenceType: to.Ptr(armalertprocessingrules.RecurrenceTypeDaily),
		// 								StartTime: to.Ptr("06:00:00"),
		// 							},
		// 						},
		// 						TimeZone: to.Ptr("Pacific Standard Time"),
		// 					},
		// 					Scopes: []*string{
		// 						to.Ptr("/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/resourceGroups/alertscorrelationrg"),
		// 					},
		// 				},
		// 				SystemData: &armalertprocessingrules.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09Z"); return t}()),
		// 					CreatedBy: to.Ptr("abc@microsoft.com"),
		// 					CreatedByType: to.Ptr(armalertprocessingrules.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("xyz@microsoft.com"),
		// 					LastModifiedByType: to.Ptr(armalertprocessingrules.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("WeeklySuppression"),
		// 				Type: to.Ptr("Microsoft.AlertsManagement/actionRules"),
		// 				ID: to.Ptr("/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/resourceGroups/alertscorrelationrg/providers/Microsoft.AlertsManagement/actionRules/WeeklySuppression"),
		// 				Location: to.Ptr("Global"),
		// 				Properties: &armalertprocessingrules.AlertProcessingRuleProperties{
		// 					Description: to.Ptr("Alert processing rule on resource group for sending email"),
		// 					Actions: []armalertprocessingrules.ActionClassification{
		// 						&armalertprocessingrules.AddActionGroups{
		// 							ActionGroupIDs: []*string{
		// 								to.Ptr("/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/resourceGroups/alertscorrelationrg/providers/Microsoft.insights/actiongroups/testAG"),
		// 							},
		// 							ActionType: to.Ptr(armalertprocessingrules.ActionTypeAddActionGroups),
		// 						},
		// 					},
		// 					Conditions: []*armalertprocessingrules.Condition{
		// 						{
		// 							Field: to.Ptr(armalertprocessingrules.FieldMonitorCondition),
		// 							Operator: to.Ptr(armalertprocessingrules.OperatorEquals),
		// 							Values: []*string{
		// 								to.Ptr("Fired"),
		// 							},
		// 						},
		// 						{
		// 							Field: to.Ptr(armalertprocessingrules.FieldDescription),
		// 							Operator: to.Ptr(armalertprocessingrules.OperatorContains),
		// 							Values: []*string{
		// 								to.Ptr("Percentage CPU greater than 80%"),
		// 								to.Ptr("Metric alert on resource foo"),
		// 							},
		// 						},
		// 						{
		// 							Field: to.Ptr(armalertprocessingrules.FieldAlertContext),
		// 							Operator: to.Ptr(armalertprocessingrules.OperatorDoesNotContain),
		// 							Values: []*string{
		// 								to.Ptr("testresource"),
		// 								to.Ptr("foo"),
		// 							},
		// 						},
		// 					},
		// 					Enabled: to.Ptr(true),
		// 					Scopes: []*string{
		// 						to.Ptr("/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/resourceGroups/alertscorrelationrg/providers/Microsoft.Compute/VirtualMachines/testResource"),
		// 					},
		// 				},
		// 				SystemData: &armalertprocessingrules.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09Z"); return t}()),
		// 					CreatedBy: to.Ptr("abc@microsoft.com"),
		// 					CreatedByType: to.Ptr(armalertprocessingrules.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("xyz@microsoft.com"),
		// 					LastModifiedByType: to.Ptr(armalertprocessingrules.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
