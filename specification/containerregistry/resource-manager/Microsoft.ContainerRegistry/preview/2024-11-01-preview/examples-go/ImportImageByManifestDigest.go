package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/300ff7c27c481d7074af06cd95a152aaea80ed2b/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2024-11-01-preview/examples/ImportImageByManifestDigest.json
func ExampleRegistriesClient_BeginImportImage_importImageByManifestDigest() {
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
