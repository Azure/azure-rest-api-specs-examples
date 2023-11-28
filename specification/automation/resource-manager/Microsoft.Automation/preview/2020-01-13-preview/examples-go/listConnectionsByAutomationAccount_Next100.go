package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listConnectionsByAutomationAccount_Next100.json
func ExampleConnectionClient_NewListByAutomationAccountPager_listConnectionsByAutomationAccountNext100() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectionClient().NewListByAutomationAccountPager("rg", "myAutomationAccount28", nil)
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
		// page.ConnectionListResult = armautomation.ConnectionListResult{
		// 	Value: []*armautomation.Connection{
		// 		{
		// 			Name: to.Ptr("myConnection3"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection3"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:37.290Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:37.290Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection30"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection30"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:45.040Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:45.040Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection31"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection31"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:45.307Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:45.307Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection32"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection32"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:45.557Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:45.557Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection33"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection33"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:45.853Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:45.853Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection34"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection34"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:46.087Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:46.087Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection35"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection35"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:46.353Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:46.353Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection36"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection36"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:46.603Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:46.603Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection37"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection37"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:46.853Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:46.853Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection38"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection38"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:47.103Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:47.103Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection39"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection39"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:47.353Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:47.353Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection4"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection4"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:37.540Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:37.540Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection40"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection40"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:47.620Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:47.620Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection41"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection41"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:47.883Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:47.883Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection42"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection42"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:48.167Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:48.167Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection43"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection43"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:48.430Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:48.430Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection44"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection44"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:48.820Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:48.820Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection45"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection45"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:49.180Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:49.180Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection46"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection46"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:49.477Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:49.477Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection47"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection47"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:49.773Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:49.773Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection48"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection48"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:50.197Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:50.197Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection49"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection49"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:50.603Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:50.603Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection5"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection5"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:37.790Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:37.790Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection50"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection50"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:50.837Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:50.837Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection51"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection51"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:51.087Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:51.087Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection52"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection52"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:51.337Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:51.337Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection53"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection53"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:51.587Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:51.587Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection54"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection54"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:51.870Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:51.870Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection55"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection55"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:52.150Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:52.150Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection56"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection56"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:52.400Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:52.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection57"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection57"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:52.667Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:52.667Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection58"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection58"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:52.917Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:52.917Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection59"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection59"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:53.167Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:53.167Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection6"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection6"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:38.103Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:38.103Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection60"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection60"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:53.463Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:53.463Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection61"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection61"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:53.727Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:53.727Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection62"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection62"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:54.010Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:54.010Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection63"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection63"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:54.273Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:54.273Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection64"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection64"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:54.523Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:54.523Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection65"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection65"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:54.807Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:54.807Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection66"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection66"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:55.103Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:55.103Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection67"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection67"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:55.370Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:55.370Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection68"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection68"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:55.620Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:55.620Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection69"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection69"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:55.883Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:55.883Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection7"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection7"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:38.353Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:38.353Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection70"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection70"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:56.133Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:56.133Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection71"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection71"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:56.383Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:56.383Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection72"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection72"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:56.633Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:56.633Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection73"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection73"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:56.870Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:56.870Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection74"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection74"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:57.120Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:57.120Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection75"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection75"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:57.370Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:57.370Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection76"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection76"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:57.620Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:57.620Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection77"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection77"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:57.870Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:57.870Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection78"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection78"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:58.120Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:58.120Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection79"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection79"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:58.370Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:58.370Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection8"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection8"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:38.587Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:38.587Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection80"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection80"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:58.633Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:58.633Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection81"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection81"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:58.917Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:58.917Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection82"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection82"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:59.180Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:59.180Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection83"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection83"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:59.477Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:59.477Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection84"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection84"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:59.713Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:59.713Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection85"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection85"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:00.010Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:00.010Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection86"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection86"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:00.260Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:00.260Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection87"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection87"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:00.620Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:00.620Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection88"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection88"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:00.980Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:00.980Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection89"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection89"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:01.353Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:01.353Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection9"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection9"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:38.867Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:38.867Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection90"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection90"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:01.697Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:01.697Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection91"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection91"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:01.963Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:01.963Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection92"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection92"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:02.290Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:02.290Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection93"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection93"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:02.587Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:02.587Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection94"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection94"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:02.870Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:02.870Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection95"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection95"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:03.260Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:03.260Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection96"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection96"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:03.510Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:03.510Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection97"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection97"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:03.760Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:03.760Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection98"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection98"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:04.027Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:04.027Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection99"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection99"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:04.260Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:04.260Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
