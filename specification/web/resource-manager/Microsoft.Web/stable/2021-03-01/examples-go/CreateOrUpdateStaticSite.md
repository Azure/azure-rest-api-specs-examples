```go
package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateOrUpdateStaticSite.json
func ExampleStaticSitesClient_BeginCreateOrUpdateStaticSite() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewStaticSitesClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdateStaticSite(ctx,
		"rg",
		"testStaticSite0",
		armappservice.StaticSiteARMResource{
			Location: to.Ptr("West US 2"),
			Properties: &armappservice.StaticSite{
				Branch: to.Ptr("master"),
				BuildProperties: &armappservice.StaticSiteBuildProperties{
					APILocation:         to.Ptr("api"),
					AppArtifactLocation: to.Ptr("build"),
					AppLocation:         to.Ptr("app"),
				},
				RepositoryToken: to.Ptr("repoToken123"),
				RepositoryURL:   to.Ptr("https://github.com/username/RepoName"),
			},
			SKU: &armappservice.SKUDescription{
				Name: to.Ptr("Basic"),
				Tier: to.Ptr("Basic"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappservice%2Farmappservice%2Fv1.0.0/sdk/resourcemanager/appservice/armappservice/README.md) on how to add the SDK to your project and authenticate.
