package armtenantactivitylogalerts_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/tenantactivitylogalerts/armtenantactivitylogalerts"
)

// Generated from example definition: 2023-04-01-preview/TenantActivityLogAlertRule_CreateOrUpdate.json
func ExampleClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtenantactivitylogalerts.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().CreateOrUpdate(ctx, "72f988bf-86f1-41af-91ab-2d7cd011db47", "SampleActivityLogAlertSHRuleOnTenantLevel", armtenantactivitylogalerts.TenantActivityLogAlertResource{
		Location: to.Ptr("Global"),
		Properties: &armtenantactivitylogalerts.AlertRuleProperties{
			Description: to.Ptr("Description of sample Activity Log Alert service health rule on tenant level events."),
			Actions: &armtenantactivitylogalerts.ActionList{
				ActionGroups: []*armtenantactivitylogalerts.ActionGroup{
					{
						ActionGroupID: to.Ptr("/providers/Microsoft.Management/ManagementGroups/72f988bf-86f1-41af-91ab-2d7cd011db47/providers/Microsoft.Insights/actionGroups/SampleActionGroup"),
						ActionProperties: map[string]*string{
							"Email.Title": to.Ptr("my email title"),
						},
						WebhookProperties: map[string]*string{
							"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
						},
					},
				},
			},
			Condition: &armtenantactivitylogalerts.AlertRuleAllOfCondition{
				AllOf: []*armtenantactivitylogalerts.AlertRuleAnyOfOrLeafCondition{
					{
						Equals: to.Ptr("ServiceHealth"),
						Field:  to.Ptr("category"),
					},
				},
			},
			Enabled:     to.Ptr(true),
			TenantScope: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armtenantactivitylogalerts.ClientCreateOrUpdateResponse{
	// 	TenantActivityLogAlertResource: &armtenantactivitylogalerts.TenantActivityLogAlertResource{
	// 		Name: to.Ptr("SampleActivityLogAlertSHRuleOnTenantLevel"),
	// 		Type: to.Ptr("Microsoft.AlertsManagement/TenantActivityLogAlerts"),
	// 		ID: to.Ptr("/providers/Microsoft.Management/ManagementGroups/72f988bf-86f1-41af-91ab-2d7cd011db47/providers/Microsoft.AlertsManagement/TenantActivityLogAlerts/SampleActivityLogAlertSHRuleOnTenantLevel"),
	// 		Location: to.Ptr("Global"),
	// 		Properties: &armtenantactivitylogalerts.AlertRuleProperties{
	// 			Description: to.Ptr("Description of sample Activity Log Alert service health rule on tenant level events."),
	// 			Actions: &armtenantactivitylogalerts.ActionList{
	// 				ActionGroups: []*armtenantactivitylogalerts.ActionGroup{
	// 					{
	// 						ActionGroupID: to.Ptr("/providers/Microsoft.Management/ManagementGroups/72f988bf-86f1-41af-91ab-2d7cd011db47/providers/Microsoft.Insights/actionGroups/SampleActionGroup"),
	// 						ActionProperties: map[string]*string{
	// 							"Email.Title": to.Ptr("my email title"),
	// 						},
	// 						WebhookProperties: map[string]*string{
	// 							"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Condition: &armtenantactivitylogalerts.AlertRuleAllOfCondition{
	// 				AllOf: []*armtenantactivitylogalerts.AlertRuleAnyOfOrLeafCondition{
	// 					{
	// 						Equals: to.Ptr("ServiceHealth"),
	// 						Field: to.Ptr("category"),
	// 					},
	// 				},
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			TenantScope: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
