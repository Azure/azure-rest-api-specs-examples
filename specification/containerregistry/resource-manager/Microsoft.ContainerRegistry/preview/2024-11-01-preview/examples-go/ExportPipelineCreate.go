package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/300ff7c27c481d7074af06cd95a152aaea80ed2b/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2024-11-01-preview/examples/ExportPipelineCreate.json
func ExampleExportPipelinesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExportPipelinesClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", "myExportPipeline", armcontainerregistry.ExportPipeline{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExportPipeline = armcontainerregistry.ExportPipeline{
	// 	Name: to.Ptr("myExportPipeline"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/exportPipelines"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/exportPipelines/myExportPipeline"),
	// 	Identity: &armcontainerregistry.IdentityProperties{
	// 		Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("fa153151-b9fd-46f4-9088-5e6600f2689v"),
	// 		TenantID: to.Ptr("f686d426-8d16-42db-81b7-abu4gm510ccd"),
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcontainerregistry.ExportPipelineProperties{
	// 		Options: []*armcontainerregistry.PipelineOptions{
	// 			to.Ptr(armcontainerregistry.PipelineOptionsOverwriteBlobs)},
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 			Target: &armcontainerregistry.ExportPipelineTargetProperties{
	// 				Type: to.Ptr("AzureStorageBlobContainer"),
	// 				KeyVaultURI: to.Ptr("https://myvault.vault.azure.net/secrets/acrexportsas"),
	// 				URI: to.Ptr("https://accountname.blob.core.windows.net/containername"),
	// 			},
	// 		},
	// 	}
}
