package armcontainerregistry_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2021-12-01-preview/examples/ExportPipelineCreate.json
func ExampleExportPipelinesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewExportPipelinesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<registry-name>",
		"<export-pipeline-name>",
		armcontainerregistry.ExportPipeline{
			Identity: &armcontainerregistry.IdentityProperties{
				Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeSystemAssigned),
			},
			Location: to.Ptr("<location>"),
			Properties: &armcontainerregistry.ExportPipelineProperties{
				Options: []*armcontainerregistry.PipelineOptions{
					to.Ptr(armcontainerregistry.PipelineOptionsOverwriteBlobs)},
				Target: &armcontainerregistry.ExportPipelineTargetProperties{
					Type:        to.Ptr("<type>"),
					KeyVaultURI: to.Ptr("<key-vault-uri>"),
					URI:         to.Ptr("<uri>"),
				},
			},
		},
		&armcontainerregistry.ExportPipelinesClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
