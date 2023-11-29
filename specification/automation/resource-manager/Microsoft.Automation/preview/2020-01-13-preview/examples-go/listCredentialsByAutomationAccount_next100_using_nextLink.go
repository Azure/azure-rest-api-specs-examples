package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listCredentialsByAutomationAccount_next100_using_nextLink.json
func ExampleCredentialClient_NewListByAutomationAccountPager_listCredentialsByAutomationAccountNext100() {
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
		// 			Name: to.Ptr("myCredential188"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential188"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:54.397Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:54.397Z"); return t}()),
		// 				UserName: to.Ptr("myCredential188"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential189"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential189"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:54.710Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:54.710Z"); return t}()),
		// 				UserName: to.Ptr("myCredential189"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential19"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential19"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:58.330Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:58.330Z"); return t}()),
		// 				UserName: to.Ptr("myCredential19"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential190"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential190"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:55.037Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:55.037Z"); return t}()),
		// 				UserName: to.Ptr("myCredential190"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential191"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential191"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:55.350Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:55.350Z"); return t}()),
		// 				UserName: to.Ptr("myCredential191"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential192"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential192"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:55.677Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:55.677Z"); return t}()),
		// 				UserName: to.Ptr("myCredential192"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential193"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential193"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:56.007Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:56.007Z"); return t}()),
		// 				UserName: to.Ptr("myCredential193"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential194"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential194"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:56.333Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:56.333Z"); return t}()),
		// 				UserName: to.Ptr("myCredential194"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential195"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential195"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:56.660Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:56.660Z"); return t}()),
		// 				UserName: to.Ptr("myCredential195"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential196"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential196"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:57.007Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:57.007Z"); return t}()),
		// 				UserName: to.Ptr("myCredential196"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential197"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential197"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:57.333Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:57.333Z"); return t}()),
		// 				UserName: to.Ptr("myCredential197"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential198"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential198"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:57.660Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:57.660Z"); return t}()),
		// 				UserName: to.Ptr("myCredential198"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential199"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential199"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:57.973Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:57.973Z"); return t}()),
		// 				UserName: to.Ptr("myCredential199"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential2"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential2"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:52.440Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:52.440Z"); return t}()),
		// 				UserName: to.Ptr("myCredential2"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential20"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential20"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:58.660Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:58.660Z"); return t}()),
		// 				UserName: to.Ptr("myCredential20"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential200"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential200"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:58.303Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:58.303Z"); return t}()),
		// 				UserName: to.Ptr("myCredential200"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential21"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential21"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:59.033Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:59.033Z"); return t}()),
		// 				UserName: to.Ptr("myCredential21"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential22"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential22"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:59.363Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:59.363Z"); return t}()),
		// 				UserName: to.Ptr("myCredential22"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential23"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential23"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:59.707Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:59.707Z"); return t}()),
		// 				UserName: to.Ptr("myCredential23"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential24"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential24"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:00.020Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:00.020Z"); return t}()),
		// 				UserName: to.Ptr("myCredential24"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential25"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential25"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:00.330Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:00.330Z"); return t}()),
		// 				UserName: to.Ptr("myCredential25"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential26"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential26"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:00.707Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:00.707Z"); return t}()),
		// 				UserName: to.Ptr("myCredential26"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential27"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential27"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:01.020Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:01.020Z"); return t}()),
		// 				UserName: to.Ptr("myCredential27"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential28"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential28"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:01.347Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:01.347Z"); return t}()),
		// 				UserName: to.Ptr("myCredential28"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential29"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential29"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:01.847Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:01.847Z"); return t}()),
		// 				UserName: to.Ptr("myCredential29"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential3"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential3"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:52.863Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:52.863Z"); return t}()),
		// 				UserName: to.Ptr("myCredential3"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential30"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential30"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:02.160Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:02.160Z"); return t}()),
		// 				UserName: to.Ptr("myCredential30"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential31"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential31"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:02.503Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:02.503Z"); return t}()),
		// 				UserName: to.Ptr("myCredential31"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential32"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential32"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:02.880Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:02.880Z"); return t}()),
		// 				UserName: to.Ptr("myCredential32"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential33"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential33"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:03.177Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:03.177Z"); return t}()),
		// 				UserName: to.Ptr("myCredential33"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential34"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential34"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:03.550Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:03.550Z"); return t}()),
		// 				UserName: to.Ptr("myCredential34"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential35"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential35"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:03.893Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:03.893Z"); return t}()),
		// 				UserName: to.Ptr("myCredential35"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential36"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential36"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:04.207Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:04.207Z"); return t}()),
		// 				UserName: to.Ptr("myCredential36"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential37"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential37"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:04.550Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:04.550Z"); return t}()),
		// 				UserName: to.Ptr("myCredential37"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential38"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential38"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:04.847Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:04.847Z"); return t}()),
		// 				UserName: to.Ptr("myCredential38"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential39"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential39"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:05.177Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:05.177Z"); return t}()),
		// 				UserName: to.Ptr("myCredential39"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential4"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential4"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:53.253Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:53.253Z"); return t}()),
		// 				UserName: to.Ptr("myCredential4"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential40"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential40"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:05.490Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:05.490Z"); return t}()),
		// 				UserName: to.Ptr("myCredential40"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential41"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential41"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:05.830Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:05.830Z"); return t}()),
		// 				UserName: to.Ptr("myCredential41"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential42"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential42"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:06.143Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:06.143Z"); return t}()),
		// 				UserName: to.Ptr("myCredential42"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential43"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential43"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:06.457Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:06.457Z"); return t}()),
		// 				UserName: to.Ptr("myCredential43"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential44"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential44"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:06.940Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:06.940Z"); return t}()),
		// 				UserName: to.Ptr("myCredential44"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential45"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential45"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:07.300Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:07.300Z"); return t}()),
		// 				UserName: to.Ptr("myCredential45"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential46"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential46"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:07.630Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:07.630Z"); return t}()),
		// 				UserName: to.Ptr("myCredential46"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential47"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential47"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:07.957Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:07.957Z"); return t}()),
		// 				UserName: to.Ptr("myCredential47"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential48"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential48"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:08.270Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:08.270Z"); return t}()),
		// 				UserName: to.Ptr("myCredential48"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential49"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential49"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:08.580Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:08.580Z"); return t}()),
		// 				UserName: to.Ptr("myCredential49"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential5"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential5"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:53.567Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:53.567Z"); return t}()),
		// 				UserName: to.Ptr("myCredential5"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential50"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential50"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:09.020Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:09.020Z"); return t}()),
		// 				UserName: to.Ptr("myCredential50"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential51"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential51"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:09.330Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:09.330Z"); return t}()),
		// 				UserName: to.Ptr("myCredential51"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential52"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential52"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:09.643Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:09.643Z"); return t}()),
		// 				UserName: to.Ptr("myCredential52"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential53"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential53"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:09.973Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:09.973Z"); return t}()),
		// 				UserName: to.Ptr("myCredential53"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential54"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential54"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:10.287Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:10.287Z"); return t}()),
		// 				UserName: to.Ptr("myCredential54"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential55"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential55"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:10.660Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:10.660Z"); return t}()),
		// 				UserName: to.Ptr("myCredential55"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential56"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential56"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:10.973Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:10.973Z"); return t}()),
		// 				UserName: to.Ptr("myCredential56"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential57"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential57"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:11.330Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:11.330Z"); return t}()),
		// 				UserName: to.Ptr("myCredential57"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential58"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential58"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:11.643Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:11.643Z"); return t}()),
		// 				UserName: to.Ptr("myCredential58"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential59"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential59"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:11.990Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:11.990Z"); return t}()),
		// 				UserName: to.Ptr("myCredential59"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential6"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential6"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:53.863Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:53.863Z"); return t}()),
		// 				UserName: to.Ptr("myCredential6"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential60"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential60"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:12.300Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:12.300Z"); return t}()),
		// 				UserName: to.Ptr("myCredential60"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential61"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential61"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:12.630Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:12.630Z"); return t}()),
		// 				UserName: to.Ptr("myCredential61"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential62"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential62"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:12.957Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:12.957Z"); return t}()),
		// 				UserName: to.Ptr("myCredential62"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential63"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential63"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:13.270Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:13.270Z"); return t}()),
		// 				UserName: to.Ptr("myCredential63"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential64"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential64"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:13.583Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:13.583Z"); return t}()),
		// 				UserName: to.Ptr("myCredential64"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential65"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential65"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:13.893Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:13.893Z"); return t}()),
		// 				UserName: to.Ptr("myCredential65"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential66"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential66"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:14.207Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:14.207Z"); return t}()),
		// 				UserName: to.Ptr("myCredential66"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential67"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential67"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:14.537Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:14.537Z"); return t}()),
		// 				UserName: to.Ptr("myCredential67"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential68"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential68"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:14.847Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:14.847Z"); return t}()),
		// 				UserName: to.Ptr("myCredential68"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential69"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential69"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:15.190Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:15.190Z"); return t}()),
		// 				UserName: to.Ptr("myCredential69"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential7"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential7"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:54.190Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:54.190Z"); return t}()),
		// 				UserName: to.Ptr("myCredential7"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential70"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential70"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:15.503Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:15.503Z"); return t}()),
		// 				UserName: to.Ptr("myCredential70"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential71"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential71"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:15.910Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:15.910Z"); return t}()),
		// 				UserName: to.Ptr("myCredential71"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential72"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential72"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:16.253Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:16.253Z"); return t}()),
		// 				UserName: to.Ptr("myCredential72"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential73"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential73"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:16.567Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:16.567Z"); return t}()),
		// 				UserName: to.Ptr("myCredential73"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential74"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential74"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:16.927Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:16.927Z"); return t}()),
		// 				UserName: to.Ptr("myCredential74"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential75"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential75"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:17.253Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:17.253Z"); return t}()),
		// 				UserName: to.Ptr("myCredential75"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential76"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential76"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:17.567Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:17.567Z"); return t}()),
		// 				UserName: to.Ptr("myCredential76"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential77"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential77"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:17.880Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:17.880Z"); return t}()),
		// 				UserName: to.Ptr("myCredential77"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential78"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential78"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:18.190Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:18.190Z"); return t}()),
		// 				UserName: to.Ptr("myCredential78"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential79"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential79"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:18.520Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:18.520Z"); return t}()),
		// 				UserName: to.Ptr("myCredential79"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential8"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential8"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:54.503Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:54.503Z"); return t}()),
		// 				UserName: to.Ptr("myCredential8"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential80"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential80"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:18.833Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:18.833Z"); return t}()),
		// 				UserName: to.Ptr("myCredential80"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential81"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential81"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:19.177Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:19.177Z"); return t}()),
		// 				UserName: to.Ptr("myCredential81"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential82"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential82"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:19.490Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:19.490Z"); return t}()),
		// 				UserName: to.Ptr("myCredential82"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential83"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential83"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:19.833Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:19.833Z"); return t}()),
		// 				UserName: to.Ptr("myCredential83"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential84"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential84"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:20.177Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:20.177Z"); return t}()),
		// 				UserName: to.Ptr("myCredential84"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential85"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential85"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:20.503Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:20.503Z"); return t}()),
		// 				UserName: to.Ptr("myCredential85"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential86"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential86"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:20.817Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:20.817Z"); return t}()),
		// 				UserName: to.Ptr("myCredential86"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential87"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential87"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:21.143Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:21.143Z"); return t}()),
		// 				UserName: to.Ptr("myCredential87"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential88"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential88"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:21.457Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:21.457Z"); return t}()),
		// 				UserName: to.Ptr("myCredential88"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential89"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential89"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:21.787Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:21.787Z"); return t}()),
		// 				UserName: to.Ptr("myCredential89"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential9"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential9"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:54.957Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:43:54.957Z"); return t}()),
		// 				UserName: to.Ptr("myCredential9"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential90"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential90"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:22.100Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:22.100Z"); return t}()),
		// 				UserName: to.Ptr("myCredential90"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential91"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential91"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:22.427Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:22.427Z"); return t}()),
		// 				UserName: to.Ptr("myCredential91"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential92"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential92"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:22.770Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:22.770Z"); return t}()),
		// 				UserName: to.Ptr("myCredential92"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential93"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential93"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:23.083Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:23.083Z"); return t}()),
		// 				UserName: to.Ptr("myCredential93"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential94"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential94"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:23.410Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:23.410Z"); return t}()),
		// 				UserName: to.Ptr("myCredential94"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential95"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential95"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:23.753Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:23.753Z"); return t}()),
		// 				UserName: to.Ptr("myCredential95"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential96"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential96"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:24.100Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:24.100Z"); return t}()),
		// 				UserName: to.Ptr("myCredential96"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCredential97"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount20/credentials/myCredential97"),
		// 			Properties: &armautomation.CredentialProperties{
		// 				Description: to.Ptr("my description goes here"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:24.427Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T21:44:24.427Z"); return t}()),
		// 				UserName: to.Ptr("myCredential97"),
		// 			},
		// 	}},
		// }
	}
}
