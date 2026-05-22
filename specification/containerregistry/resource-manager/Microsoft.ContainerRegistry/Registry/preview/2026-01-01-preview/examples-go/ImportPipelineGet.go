package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: 2026-01-01-preview/ImportPipelineGet.json
func ExampleImportPipelinesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewImportPipelinesClient().Get(ctx, "myResourceGroup", "myRegistry", "myImportPipeline", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerregistry.ImportPipelinesClientGetResponse{
	// 	ImportPipeline: &armcontainerregistry.ImportPipeline{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/importPipelines/myImportPipeline"),
	// 		Name: to.Ptr("myImportPipeline"),
	// 		Type: to.Ptr("Microsoft.ContainerRegistry/registries/importPipelines"),
	// 		Properties: &armcontainerregistry.ImportPipelineProperties{
	// 			Source: &armcontainerregistry.ImportPipelineSourceProperties{
	// 				Type: to.Ptr(armcontainerregistry.PipelineSourceTypeAzureStorageBlobContainer),
	// 				URI: to.Ptr("https://accountname.blob.core.windows.net/containername"),
	// 				KeyVaultURI: to.Ptr("https://myvault.vault.azure.net/secrets/acrimportsas"),
	// 				StorageAccessMode: to.Ptr(armcontainerregistry.StorageAccessModeSasToken),
	// 			},
	// 			Trigger: &armcontainerregistry.PipelineTriggerProperties{
	// 				SourceTrigger: &armcontainerregistry.PipelineSourceTriggerProperties{
	// 					Status: to.Ptr(armcontainerregistry.TriggerStatusEnabled),
	// 				},
	// 			},
	// 			Options: []*armcontainerregistry.PipelineOptions{
	// 				to.Ptr(armcontainerregistry.PipelineOptionsOverwriteTags),
	// 				to.Ptr(armcontainerregistry.PipelineOptionsDeleteSourceBlobOnSuccess),
	// 				to.Ptr(armcontainerregistry.PipelineOptionsContinueOnErrors),
	// 			},
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Identity: &armcontainerregistry.IdentityProperties{
	// 			Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armcontainerregistry.UserIdentityProperties{
	// 				"/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2": &armcontainerregistry.UserIdentityProperties{
	// 					ClientID: to.Ptr("d3ce1bc2-f7d7-4a5b-9979-950f4e57680e"),
	// 					PrincipalID: to.Ptr("b6p9f58b-6fbf-4efd-a7e0-fvd46911a466"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
