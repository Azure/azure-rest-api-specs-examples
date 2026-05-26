package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/sourceControl/createOrUpdateSourceControl.json
func ExampleSourceControlClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
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
				AccessToken: to.Ptr("******"),
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
	// res = armautomation.SourceControlClientCreateOrUpdateResponse{
	// 	SourceControl: armautomation.SourceControl{
	// 		Name: to.Ptr("sampleSourceControl"),
	// 		ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/sourcecontrols/sampleSourceControl"),
	// 		Properties: &armautomation.SourceControlProperties{
	// 			Description: to.Ptr("my description"),
	// 			AutoSync: to.Ptr(true),
	// 			Branch: to.Ptr("master"),
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937+00:00"); return t}()),
	// 			FolderPath: to.Ptr("/folderOne/folderTwo"),
	// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937+00:00"); return t}()),
	// 			PublishRunbook: to.Ptr(true),
	// 			RepoURL: to.Ptr("https://sampleUser.visualstudio.com/myProject/_git/myRepository"),
	// 			SourceType: to.Ptr(armautomation.SourceTypeVsoGit),
	// 		},
	// 	},
	// }
}
