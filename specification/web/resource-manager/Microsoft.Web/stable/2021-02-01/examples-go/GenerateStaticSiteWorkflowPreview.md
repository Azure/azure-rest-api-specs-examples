```go
package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/GenerateStaticSiteWorkflowPreview.json
func ExampleStaticSitesClient_PreviewWorkflow() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewStaticSitesClient("<subscription-id>", cred, nil)
	res, err := client.PreviewWorkflow(ctx,
		"<location>",
		armappservice.StaticSitesWorkflowPreviewRequest{
			Properties: &armappservice.StaticSitesWorkflowPreviewRequestProperties{
				Branch: to.StringPtr("<branch>"),
				BuildProperties: &armappservice.StaticSiteBuildProperties{
					APILocation:         to.StringPtr("<apilocation>"),
					AppArtifactLocation: to.StringPtr("<app-artifact-location>"),
					AppLocation:         to.StringPtr("<app-location>"),
				},
				RepositoryURL: to.StringPtr("<repository-url>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("StaticSitesWorkflowPreview.ID: %s\n", *res.ID)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappservice%2Farmappservice%2Fv0.1.0/sdk/resourcemanager/appservice/armappservice/README.md) on how to add the SDK to your project and authenticate.
