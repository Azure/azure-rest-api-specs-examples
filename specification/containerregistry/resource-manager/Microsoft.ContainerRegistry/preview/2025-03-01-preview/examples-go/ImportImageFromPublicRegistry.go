package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dc4c1eaef16e0bc8b1e96c3d1e014deb96259b35/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2025-03-01-preview/examples/ImportImageFromPublicRegistry.json
func ExampleRegistriesClient_BeginImportImage_importImageFromPublicRegistry() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRegistriesClient().BeginImportImage(ctx, "myResourceGroup", "myRegistry", armcontainerregistry.ImportImageParameters{
		Mode: to.Ptr(armcontainerregistry.ImportModeForce),
		Source: &armcontainerregistry.ImportSource{
			RegistryURI: to.Ptr("registry.hub.docker.com"),
			SourceImage: to.Ptr("library/hello-world"),
		},
		TargetTags: []*string{
			to.Ptr("targetRepository:targetTag")},
		UntaggedTargetRepositories: []*string{
			to.Ptr("targetRepository1")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
