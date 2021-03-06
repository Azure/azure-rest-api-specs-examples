package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControl/updateSourceControl_patch.json
func ExampleSourceControlClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomation.NewSourceControlClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"rg",
		"sampleAccount9",
		"sampleSourceControl",
		armautomation.SourceControlUpdateParameters{
			Properties: &armautomation.SourceControlUpdateProperties{
				Description:    to.Ptr("my description"),
				AutoSync:       to.Ptr(true),
				Branch:         to.Ptr("master"),
				FolderPath:     to.Ptr("/folderOne/folderTwo"),
				PublishRunbook: to.Ptr(true),
				SecurityToken: &armautomation.SourceControlSecurityTokenProperties{
					AccessToken: to.Ptr("3a326f7a0dcd343ea58fee21f2fd5fb4c1234567"),
					TokenType:   to.Ptr(armautomation.TokenTypePersonalAccessToken),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
