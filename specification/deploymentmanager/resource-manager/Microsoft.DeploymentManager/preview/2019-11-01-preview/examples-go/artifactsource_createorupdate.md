Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdeploymentmanager%2Farmdeploymentmanager%2Fv0.4.0/sdk/resourcemanager/deploymentmanager/armdeploymentmanager/README.md) on how to add the SDK to your project and authenticate.

```go
package armdeploymentmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/artifactsource_createorupdate.json
func ExampleArtifactSourcesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdeploymentmanager.NewArtifactSourcesClient("caac1590-e859-444f-a9e0-62091c0f5929", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.CreateOrUpdate(ctx,
		"myResourceGroup",
		"myArtifactSource",
		&armdeploymentmanager.ArtifactSourcesClientCreateOrUpdateOptions{ArtifactSourceInfo: &armdeploymentmanager.ArtifactSource{
			Location: to.Ptr("centralus"),
			Tags:     map[string]*string{},
			Properties: &armdeploymentmanager.ArtifactSourceProperties{
				Authentication: &armdeploymentmanager.SasAuthentication{
					Type: to.Ptr("Sas"),
					Properties: &armdeploymentmanager.SasProperties{
						SasURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myartifactsource?st=2018-07-07T14%3A10%3A00Z&se=2019-12-31T15%3A10%3A00Z&sp=rl&sv=2017-04-17&sr=c&sig=Yh2SoJ1NhhLRwCLln7de%2Fkabcdefghijklmno5sWEIk%3D"),
					},
				},
				SourceType: to.Ptr("AzureStorage"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
```
