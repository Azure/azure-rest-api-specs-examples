package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ImportImageByManifestDigest.json
func ExampleRegistriesClient_BeginImportImage_importImageByManifestDigest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewRegistriesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginImportImage(ctx, "myResourceGroup", "myRegistry", armcontainerregistry.ImportImageParameters{
		Mode: to.Ptr(armcontainerregistry.ImportModeForce),
		Source: &armcontainerregistry.ImportSource{
			ResourceID:  to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/sourceResourceGroup/providers/Microsoft.ContainerRegistry/registries/sourceRegistry"),
			SourceImage: to.Ptr("sourceRepository@sha256:0000000000000000000000000000000000000000000000000000000000000000"),
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
