package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listConnectionsByAutomationAccount_First100.json
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
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:29.29+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:33.617+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection0"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection0"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:36.493+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:36.493+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection1"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection1"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:36.76+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:36.76+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection10"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection10"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.117+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.117+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection100"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection100"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:04.527+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:04.527+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection101"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection101"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:04.777+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:04.777+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection102"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection102"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.027+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.027+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection103"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection103"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.277+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.277+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection104"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection104"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.527+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.527+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection105"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection105"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.82+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:05.82+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection106"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection106"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.07+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.07+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection107"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection107"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.37+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.37+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection108"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection108"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.62+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.62+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection109"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection109"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.87+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:06.87+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection11"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection11"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.367+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.367+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection110"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection110"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.167+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.167+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection111"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection111"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.463+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.463+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection112"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection112"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.73+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.73+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection113"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection113"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.963+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:07.963+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection114"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection114"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:08.213+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:08.213+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection115"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection115"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:08.527+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:08.527+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection116"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection116"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:08.76+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:08.76+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection117"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection117"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.103+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.103+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection118"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection118"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.37+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.37+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection119"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection119"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.65+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.65+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection12"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection12"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.617+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.617+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection120"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection120"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.883+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:09.883+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection121"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection121"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.167+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.167+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection122"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection122"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.43+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.43+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection123"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection123"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.68+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.68+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection124"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection124"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.93+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:10.93+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection125"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection125"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.18+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.18+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection126"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection126"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.43+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.43+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection127"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection127"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.667+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.667+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection128"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection128"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.917+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:11.917+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection129"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection129"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.18+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.18+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection13"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection13"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.883+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:39.883+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection130"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection130"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.43+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.43+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection131"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection131"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.667+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.667+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection132"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection132"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.917+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:12.917+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection133"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection133"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:13.213+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:13.213+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection134"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection134"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:13.463+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:13.463+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection135"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection135"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:13.73+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:13.73+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection136"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection136"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.01+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.01+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection137"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection137"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.29+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.29+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection138"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection138"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.54+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.54+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection139"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection139"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.807+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:14.807+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection14"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection14"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.133+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.133+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection140"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection140"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.057+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.057+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection141"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection141"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.307+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.307+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection142"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection142"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.603+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.603+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection143"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection143"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.853+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:15.853+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection144"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection144"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.12+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.12+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection145"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection145"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.37+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.37+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection146"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection146"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.633+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.633+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection147"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection147"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.9+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:16.9+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection148"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection148"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.167+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.167+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection149"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection149"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.43+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.43+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection15"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection15"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.37+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.37+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection150"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection150"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.68+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.68+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection151"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection151"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.963+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:17.963+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection152"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection152"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:18.527+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:18.527+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection153"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection153"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:18.79+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:18.79+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection154"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection154"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.073+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.073+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection155"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection155"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.353+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.353+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection156"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection156"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.68+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.68+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection157"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection157"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.93+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:19.93+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection158"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection158"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:20.243+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:20.243+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection159"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection159"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:20.557+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:20.557+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection16"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection16"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.62+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.62+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection160"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection160"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:20.823+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:20.823+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection161"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection161"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.15+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.15+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection162"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection162"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.417+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.417+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection163"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection163"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.65+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.65+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection164"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection164"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.993+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:21.993+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection165"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection165"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:22.323+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:22.323+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection166"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection166"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:22.68+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:22.68+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection167"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection167"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.057+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.057+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection168"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection168"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.417+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.417+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection169"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection169"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.697+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.697+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection17"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection17"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.87+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:40.87+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection170"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection170"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.963+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:23.963+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection171"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection171"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:24.243+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:24.243+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection172"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection172"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:24.51+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:24.51+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection173"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection173"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:24.777+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:24.777+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection174"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection174"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:25.057+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:25.057+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection175"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection175"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:25.34+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:53:25.34+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection18"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection18"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:41.197+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:41.197+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection19"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection19"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:41.57+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:41.57+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection2"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection2"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:37.01+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:37.01+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection20"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection20"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:41.913+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:41.913+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection21"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection21"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:42.273+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:42.273+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection22"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection22"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:42.663+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:42.663+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection23"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection23"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:42.993+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:42.993+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection24"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection24"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:43.26+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:43.26+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection25"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection25"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:43.587+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:43.587+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection26"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection26"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:43.87+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:43.87+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection27"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection27"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:44.227+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:44.227+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection28"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection28"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:44.557+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:44.557+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myConnection29"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/myConnection29"),
		// 			Properties: &armautomation.ConnectionProperties{
		// 				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
		// 					Name: to.Ptr("Azure"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:44.807+00:00"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:44.807+00:00"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
