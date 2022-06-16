package armcontainerregistry_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2021-12-01-preview/examples/ImportImageByManifestDigest.json
func ExampleRegistriesClient_BeginImportImage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewRegistriesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginImportImage(ctx,
		"<resource-group-name>",
		"<registry-name>",
		armcontainerregistry.ImportImageParameters{
			Mode: to.Ptr(armcontainerregistry.ImportModeForce),
			Source: &armcontainerregistry.ImportSource{
				ResourceID:  to.Ptr("<resource-id>"),
				SourceImage: to.Ptr("<source-image>"),
			},
			TargetTags: []*string{
				to.Ptr("targetRepository:targetTag")},
			UntaggedTargetRepositories: []*string{
				to.Ptr("targetRepository1")},
		},
		&armcontainerregistry.RegistriesClientBeginImportImageOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
