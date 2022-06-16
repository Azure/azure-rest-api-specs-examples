package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/sourcecontrols/CreateSourceControl.json
func ExampleSourceControlsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewSourceControlsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<source-control-id>",
		armsecurityinsights.SourceControl{
			Etag: to.StringPtr("<etag>"),
			Properties: &armsecurityinsights.SourceControlProperties{
				Description: to.StringPtr("<description>"),
				ContentTypes: []*armsecurityinsights.ContentType{
					armsecurityinsights.ContentType("AnalyticRules").ToPtr(),
					armsecurityinsights.ContentType("Workbook").ToPtr()},
				DisplayName: to.StringPtr("<display-name>"),
				RepoType:    armsecurityinsights.RepoType("Github").ToPtr(),
				Repository: &armsecurityinsights.Repository{
					Branch:     to.StringPtr("<branch>"),
					DisplayURL: to.StringPtr("<display-url>"),
					PathMapping: []*armsecurityinsights.ContentPathMap{
						{
							Path:        to.StringPtr("<path>"),
							ContentType: armsecurityinsights.ContentType("AnalyticRules").ToPtr(),
						},
						{
							Path:        to.StringPtr("<path>"),
							ContentType: armsecurityinsights.ContentType("Workbook").ToPtr(),
						}},
					URL: to.StringPtr("<url>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SourceControlsClientCreateResult)
}
