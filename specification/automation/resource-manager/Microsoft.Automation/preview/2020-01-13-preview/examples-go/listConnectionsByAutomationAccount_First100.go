package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listConnectionsByAutomationAccount_First100.json
func ExampleConnectionClient_NewListByAutomationAccountPager_listConnectionsByAutomationAccountFirst100() {
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
		// 			Name: to.Ptr("myConnection"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:29.290Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:33.617Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection0"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection0"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:36.493Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:36.493Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection1"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection1"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:36.760Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:36.760Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection10"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection10"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.117Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.117Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection100"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection100"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:04.527Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:04.527Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection101"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection101"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:04.777Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:04.777Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection102"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection102"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.027Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.027Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection103"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection103"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.277Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.277Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection104"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection104"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.527Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.527Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection105"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection105"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.820Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.820Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection106"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection106"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.070Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.070Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection107"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection107"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.370Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.370Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection108"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection108"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.620Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.620Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection109"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection109"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.870Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.870Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection11"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection11"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.367Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.367Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection110"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection110"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.167Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.167Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection111"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection111"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.463Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.463Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection112"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection112"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.730Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.730Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection113"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection113"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.963Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.963Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection114"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection114"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:08.213Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:08.213Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection115"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection115"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:08.527Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:08.527Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection116"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection116"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:08.760Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:08.760Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection117"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection117"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.103Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.103Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection118"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection118"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.370Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.370Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection119"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection119"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.650Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.650Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection12"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection12"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.617Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.617Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection120"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection120"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.883Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.883Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection121"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection121"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.167Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.167Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection122"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection122"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.430Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.430Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection123"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection123"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.680Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.680Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection124"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection124"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.930Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.930Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection125"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection125"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.180Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.180Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection126"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection126"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.430Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.430Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection127"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection127"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.667Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.667Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection128"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection128"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.917Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.917Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection129"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection129"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.180Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.180Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection13"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection13"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.883Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.883Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection130"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection130"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.430Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.430Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection131"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection131"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.667Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.667Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection132"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection132"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.917Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.917Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection133"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection133"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:13.213Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:13.213Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection134"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection134"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:13.463Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:13.463Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection135"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection135"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:13.730Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:13.730Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection136"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection136"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.010Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.010Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection137"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection137"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.290Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.290Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection138"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection138"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.540Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.540Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection139"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection139"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.807Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.807Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection14"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection14"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.133Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.133Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection140"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection140"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.057Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.057Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection141"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection141"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.307Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.307Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection142"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection142"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.603Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.603Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection143"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection143"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.853Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.853Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection144"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection144"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.120Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.120Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection145"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection145"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.370Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.370Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection146"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection146"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.633Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.633Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection147"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection147"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.900Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.900Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection148"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection148"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.167Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.167Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection149"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection149"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.430Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.430Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection15"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection15"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.370Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.370Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection150"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection150"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.680Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.680Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection151"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection151"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.963Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.963Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection152"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection152"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:18.527Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:18.527Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection153"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection153"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:18.790Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:18.790Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection154"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection154"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.073Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.073Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection155"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection155"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.353Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.353Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection156"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection156"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.680Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.680Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection157"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection157"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.930Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.930Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection158"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection158"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:20.243Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:20.243Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection159"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection159"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:20.557Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:20.557Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection16"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection16"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.620Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.620Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection160"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection160"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:20.823Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:20.823Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection161"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection161"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.150Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.150Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection162"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection162"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.417Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.417Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection163"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection163"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.650Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.650Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection164"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection164"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.993Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.993Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection165"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection165"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:22.323Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:22.323Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection166"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection166"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:22.680Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:22.680Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection167"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection167"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.057Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.057Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection168"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection168"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.417Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.417Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection169"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection169"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.697Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.697Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection17"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection17"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.870Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.870Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection170"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection170"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.963Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.963Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection171"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection171"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:24.243Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:24.243Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection172"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection172"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:24.510Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:24.510Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection173"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection173"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:24.777Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:24.777Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection174"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection174"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:25.057Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:25.057Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection175"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection175"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:25.340Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:25.340Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection18"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection18"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:41.197Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:41.197Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection19"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection19"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:41.570Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:41.570Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection2"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection2"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:37.010Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:37.010Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection20"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection20"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:41.913Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:41.913Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection21"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection21"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:42.273Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:42.273Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection22"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection22"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:42.663Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:42.663Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection23"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection23"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:42.993Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:42.993Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection24"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection24"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:43.260Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:43.260Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection25"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection25"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:43.587Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:43.587Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection26"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection26"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:43.870Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:43.870Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection27"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection27"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:44.227Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:44.227Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection28"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection28"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:44.557Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:44.557Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection29"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection29"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:44.807Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:44.807Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
