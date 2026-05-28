package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2023-01-01-preview/ActivityLogAlertRule_ListByResourceGroupName.json
func ExampleActivityLogAlertsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewActivityLogAlertsClient().NewListByResourceGroupPager("MyResourceGroup", nil)
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
		// page = armmonitor.ActivityLogAlertsClientListByResourceGroupResponse{
		// 	AlertRuleList: armmonitor.AlertRuleList{
		// 		Value: []*armmonitor.ActivityLogAlertResource{
		// 			{
		// 				Name: to.Ptr("SampleActivityLogAlertRule1"),
		// 				Type: to.Ptr("Microsoft.Insights/ActivityLogAlerts"),
		// 				ID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/activityLogAlerts/SampleActivityLogAlertRule1"),
		// 				Location: to.Ptr("Global"),
		// 				Properties: &armmonitor.AlertRuleProperties{
		// 					Description: to.Ptr("Description of sample Activity Log Alert rule."),
		// 					Actions: &armmonitor.ActionList{
		// 						ActionGroups: []*armmonitor.ActivityLogAlertActionGroup{
		// 							{
		// 								ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/actionGroups/SampleActionGroup"),
		// 								ActionProperties: map[string]*string{
		// 									"Email.Title": to.Ptr("my email title"),
		// 								},
		// 								WebhookProperties: map[string]*string{
		// 									"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					Condition: &armmonitor.AlertRuleAllOfCondition{
		// 						AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
		// 							{
		// 								Equals: to.Ptr("Administrative"),
		// 								Field: to.Ptr("category"),
		// 							},
		// 							{
		// 								Equals: to.Ptr("Error"),
		// 								Field: to.Ptr("level"),
		// 							},
		// 						},
		// 					},
		// 					Enabled: to.Ptr(true),
		// 					Scopes: []*string{
		// 						to.Ptr("subscriptions/187f412d-1758-44d9-b052-169e2564721d"),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("SampleActivityLogAlertRule2"),
		// 				Type: to.Ptr("Microsoft.Insights/ActivityLogAlerts"),
		// 				ID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/activityLogAlerts/SampleActivityLogAlertRule2"),
		// 				Location: to.Ptr("Global"),
		// 				Properties: &armmonitor.AlertRuleProperties{
		// 					Description: to.Ptr("Description of sample Activity Log Alert rule."),
		// 					Actions: &armmonitor.ActionList{
		// 						ActionGroups: []*armmonitor.ActivityLogAlertActionGroup{
		// 							{
		// 								ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/actionGroups/SampleActionGroup"),
		// 								ActionProperties: map[string]*string{
		// 									"Email.Title": to.Ptr("my email title"),
		// 								},
		// 								WebhookProperties: map[string]*string{
		// 								},
		// 							},
		// 						},
		// 					},
		// 					Condition: &armmonitor.AlertRuleAllOfCondition{
		// 						AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
		// 							{
		// 								Equals: to.Ptr("Administrative"),
		// 								Field: to.Ptr("category"),
		// 							},
		// 							{
		// 								Equals: to.Ptr("Succeeded"),
		// 								Field: to.Ptr("status"),
		// 							},
		// 						},
		// 					},
		// 					Enabled: to.Ptr(true),
		// 					Scopes: []*string{
		// 						to.Ptr("subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup"),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
