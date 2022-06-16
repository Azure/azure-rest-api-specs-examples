package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-05-01-preview/examples/sourcecontrols/CreateSourceControl.json
func ExampleSourceControlsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewSourceControlsClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"myRg",
		"myWorkspace",
		"789e0c1f-4a3d-43ad-809c-e713b677b04a",
		armsecurityinsights.SourceControl{
			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
			Properties: &armsecurityinsights.SourceControlProperties{
				Description: to.Ptr("This is a source control"),
				ContentTypes: []*armsecurityinsights.ContentType{
					to.Ptr(armsecurityinsights.ContentType("AnalyticRules")),
					to.Ptr(armsecurityinsights.ContentTypeWorkbook)},
				DisplayName: to.Ptr("My Source Control"),
				RepoType:    to.Ptr(armsecurityinsights.RepoTypeGithub),
				Repository: &armsecurityinsights.Repository{
					Branch:     to.Ptr("master"),
					DisplayURL: to.Ptr("https://github.com/user/repo"),
					PathMapping: []*armsecurityinsights.ContentPathMap{
						{
							Path:        to.Ptr("path/to/rules"),
							ContentType: to.Ptr(armsecurityinsights.ContentType("AnalyticRules")),
						},
						{
							Path:        to.Ptr("path/to/workbooks"),
							ContentType: to.Ptr(armsecurityinsights.ContentTypeWorkbook),
						}},
					URL: to.Ptr("https://github.com/user/repo"),
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
