package armtenantactivitylogalerts_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/tenantactivitylogalerts/armtenantactivitylogalerts"
)

// Generated from example definition: 2023-04-01-preview/TenantActivityLogAlertRule_ListByTenant.json
func ExampleClient_NewListByTenantPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtenantactivitylogalerts.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListByTenantPager(nil)
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
		// page = armtenantactivitylogalerts.ClientListByTenantResponse{
		// 	TenantAlertRuleList: armtenantactivitylogalerts.TenantAlertRuleList{
		// 		Value: []*armtenantactivitylogalerts.TenantActivityLogAlertResource{
		// 			{
		// 				ID: to.Ptr("/providers/Microsoft.Management/ManagementGroups/72f988bf-86f1-41af-91ab-2d7cd011db47/providers/Microsoft.AlertsManagement/TenantActivityLogAlerts/SampleActivityLogAlertSHRuleOnTenantLevel"),
		// 				Type: to.Ptr("Microsoft.AlertsManagement/TenantActivityLogAlerts"),
		// 				Name: to.Ptr("SampleActivityLogAlertSHRuleOnTenantLevel"),
		// 				Location: to.Ptr("Global"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armtenantactivitylogalerts.AlertRuleProperties{
		// 					TenantScope: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 					Condition: &armtenantactivitylogalerts.AlertRuleAllOfCondition{
		// 						AllOf: []*armtenantactivitylogalerts.AlertRuleAnyOfOrLeafCondition{
		// 							{
		// 								Field: to.Ptr("category"),
		// 								Equals: to.Ptr("ServiceHealth"),
		// 							},
		// 						},
		// 					},
		// 					Actions: &armtenantactivitylogalerts.ActionList{
		// 						ActionGroups: []*armtenantactivitylogalerts.ActionGroup{
		// 							{
		// 								ActionGroupID: to.Ptr("/providers/Microsoft.Management/ManagementGroups/72f988bf-86f1-41af-91ab-2d7cd011db47/providers/Microsoft.Insights/actionGroups/SampleActionGroup"),
		// 								WebhookProperties: map[string]*string{
		// 									"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
		// 								},
		// 								ActionProperties: map[string]*string{
		// 									"Email.Title": to.Ptr("my email title"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					Enabled: to.Ptr(true),
		// 					Description: to.Ptr("Description of sample Activity Log Alert service health rule on tenant level events."),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/providers/Microsoft.Management/ManagementGroups/MyManagementGroup/providers/Microsoft.AlertsManagement/TenantActivityLogAlerts/SampleActivityLogAlertSHRuleOnTenantLevel"),
		// 				Type: to.Ptr("Microsoft.AlertsManagement/TenantActivityLogAlerts"),
		// 				Name: to.Ptr("SampleActivityLogAlertSHRuleOnTenantLevel"),
		// 				Location: to.Ptr("Global"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armtenantactivitylogalerts.AlertRuleProperties{
		// 					TenantScope: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 					Condition: &armtenantactivitylogalerts.AlertRuleAllOfCondition{
		// 						AllOf: []*armtenantactivitylogalerts.AlertRuleAnyOfOrLeafCondition{
		// 							{
		// 								Field: to.Ptr("category"),
		// 								Equals: to.Ptr("ServiceHealth"),
		// 							},
		// 						},
		// 					},
		// 					Actions: &armtenantactivitylogalerts.ActionList{
		// 						ActionGroups: []*armtenantactivitylogalerts.ActionGroup{
		// 							{
		// 								ActionGroupID: to.Ptr("/providers/Microsoft.Management/ManagementGroups/MyManagementGroup/providers/Microsoft.Insights/actionGroups/SampleActionGroup"),
		// 								WebhookProperties: map[string]*string{
		// 									"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
		// 								},
		// 								ActionProperties: map[string]*string{
		// 									"Email.Title": to.Ptr("my email title"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					Enabled: to.Ptr(true),
		// 					Description: to.Ptr("Description of sample Activity Log Alert service health rule on tenant level events."),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
