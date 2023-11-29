package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControl/getAllSourceControls.json
func ExampleSourceControlClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSourceControlClient().NewListByAutomationAccountPager("rg", "sampleAccount9", &armautomation.SourceControlClientListByAutomationAccountOptions{Filter: nil})
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
		// page.SourceControlListResult = armautomation.SourceControlListResult{
		// 	Value: []*armautomation.SourceControl{
		// 		{
		// 			Name: to.Ptr("sampleSourceControl1"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/sourcecontrols/sampleSourceControl1"),
		// 			Properties: &armautomation.SourceControlProperties{
		// 				Description: to.Ptr("my description"),
		// 				AutoSync: to.Ptr(true),
		// 				Branch: to.Ptr("master"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937Z"); return t}()),
		// 				FolderPath: to.Ptr("/sampleFolder/sampleFolder2"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937Z"); return t}()),
		// 				PublishRunbook: to.Ptr(true),
		// 				RepoURL: to.Ptr("https://github.com/SampleUserRepro/PowerShell-1"),
		// 				SourceType: to.Ptr(armautomation.SourceTypeGitHub),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleSourceControl2"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/sourcecontrols/sampleSourceControl2"),
		// 			Properties: &armautomation.SourceControlProperties{
		// 				Description: to.Ptr("my description"),
		// 				AutoSync: to.Ptr(true),
		// 				Branch: to.Ptr("master"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937Z"); return t}()),
		// 				FolderPath: to.Ptr("/sampleFolder/sampleFolder2"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937Z"); return t}()),
		// 				PublishRunbook: to.Ptr(true),
		// 				RepoURL: to.Ptr("https://github.com/SampleUserRepro/PowerShell-2"),
		// 				SourceType: to.Ptr(armautomation.SourceTypeGitHub),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleSourceControl3"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/sourcecontrols/sampleSourceControl3"),
		// 			Properties: &armautomation.SourceControlProperties{
		// 				Description: to.Ptr("my description"),
		// 				AutoSync: to.Ptr(true),
		// 				Branch: to.Ptr("master"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937Z"); return t}()),
		// 				FolderPath: to.Ptr("/sampleFolder/sampleFolder2"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937Z"); return t}()),
		// 				PublishRunbook: to.Ptr(true),
		// 				RepoURL: to.Ptr("https://github.com/SampleUserRepro/PowerShell-3"),
		// 				SourceType: to.Ptr(armautomation.SourceTypeGitHub),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleSourceControl4"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/sourcecontrols/sampleSourceControl4"),
		// 			Properties: &armautomation.SourceControlProperties{
		// 				Description: to.Ptr("my description"),
		// 				AutoSync: to.Ptr(true),
		// 				Branch: to.Ptr("master"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937Z"); return t}()),
		// 				FolderPath: to.Ptr("/sampleFolder/sampleFolder2"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937Z"); return t}()),
		// 				PublishRunbook: to.Ptr(true),
		// 				RepoURL: to.Ptr("https://github.com/SampleUserRepro/PowerShell-4"),
		// 				SourceType: to.Ptr(armautomation.SourceTypeGitHub),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleSourceControl5"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/sourcecontrols/sampleSourceControl5"),
		// 			Properties: &armautomation.SourceControlProperties{
		// 				Description: to.Ptr("my description"),
		// 				AutoSync: to.Ptr(true),
		// 				Branch: to.Ptr("master"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937Z"); return t}()),
		// 				FolderPath: to.Ptr("/sampleFolder/sampleFolder2"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937Z"); return t}()),
		// 				PublishRunbook: to.Ptr(true),
		// 				RepoURL: to.Ptr("https://github.com/SampleUserRepro/PowerShell-5"),
		// 				SourceType: to.Ptr(armautomation.SourceTypeGitHub),
		// 			},
		// 	}},
		// }
	}
}
