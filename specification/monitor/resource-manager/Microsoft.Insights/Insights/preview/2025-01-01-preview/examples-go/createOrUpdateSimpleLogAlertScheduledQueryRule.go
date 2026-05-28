package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2025-01-01-preview/createOrUpdateSimpleLogAlertScheduledQueryRule.json
func ExampleScheduledQueryRulesClient_CreateOrUpdate_createOrUpdateASimpleLogAlertScheduledQueryRuleOnSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("dd4bfc94-a096-412b-9c43-4bd13e35afbc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledQueryRulesClient().CreateOrUpdate(ctx, "QueryResourceGroupName", "perf", armmonitor.ScheduledQueryRuleResource{
		Kind:     to.Ptr(armmonitor.KindSimpleLogAlert),
		Location: to.Ptr("eastus"),
		Properties: &armmonitor.ScheduledQueryRuleProperties{
			Description: to.Ptr("Performance rule"),
			Actions: &armmonitor.Actions{
				ActionGroups: []*string{
					to.Ptr("/subscriptions/1cf177ed-1330-4692-80ea-fd3d7783b147/resourcegroups/sqrapi/providers/microsoft.insights/actiongroups/myactiongroup"),
				},
				ActionProperties: map[string]*string{
					"Icm.Title": to.Ptr("Custom title in ICM"),
					"Icm.TsgId": to.Ptr("https://tsg.url"),
				},
				CustomProperties: map[string]*string{
					"key11": to.Ptr("value11"),
					"key12": to.Ptr("value12"),
				},
			},
			AutoMitigate:                          to.Ptr(false),
			CheckWorkspaceAlertsStorageConfigured: to.Ptr(true),
			Criteria: &armmonitor.ScheduledQueryRuleCriteria{
				AllOf: []*armmonitor.Condition{
					{
						Query: to.Ptr("Perf | where ObjectName == \"Processor\""),
					},
				},
			},
			Enabled: to.Ptr(true),
			Scopes: []*string{
				to.Ptr("/subscriptions/aaf177ed-1330-a9f2-80ea-fd3d7783b147/resourceGroups/scopeResourceGroup1/providers/Microsoft.Compute/virtualMachines/vm1"),
			},
			Severity:            to.Ptr(armmonitor.AlertSeverityFour),
			SkipQueryValidation: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.ScheduledQueryRulesClientCreateOrUpdateResponse{
	// 	ScheduledQueryRuleResource: armmonitor.ScheduledQueryRuleResource{
	// 		Name: to.Ptr("perf"),
	// 		Type: to.Ptr("microsoft.insights/scheduledqueryrules"),
	// 		ID: to.Ptr("/subscriptions/dd4bfc94-a096-412b-9c43-4bd13e35afbc/resourcegroups/QueryResourceGroupName/providers/microsoft.insights/scheduledqueryrules/perf"),
	// 		Kind: to.Ptr(armmonitor.KindSimpleLogAlert),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armmonitor.ScheduledQueryRuleProperties{
	// 			Description: to.Ptr("Performance rule"),
	// 			Actions: &armmonitor.Actions{
	// 				ActionGroups: []*string{
	// 					to.Ptr("/subscriptions/1cf177ed-1330-4692-80ea-fd3d7783b147/resourcegroups/sqrapi/providers/microsoft.insights/actiongroups/myactiongroup"),
	// 				},
	// 				ActionProperties: map[string]*string{
	// 					"Icm.Title": to.Ptr("Custom title in ICM"),
	// 					"Icm.TsgId": to.Ptr("https://tsg.url"),
	// 				},
	// 				CustomProperties: map[string]*string{
	// 					"key11": to.Ptr("value11"),
	// 					"key12": to.Ptr("value12"),
	// 				},
	// 			},
	// 			AutoMitigate: to.Ptr(false),
	// 			CheckWorkspaceAlertsStorageConfigured: to.Ptr(true),
	// 			Criteria: &armmonitor.ScheduledQueryRuleCriteria{
	// 				AllOf: []*armmonitor.Condition{
	// 					{
	// 						Query: to.Ptr("Perf | where ObjectName == \"Processor\""),
	// 					},
	// 				},
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			IsWorkspaceAlertsStorageConfigured: to.Ptr(true),
	// 			Scopes: []*string{
	// 				to.Ptr("/subscriptions/aaf177ed-1330-a9f2-80ea-fd3d7783b147/resourceGroups/scopeResourceGroup1/providers/Microsoft.Compute/virtualMachines/vm1"),
	// 			},
	// 			Severity: to.Ptr(		armmonitor.AlertSeverityFour),
	// 			SkipQueryValidation: to.Ptr(true),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
