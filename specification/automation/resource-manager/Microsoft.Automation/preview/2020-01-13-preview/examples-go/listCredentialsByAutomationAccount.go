package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listCredentialsByAutomationAccount.json
func ExampleCredentialClient_NewListByAutomationAccountPager_listCredentialsByAutomationAccountFirst100() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCredentialClient().NewListByAutomationAccountPager("rg", "myAutomationAccount20", nil)
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
		// page.CredentialListResult = armautomation.CredentialListResult{
		// 	Value: []*armautomation.Credential{
		// 		{
		// 			Name: to.Ptr("myCredential"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:38.910Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:45.393Z"); return t}()),
		// 				UserName: to.Ptr("mylingaiah"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential0"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential0"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:51.770Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:51.770Z"); return t}()),
		// 				UserName: to.Ptr("myCredential0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential1"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential1"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:52.113Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:52.113Z"); return t}()),
		// 				UserName: to.Ptr("myCredential1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential10"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential10"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:55.300Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:55.300Z"); return t}()),
		// 				UserName: to.Ptr("myCredential10"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential100"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential100"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:25.410Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:25.410Z"); return t}()),
		// 				UserName: to.Ptr("myCredential100"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential101"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential101"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:25.690Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:25.690Z"); return t}()),
		// 				UserName: to.Ptr("myCredential101"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential102"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential102"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:26.003Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:26.003Z"); return t}()),
		// 				UserName: to.Ptr("myCredential102"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential103"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential103"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:26.337Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:26.337Z"); return t}()),
		// 				UserName: to.Ptr("myCredential103"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential104"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential104"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:26.677Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:26.677Z"); return t}()),
		// 				UserName: to.Ptr("myCredential104"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential105"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential105"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:27.020Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:27.020Z"); return t}()),
		// 				UserName: to.Ptr("myCredential105"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential106"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential106"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:27.350Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:27.350Z"); return t}()),
		// 				UserName: to.Ptr("myCredential106"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential107"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential107"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:27.660Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:27.660Z"); return t}()),
		// 				UserName: to.Ptr("myCredential107"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential108"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential108"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:27.990Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:27.990Z"); return t}()),
		// 				UserName: to.Ptr("myCredential108"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential109"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential109"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:28.300Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:28.300Z"); return t}()),
		// 				UserName: to.Ptr("myCredential109"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential11"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential11"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:55.677Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:55.677Z"); return t}()),
		// 				UserName: to.Ptr("myCredential11"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential110"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential110"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:28.630Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:28.630Z"); return t}()),
		// 				UserName: to.Ptr("myCredential110"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential111"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential111"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:28.943Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:28.943Z"); return t}()),
		// 				UserName: to.Ptr("myCredential111"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential112"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential112"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:29.253Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:29.253Z"); return t}()),
		// 				UserName: to.Ptr("myCredential112"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential113"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential113"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:29.567Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:29.567Z"); return t}()),
		// 				UserName: to.Ptr("myCredential113"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential114"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential114"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:29.880Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:29.880Z"); return t}()),
		// 				UserName: to.Ptr("myCredential114"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential115"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential115"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:30.207Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:30.207Z"); return t}()),
		// 				UserName: to.Ptr("myCredential115"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential116"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential116"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:30.520Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:30.520Z"); return t}()),
		// 				UserName: to.Ptr("myCredential116"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential117"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential117"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:30.833Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:30.833Z"); return t}()),
		// 				UserName: to.Ptr("myCredential117"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential118"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential118"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:31.240Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:31.240Z"); return t}()),
		// 				UserName: to.Ptr("myCredential118"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential119"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential119"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:31.600Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:31.600Z"); return t}()),
		// 				UserName: to.Ptr("myCredential119"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential12"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential12"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:56.020Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:56.020Z"); return t}()),
		// 				UserName: to.Ptr("myCredential12"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential120"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential120"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:31.943Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:31.943Z"); return t}()),
		// 				UserName: to.Ptr("myCredential120"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential121"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential121"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:32.333Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:32.333Z"); return t}()),
		// 				UserName: to.Ptr("myCredential121"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential122"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential122"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:32.660Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:32.660Z"); return t}()),
		// 				UserName: to.Ptr("myCredential122"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential123"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential123"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:32.957Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:32.957Z"); return t}()),
		// 				UserName: to.Ptr("myCredential123"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential124"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential124"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:33.270Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:33.270Z"); return t}()),
		// 				UserName: to.Ptr("myCredential124"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential125"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential125"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:33.600Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:33.600Z"); return t}()),
		// 				UserName: to.Ptr("myCredential125"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential126"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential126"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:33.957Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:33.957Z"); return t}()),
		// 				UserName: to.Ptr("myCredential126"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential127"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential127"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:34.287Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:34.287Z"); return t}()),
		// 				UserName: to.Ptr("myCredential127"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential128"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential128"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:34.613Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:34.613Z"); return t}()),
		// 				UserName: to.Ptr("myCredential128"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential129"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential129"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:34.973Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:34.973Z"); return t}()),
		// 				UserName: to.Ptr("myCredential129"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential13"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential13"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:56.393Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:56.393Z"); return t}()),
		// 				UserName: to.Ptr("myCredential13"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential130"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential130"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:35.363Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:35.363Z"); return t}()),
		// 				UserName: to.Ptr("myCredential130"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential131"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential131"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:35.707Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:35.707Z"); return t}()),
		// 				UserName: to.Ptr("myCredential131"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential132"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential132"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:36.037Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:36.037Z"); return t}()),
		// 				UserName: to.Ptr("myCredential132"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential133"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential133"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:36.380Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:36.380Z"); return t}()),
		// 				UserName: to.Ptr("myCredential133"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential134"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential134"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:36.740Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:36.740Z"); return t}()),
		// 				UserName: to.Ptr("myCredential134"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential135"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential135"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:37.050Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:37.050Z"); return t}()),
		// 				UserName: to.Ptr("myCredential135"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential136"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential136"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:37.410Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:37.410Z"); return t}()),
		// 				UserName: to.Ptr("myCredential136"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential137"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential137"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:37.723Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:37.723Z"); return t}()),
		// 				UserName: to.Ptr("myCredential137"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential138"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential138"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:38.037Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:38.037Z"); return t}()),
		// 				UserName: to.Ptr("myCredential138"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential139"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential139"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:38.350Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:38.350Z"); return t}()),
		// 				UserName: to.Ptr("myCredential139"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential14"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential14"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:56.723Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:56.723Z"); return t}()),
		// 				UserName: to.Ptr("myCredential14"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential140"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential140"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:38.710Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:38.710Z"); return t}()),
		// 				UserName: to.Ptr("myCredential140"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential141"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential141"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:39.037Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:39.037Z"); return t}()),
		// 				UserName: to.Ptr("myCredential141"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential142"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential142"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:39.350Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:39.350Z"); return t}()),
		// 				UserName: to.Ptr("myCredential142"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential143"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential143"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:39.677Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:39.677Z"); return t}()),
		// 				UserName: to.Ptr("myCredential143"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential144"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential144"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:39.990Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:39.990Z"); return t}()),
		// 				UserName: to.Ptr("myCredential144"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential145"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential145"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:40.317Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:40.317Z"); return t}()),
		// 				UserName: to.Ptr("myCredential145"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential146"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential146"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:40.630Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:40.630Z"); return t}()),
		// 				UserName: to.Ptr("myCredential146"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential147"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential147"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:40.943Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:40.943Z"); return t}()),
		// 				UserName: to.Ptr("myCredential147"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential148"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential148"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:41.270Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:41.270Z"); return t}()),
		// 				UserName: to.Ptr("myCredential148"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential149"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential149"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:41.600Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:41.600Z"); return t}()),
		// 				UserName: to.Ptr("myCredential149"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential15"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential15"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:57.033Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:57.033Z"); return t}()),
		// 				UserName: to.Ptr("myCredential15"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential150"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential150"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:41.910Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:41.910Z"); return t}()),
		// 				UserName: to.Ptr("myCredential150"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential151"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential151"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:42.240Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:42.240Z"); return t}()),
		// 				UserName: to.Ptr("myCredential151"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential152"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential152"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:42.567Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:42.567Z"); return t}()),
		// 				UserName: to.Ptr("myCredential152"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential153"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential153"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:42.880Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:42.880Z"); return t}()),
		// 				UserName: to.Ptr("myCredential153"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential154"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential154"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:43.193Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:43.193Z"); return t}()),
		// 				UserName: to.Ptr("myCredential154"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential155"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential155"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:43.520Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:43.520Z"); return t}()),
		// 				UserName: to.Ptr("myCredential155"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential156"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential156"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:43.833Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:43.833Z"); return t}()),
		// 				UserName: to.Ptr("myCredential156"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential157"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential157"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:44.147Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:44.147Z"); return t}()),
		// 				UserName: to.Ptr("myCredential157"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential158"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential158"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:44.473Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:44.473Z"); return t}()),
		// 				UserName: to.Ptr("myCredential158"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential159"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential159"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:44.787Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:44.787Z"); return t}()),
		// 				UserName: to.Ptr("myCredential159"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential16"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential16"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:57.363Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:57.363Z"); return t}()),
		// 				UserName: to.Ptr("myCredential16"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential160"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential160"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:45.100Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:45.100Z"); return t}()),
		// 				UserName: to.Ptr("myCredential160"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential161"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential161"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:45.427Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:45.427Z"); return t}()),
		// 				UserName: to.Ptr("myCredential161"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential162"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential162"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:45.740Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:45.740Z"); return t}()),
		// 				UserName: to.Ptr("myCredential162"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential163"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential163"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:46.067Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:46.067Z"); return t}()),
		// 				UserName: to.Ptr("myCredential163"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential164"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential164"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:46.380Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:46.380Z"); return t}()),
		// 				UserName: to.Ptr("myCredential164"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential165"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential165"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:46.710Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:46.710Z"); return t}()),
		// 				UserName: to.Ptr("myCredential165"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential166"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential166"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:47.037Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:47.037Z"); return t}()),
		// 				UserName: to.Ptr("myCredential166"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential167"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential167"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:47.350Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:47.350Z"); return t}()),
		// 				UserName: to.Ptr("myCredential167"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential168"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential168"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:47.677Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:47.677Z"); return t}()),
		// 				UserName: to.Ptr("myCredential168"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential169"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential169"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:47.990Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:47.990Z"); return t}()),
		// 				UserName: to.Ptr("myCredential169"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential17"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential17"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:57.677Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:57.677Z"); return t}()),
		// 				UserName: to.Ptr("myCredential17"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential170"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential170"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:48.287Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:48.287Z"); return t}()),
		// 				UserName: to.Ptr("myCredential170"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential171"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential171"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:48.613Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:48.613Z"); return t}()),
		// 				UserName: to.Ptr("myCredential171"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential172"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential172"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:48.927Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:48.927Z"); return t}()),
		// 				UserName: to.Ptr("myCredential172"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential173"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential173"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:49.257Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:49.257Z"); return t}()),
		// 				UserName: to.Ptr("myCredential173"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential174"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential174"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:49.567Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:49.567Z"); return t}()),
		// 				UserName: to.Ptr("myCredential174"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential175"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential175"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:49.880Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:49.880Z"); return t}()),
		// 				UserName: to.Ptr("myCredential175"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential176"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential176"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:50.210Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:50.210Z"); return t}()),
		// 				UserName: to.Ptr("myCredential176"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential177"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential177"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:50.537Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:50.537Z"); return t}()),
		// 				UserName: to.Ptr("myCredential177"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential178"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential178"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:50.863Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:50.863Z"); return t}()),
		// 				UserName: to.Ptr("myCredential178"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential179"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential179"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:51.193Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:51.193Z"); return t}()),
		// 				UserName: to.Ptr("myCredential179"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential18"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential18"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:58.003Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:58.003Z"); return t}()),
		// 				UserName: to.Ptr("myCredential18"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential180"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential180"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:51.600Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:51.600Z"); return t}()),
		// 				UserName: to.Ptr("myCredential180"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential181"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential181"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:51.910Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:51.910Z"); return t}()),
		// 				UserName: to.Ptr("myCredential181"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential182"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential182"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:52.223Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:52.223Z"); return t}()),
		// 				UserName: to.Ptr("myCredential182"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential183"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential183"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:52.600Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:52.600Z"); return t}()),
		// 				UserName: to.Ptr("myCredential183"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential184"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential184"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:52.927Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:52.927Z"); return t}()),
		// 				UserName: to.Ptr("myCredential184"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential185"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential185"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:53.270Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:53.270Z"); return t}()),
		// 				UserName: to.Ptr("myCredential185"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential186"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential186"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:53.583Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:53.583Z"); return t}()),
		// 				UserName: to.Ptr("myCredential186"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential187"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential187"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:54.037Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:54.037Z"); return t}()),
		// 				UserName: to.Ptr("myCredential187"),
		// 			},
		// 	}},
		// }
	}
}
