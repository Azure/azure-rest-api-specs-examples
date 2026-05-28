package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2023-01-01-preview/ActivityLogAlertRule_Update.json
func ExampleActivityLogAlertsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewActivityLogAlertsClient().Update(ctx, "MyResourceGroup", "SampleActivityLogAlertRule", armmonitor.AlertRulePatchObject{
		Properties: &armmonitor.AlertRulePatchProperties{
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
	// res = armmonitor.ActivityLogAlertsClientUpdateResponse{
	// 	ActivityLogAlertResource: armmonitor.ActivityLogAlertResource{
	// 		Name: to.Ptr("SampleActivityLogAlertRule"),
	// 		Type: to.Ptr("Microsoft.Insights/ActivityLogAlerts"),
	// 		ID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/activityLogAlerts/SampleActivityLogAlertRule"),
	// 		Location: to.Ptr("Global"),
	// 		Properties: &armmonitor.AlertRuleProperties{
	// 			Description: to.Ptr("Description of sample Activity Log Alert rule."),
	// 			Actions: &armmonitor.ActionList{
	// 				ActionGroups: []*armmonitor.ActivityLogAlertActionGroup{
	// 					{
	// 						ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/actionGroups/SampleActionGroup"),
	// 						ActionProperties: map[string]*string{
	// 							"Email.Title": to.Ptr("my email title"),
	// 						},
	// 						WebhookProperties: map[string]*string{
	// 							"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Condition: &armmonitor.AlertRuleAllOfCondition{
	// 				AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
	// 					{
	// 						Equals: to.Ptr("Administrative"),
	// 						Field: to.Ptr("category"),
	// 					},
	// 					{
	// 						Equals: to.Ptr("Error"),
	// 						Field: to.Ptr("level"),
	// 					},
	// 				},
	// 			},
	// 			Enabled: to.Ptr(false),
	// 			Scopes: []*string{
	// 				to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d"),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
