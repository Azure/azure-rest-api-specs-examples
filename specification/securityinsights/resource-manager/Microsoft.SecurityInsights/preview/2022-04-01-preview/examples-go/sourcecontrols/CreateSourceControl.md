```go
package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/sourcecontrols/CreateSourceControl.json
func ExampleSourceControlsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewSourceControlsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<source-control-id>",
		armsecurityinsights.SourceControl{
			Etag: to.Ptr("<etag>"),
			Properties: &armsecurityinsights.SourceControlProperties{
				Description: to.Ptr("<description>"),
				ContentTypes: []*armsecurityinsights.ContentType{
					to.Ptr(armsecurityinsights.ContentType("AnalyticRules")),
					to.Ptr(armsecurityinsights.ContentTypeWorkbook)},
				DisplayName: to.Ptr("<display-name>"),
				RepoType:    to.Ptr(armsecurityinsights.RepoTypeGithub),
				Repository: &armsecurityinsights.Repository{
					Branch:     to.Ptr("<branch>"),
					DisplayURL: to.Ptr("<display-url>"),
					PathMapping: []*armsecurityinsights.ContentPathMap{
						{
							Path:        to.Ptr("<path>"),
							ContentType: to.Ptr(armsecurityinsights.ContentType("AnalyticRules")),
						},
						{
							Path:        to.Ptr("<path>"),
							ContentType: to.Ptr(armsecurityinsights.ContentTypeWorkbook),
						}},
					URL: to.Ptr("<url>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsights%2Farmsecurityinsights%2Fv0.3.0/sdk/resourcemanager/securityinsights/armsecurityinsights/README.md) on how to add the SDK to your project and authenticate.
