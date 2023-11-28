package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControl/createOrUpdateSourceControl.json
func ExampleSourceControlClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSourceControlClient().CreateOrUpdate(ctx, "rg", "sampleAccount9", "sampleSourceControl", armautomation.SourceControlCreateOrUpdateParameters{
		Properties: &armautomation.SourceControlCreateOrUpdateProperties{
			Description:    to.Ptr("my description"),
			AutoSync:       to.Ptr(true),
			Branch:         to.Ptr("master"),
			FolderPath:     to.Ptr("/folderOne/folderTwo"),
			PublishRunbook: to.Ptr(true),
			RepoURL:        to.Ptr("https://sampleUser.visualstudio.com/myProject/_git/myRepository"),
			SecurityToken: &armautomation.SourceControlSecurityTokenProperties{
				AccessToken: to.Ptr("3a326f7a0dcd343ea58fee21f2fd5fb4c1234567"),
				TokenType:   to.Ptr(armautomation.TokenTypePersonalAccessToken),
			},
			SourceType: to.Ptr(armautomation.SourceTypeVsoGit),
		},
	}, nil)
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
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937Z"); return t}()),
	// 		FolderPath: to.Ptr("/folderOne/folderTwo"),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937Z"); return t}()),
	// 		PublishRunbook: to.Ptr(true),
	// 		RepoURL: to.Ptr("https://sampleUser.visualstudio.com/myProject/_git/myRepository"),
	// 		SourceType: to.Ptr(armautomation.SourceTypeVsoGit),
	// 	},
	// }
}
