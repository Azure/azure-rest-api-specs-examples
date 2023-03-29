package armdeploymentmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/artifactsource_createorupdate_artifactroot.json
func ExampleArtifactSourcesClient_CreateOrUpdate_createArtifactSourceWithArtifactRootAnOffsetIntoTheStorageContainer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewArtifactSourcesClient().CreateOrUpdate(ctx, "myResourceGroup", "myArtifactSource", &armdeploymentmanager.ArtifactSourcesClientCreateOrUpdateOptions{ArtifactSourceInfo: &armdeploymentmanager.ArtifactSource{
		Location: to.Ptr("centralus"),
		Tags:     map[string]*string{},
		Properties: &armdeploymentmanager.ArtifactSourceProperties{
			ArtifactRoot: to.Ptr("1.0.0.0"),
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
