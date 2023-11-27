package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a436672b07fb1fe276c203b086b3f0e0d0c4aa24/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_List_Subscription.json
func ExampleAlertProcessingRulesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertProcessingRulesClient().NewListBySubscriptionPager(nil)
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
		// page.AlertProcessingRulesList = armalertsmanagement.AlertProcessingRulesList{
		// 	Value: []*armalertsmanagement.AlertProcessingRule{
		// 		{
		// 			Name: to.Ptr("DailySuppression"),
		// 			Type: to.Ptr("Microsoft.AlertsManagement/actionRules"),
		// 			ID: to.Ptr("/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/resourceGroups/alertscorrelationrg/providers/Microsoft.AlertsManagement/actionRules/DailySuppression"),
		// 			Location: to.Ptr("Global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armalertsmanagement.AlertProcessingRuleProperties{
		// 				Description: to.Ptr("Alert processing rule on resource group for daily suppression"),
		// 				Actions: []armalertsmanagement.ActionClassification{
		// 					&armalertsmanagement.RemoveAllActionGroups{
		// 						ActionType: to.Ptr(armalertsmanagement.ActionTypeRemoveAllActionGroups),
		// 				}},
		// 				Conditions: []*armalertsmanagement.Condition{
		// 					{
		// 						Field: to.Ptr(armalertsmanagement.FieldSeverity),
		// 						Operator: to.Ptr(armalertsmanagement.OperatorEquals),
		// 						Values: []*string{
		// 							to.Ptr("Sev0"),
		// 							to.Ptr("Sev2")},
		// 						},
		// 						{
		// 							Field: to.Ptr(armalertsmanagement.FieldMonitorService),
		// 							Operator: to.Ptr(armalertsmanagement.OperatorEquals),
		// 							Values: []*string{
		// 								to.Ptr("Platform"),
		// 								to.Ptr("Application Insights")},
		// 							},
		// 							{
		// 								Field: to.Ptr(armalertsmanagement.FieldTargetResourceType),
		// 								Operator: to.Ptr(armalertsmanagement.OperatorNotEquals),
		// 								Values: []*string{
		// 									to.Ptr("Microsoft.Compute/VirtualMachines")},
		// 							}},
		// 							Enabled: to.Ptr(true),
		// 							Schedule: &armalertsmanagement.Schedule{
		// 								EffectiveFrom: to.Ptr("2018-01-10T22:05:09"),
		// 								EffectiveUntil: to.Ptr("2018-12-10T22:05:09"),
		// 								Recurrences: []armalertsmanagement.RecurrenceClassification{
		// 									&armalertsmanagement.DailyRecurrence{
		// 										EndTime: to.Ptr("14:00:00"),
		// 										RecurrenceType: to.Ptr(armalertsmanagement.RecurrenceTypeDaily),
		// 										StartTime: to.Ptr("06:00:00"),
		// 								}},
		// 								TimeZone: to.Ptr("Pacific Standard Time"),
		// 							},
		// 							Scopes: []*string{
		// 								to.Ptr("/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/resourceGroups/alertscorrelationrg")},
		// 							},
		// 							SystemData: &armalertsmanagement.SystemData{
		// 								CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09.000Z"); return t}()),
		// 								CreatedBy: to.Ptr("abc@microsoft.com"),
		// 								CreatedByType: to.Ptr(armalertsmanagement.CreatedByTypeUser),
		// 								LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09.000Z"); return t}()),
		// 								LastModifiedBy: to.Ptr("xyz@microsoft.com"),
		// 								LastModifiedByType: to.Ptr(armalertsmanagement.CreatedByTypeUser),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("WeeklySuppression"),
		// 							Type: to.Ptr("Microsoft.AlertsManagement/actionRules"),
		// 							ID: to.Ptr("/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/resourceGroups/alertscorrelationrg/providers/Microsoft.AlertsManagement/actionRules/WeeklySuppression"),
		// 							Location: to.Ptr("Global"),
		// 							Tags: map[string]*string{
		// 							},
		// 							Properties: &armalertsmanagement.AlertProcessingRuleProperties{
		// 								Description: to.Ptr("Alert processing rule on resource group for adding action group"),
		// 								Actions: []armalertsmanagement.ActionClassification{
		// 									&armalertsmanagement.AddActionGroups{
		// 										ActionType: to.Ptr(armalertsmanagement.ActionTypeAddActionGroups),
		// 										ActionGroupIDs: []*string{
		// 											to.Ptr("actiongGroup1")},
		// 									}},
		// 									Conditions: []*armalertsmanagement.Condition{
		// 										{
		// 											Field: to.Ptr(armalertsmanagement.FieldMonitorCondition),
		// 											Operator: to.Ptr(armalertsmanagement.OperatorEquals),
		// 											Values: []*string{
		// 												to.Ptr("Fired")},
		// 											},
		// 											{
		// 												Field: to.Ptr(armalertsmanagement.FieldDescription),
		// 												Operator: to.Ptr(armalertsmanagement.OperatorContains),
		// 												Values: []*string{
		// 													to.Ptr("Percentage CPU greater than 80%"),
		// 													to.Ptr("Metric alert on resource foo")},
		// 											}},
		// 											Enabled: to.Ptr(true),
		// 											Scopes: []*string{
		// 												to.Ptr("/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/resourceGroups/alertscorrelationrg/providers/Microsoft.Compute/VirtualMachines/testResource")},
		// 											},
		// 											SystemData: &armalertsmanagement.SystemData{
		// 												CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09.000Z"); return t}()),
		// 												CreatedBy: to.Ptr("abc@microsoft.com"),
		// 												CreatedByType: to.Ptr(armalertsmanagement.CreatedByTypeUser),
		// 												LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09.000Z"); return t}()),
		// 												LastModifiedBy: to.Ptr("xyz@microsoft.com"),
		// 												LastModifiedByType: to.Ptr(armalertsmanagement.CreatedByTypeUser),
		// 											},
		// 									}},
		// 								}
	}
}
