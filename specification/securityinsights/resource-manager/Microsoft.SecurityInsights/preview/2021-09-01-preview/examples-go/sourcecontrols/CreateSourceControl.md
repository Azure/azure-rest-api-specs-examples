```go
package armsecurityinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/sourcecontrols/CreateSourceControl.json
func ExampleSourceControlsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewSourceControlsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<source-control-id>",
		armsecurityinsight.SourceControl{
			Etag: to.StringPtr("<etag>"),
			Properties: &armsecurityinsight.SourceControlProperties{
				Description: to.StringPtr("<description>"),
				ContentTypes: []*armsecurityinsight.ContentType{
					armsecurityinsight.ContentType("AnalyticRules").ToPtr(),
					armsecurityinsight.ContentType("Workbook").ToPtr()},
				DisplayName: to.StringPtr("<display-name>"),
				RepoType:    armsecurityinsight.RepoType("Github").ToPtr(),
				Repository: &armsecurityinsight.Repository{
					Branch:     to.StringPtr("<branch>"),
					DisplayURL: to.StringPtr("<display-url>"),
					PathMapping: []*armsecurityinsight.ContentPathMap{
						{
							Path:        to.StringPtr("<path>"),
							ContentType: armsecurityinsight.ContentType("AnalyticRules").ToPtr(),
						},
						{
							Path:        to.StringPtr("<path>"),
							ContentType: armsecurityinsight.ContentType("Workbook").ToPtr(),
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsight%2Farmsecurityinsight%2Fv0.2.1/sdk/resourcemanager/securityinsight/armsecurityinsight/README.md) on how to add the SDK to your project and authenticate.
