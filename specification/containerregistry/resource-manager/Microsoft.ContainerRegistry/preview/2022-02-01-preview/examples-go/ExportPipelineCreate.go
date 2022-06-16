package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/ExportPipelineCreate.json
func ExampleExportPipelinesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewExportPipelinesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"myResourceGroup",
		"myRegistry",
		"myExportPipeline",
		armcontainerregistry.ExportPipeline{
			Identity: &armcontainerregistry.IdentityProperties{
				Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeSystemAssigned),
			},
			Location: to.Ptr("westus"),
			Properties: &armcontainerregistry.ExportPipelineProperties{
				Options: []*armcontainerregistry.PipelineOptions{
					to.Ptr(armcontainerregistry.PipelineOptionsOverwriteBlobs)},
				Target: &armcontainerregistry.ExportPipelineTargetProperties{
					Type:        to.Ptr("AzureStorageBlobContainer"),
					KeyVaultURI: to.Ptr("https://myvault.vault.azure.net/secrets/acrexportsas"),
					URI:         to.Ptr("https://accountname.blob.core.windows.net/containername"),
				},
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
