package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: 2025-05-25-preview/Alerts_History.json
func ExampleAlertsClient_GetHistory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("subscriptions/9e261de7-c804-4b9d-9ebf-6f50fe350a9a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertsClient().GetHistory(ctx, "66114d64-d9d9-478b-95c9-b789d6502100", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armalertsmanagement.AlertsClientGetHistoryResponse{
	// 	AlertModification: &armalertsmanagement.AlertModification{
	// 		Name: to.Ptr("CPU Alert"),
	// 		Type: to.Ptr("Microsoft.AlertsManagement/alerts"),
	// 		ID: to.Ptr("/subscriptions/9e261de7-c804-4b9d-9ebf-6f50fe350a9a/resourceGroups/someResourceGroup/providers/Microsoft.AlertsManagement/alerts/66114d64-d9d9-478b-95c9-b789d6502100"),
	// 		Properties: &armalertsmanagement.AlertModificationProperties{
	// 			AlertID: to.Ptr("66114d64-d9d9-478b-95c9-b789d6502100"),
	// 			Modifications: []*armalertsmanagement.AlertModificationItem{
	// 				{
	// 					Description: to.Ptr("New Alert Object is created"),
	// 					Comments: to.Ptr(""),
	// 					ModificationEvent: to.Ptr(armalertsmanagement.AlertModificationEventAlertCreated),
	// 					ModifiedAt: to.Ptr("2018-06-13T06:09:01Z"),
	// 					ModifiedBy: to.Ptr("System"),
	// 					NewValue: to.Ptr(""),
	// 					OldValue: to.Ptr(""),
	// 					Details: nil,
	// 				},
	// 				{
	// 					Description: to.Ptr("State changed from 'New' to 'Acknowledged'"),
	// 					Comments: to.Ptr("The alert has been resolved"),
	// 					ModificationEvent: to.Ptr(armalertsmanagement.AlertModificationEventStateChange),
	// 					ModifiedAt: to.Ptr("2018-06-13T06:14:15.7378737Z"),
	// 					ModifiedBy: to.Ptr("vikramm@microsoft.com"),
	// 					NewValue: to.Ptr("Resolved"),
	// 					OldValue: to.Ptr("Fired"),
	// 					Details: &armalertsmanagement.PropertyChangeDetails{
	// 						Type: to.Ptr(armalertsmanagement.AlertModificationTypePropertyChange),
	// 						Comment: to.Ptr("The alert has been resolved"),
	// 						NewValue: to.Ptr("Resolved"),
	// 						OldValue: to.Ptr("Fired"),
	// 					},
	// 				},
	// 				{
	// 					Description: to.Ptr("New Alert Object is created"),
	// 					Comments: to.Ptr(""),
	// 					ModificationEvent: to.Ptr(armalertsmanagement.AlertModificationEventActionsTriggered),
	// 					ModifiedAt: to.Ptr("2018-06-13T06:09:01Z"),
	// 					ModifiedBy: to.Ptr("System"),
	// 					NewValue: to.Ptr(""),
	// 					OldValue: to.Ptr(""),
	// 					Details: &armalertsmanagement.ActionTriggeredDetails{
	// 						Type: to.Ptr(armalertsmanagement.AlertModificationTypeActionsTriggered),
	// 						ActionGroup: &armalertsmanagement.TriggeredRule{
	// 							ActionGroupID: to.Ptr("/subscriptions/9e261de7-c804-4b9d-9ebf-6f50fe350a9a/resourceGroups/rg1/providers/Microsoft.Insights/actionGroups/actionGroup1"),
	// 							RuleID: to.Ptr("RuleId1"),
	// 							RuleType: to.Ptr(armalertsmanagement.RuleTypeAlertRule),
	// 						},
	// 						NotificationResult: &armalertsmanagement.NotificationResult{
	// 							Status: to.Ptr(armalertsmanagement.ResultStatus("Accepted")),
	// 							StatusURL: to.Ptr("https://management.azure.com/subscriptions/2a784a95-81bd-41c8-ba8a-362d1098a2b9/resourceGroups/AzNSTest/providers/microsoft.insights/actionGroups/ag2/notificationStatus/00000000000?api-version=2021-09-01"),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Description: to.Ptr("New Alert Object is created"),
	// 					Comments: to.Ptr(""),
	// 					ModificationEvent: to.Ptr(armalertsmanagement.AlertModificationEventActionsSuppressed),
	// 					ModifiedAt: to.Ptr("2018-06-13T06:09:01Z"),
	// 					ModifiedBy: to.Ptr("System"),
	// 					NewValue: to.Ptr(""),
	// 					OldValue: to.Ptr(""),
	// 					Details: &armalertsmanagement.ActionSuppressedDetails{
	// 						Type: to.Ptr(armalertsmanagement.AlertModificationTypeActionsSuppressed),
	// 						SuppressedActionGroups: []*armalertsmanagement.TriggeredRule{
	// 							{
	// 								ActionGroupID: to.Ptr("/subscriptions/9e261de7-c804-4b9d-9ebf-6f50fe350a9a/resourceGroups/rg1/providers/Microsoft.Insights/actionGroups/actionGroup1"),
	// 								RuleID: to.Ptr("RuleId1"),
	// 								RuleType: to.Ptr(armalertsmanagement.RuleTypeAlertRule),
	// 							},
	// 						},
	// 						SuppressionActionRules: []*string{
	// 							to.Ptr("ActionRule 1"),
	// 							to.Ptr("ActionRule 2"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
