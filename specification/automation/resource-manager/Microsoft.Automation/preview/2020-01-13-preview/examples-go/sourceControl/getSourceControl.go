package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControl/getSourceControl.json
func ExampleSourceControlClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSourceControlClient().Get(ctx, "rg", "sampleAccount9", "sampleSourceControl", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SourceControl = armautomation.SourceControl{
	// 	Name: to.Ptr("sampleSourceControl"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/sourcecontrols/sampleSourceControl"),
	// 	Properties: &armautomation.SourceControlProperties{
	// 		Description: to.Ptr("my description"),
	// 		AutoSync: to.Ptr(true),
	// 		Branch: to.Ptr("master"),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937+00:00"); return t}()),
	// 		FolderPath: to.Ptr("/folderOne/folderTwo"),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937+00:00"); return t}()),
	// 		PublishRunbook: to.Ptr(true),
	// 		RepoURL: to.Ptr("https://github.com/SampleUserRepro/PowerShell"),
	// 		SourceType: to.Ptr(armautomation.SourceTypeGitHub),
	// 	},
	// }
}
