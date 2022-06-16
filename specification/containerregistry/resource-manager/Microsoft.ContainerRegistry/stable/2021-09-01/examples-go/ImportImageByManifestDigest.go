package armcontainerregistry_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/examples/ImportImageByManifestDigest.json
func ExampleRegistriesClient_BeginImportImage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerregistry.NewRegistriesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginImportImage(ctx,
		"<resource-group-name>",
		"<registry-name>",
		armcontainerregistry.ImportImageParameters{
			Mode: armcontainerregistry.ImportMode("Force").ToPtr(),
			Source: &armcontainerregistry.ImportSource{
				ResourceID:  to.StringPtr("<resource-id>"),
				SourceImage: to.StringPtr("<source-image>"),
			},
			TargetTags: []*string{
				to.StringPtr("targetRepository:targetTag")},
			UntaggedTargetRepositories: []*string{
				to.StringPtr("targetRepository1")},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
