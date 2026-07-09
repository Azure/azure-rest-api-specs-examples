package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: 2026-03-01-preview/ExportPipelineCreate.json
func ExampleExportPipelinesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExportPipelinesClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", "myExportPipeline", armcontainerregistry.ExportPipeline{
		Location: to.Ptr("westus"),
		Identity: &armcontainerregistry.IdentityProperties{
			Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeSystemAssigned),
		},
		Properties: &armcontainerregistry.ExportPipelineProperties{
			Target: &armcontainerregistry.ExportPipelineTargetProperties{
				Type:              to.Ptr("AzureStorageBlobContainer"),
				URI:               to.Ptr("https://accountname.blob.core.windows.net/containername"),
				KeyVaultURI:       to.Ptr("https://myvault.vault.azure.net/secrets/acrexportsas"),
				StorageAccessMode: to.Ptr(armcontainerregistry.StorageAccessModeSasToken),
			},
			Options: []*armcontainerregistry.PipelineOptions{
				to.Ptr(armcontainerregistry.PipelineOptionsOverwriteBlobs),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerregistry.ExportPipelinesClientCreateResponse{
	// 	ExportPipeline: armcontainerregistry.ExportPipeline{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/exportPipelines/myExportPipeline"),
	// 		Name: to.Ptr("myExportPipeline"),
	// 		Type: to.Ptr("Microsoft.ContainerRegistry/registries/exportPipelines"),
	// 		Properties: &armcontainerregistry.ExportPipelineProperties{
	// 			Target: &armcontainerregistry.ExportPipelineTargetProperties{
	// 				Type: to.Ptr("AzureStorageBlobContainer"),
	// 				URI: to.Ptr("https://accountname.blob.core.windows.net/containername"),
	// 				KeyVaultURI: to.Ptr("https://myvault.vault.azure.net/secrets/acrexportsas"),
	// 				StorageAccessMode: to.Ptr(armcontainerregistry.StorageAccessModeSasToken),
	// 			},
	// 			Options: []*armcontainerregistry.PipelineOptions{
	// 				to.Ptr(armcontainerregistry.PipelineOptionsOverwriteBlobs),
	// 			},
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Identity: &armcontainerregistry.IdentityProperties{
	// 			PrincipalID: to.Ptr("fa153151-b9fd-46f4-9088-5e6600f2689v"),
	// 			TenantID: to.Ptr("f686d426-8d16-42db-81b7-abu4gm510ccd"),
	// 			Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeSystemAssigned),
	// 		},
	// 	},
	// }
}
