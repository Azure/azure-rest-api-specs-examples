package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dc4c1eaef16e0bc8b1e96c3d1e014deb96259b35/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2025-03-01-preview/examples/ImportPipelineCreate.json
func ExampleImportPipelinesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewImportPipelinesClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", "myImportPipeline", armcontainerregistry.ImportPipeline{
		Identity: &armcontainerregistry.IdentityProperties{
			Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armcontainerregistry.UserIdentityProperties{
				"/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2": {},
			},
		},
		Location: to.Ptr("westus"),
		Properties: &armcontainerregistry.ImportPipelineProperties{
			Options: []*armcontainerregistry.PipelineOptions{
				to.Ptr(armcontainerregistry.PipelineOptionsOverwriteTags),
				to.Ptr(armcontainerregistry.PipelineOptionsDeleteSourceBlobOnSuccess),
				to.Ptr(armcontainerregistry.PipelineOptionsContinueOnErrors)},
			Source: &armcontainerregistry.ImportPipelineSourceProperties{
				Type:        to.Ptr(armcontainerregistry.PipelineSourceTypeAzureStorageBlobContainer),
				KeyVaultURI: to.Ptr("https://myvault.vault.azure.net/secrets/acrimportsas"),
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
	// res.ImportPipeline = armcontainerregistry.ImportPipeline{
	// 	Name: to.Ptr("myImportPipeline"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/importPipelines"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/importPipelines/myImportPipeline"),
	// 	Identity: &armcontainerregistry.IdentityProperties{
	// 		Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armcontainerregistry.UserIdentityProperties{
	// 			"/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2": &armcontainerregistry.UserIdentityProperties{
	// 				ClientID: to.Ptr("d3ce1bc2-f7d7-4a5b-9979-950f4e57680e"),
	// 				PrincipalID: to.Ptr("b6p9f58b-6fbf-4efd-a7e0-fvd46911a466"),
	// 			},
	// 		},
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcontainerregistry.ImportPipelineProperties{
	// 		Options: []*armcontainerregistry.PipelineOptions{
	// 			to.Ptr(armcontainerregistry.PipelineOptionsOverwriteTags),
	// 			to.Ptr(armcontainerregistry.PipelineOptionsDeleteSourceBlobOnSuccess),
	// 			to.Ptr(armcontainerregistry.PipelineOptionsContinueOnErrors)},
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 			Source: &armcontainerregistry.ImportPipelineSourceProperties{
	// 				Type: to.Ptr(armcontainerregistry.PipelineSourceTypeAzureStorageBlobContainer),
	// 				KeyVaultURI: to.Ptr("https://myvault.vault.azure.net/secrets/acrimportsas"),
	// 				URI: to.Ptr("https://accountname.blob.core.windows.net/containername"),
	// 			},
	// 			Trigger: &armcontainerregistry.PipelineTriggerProperties{
	// 				SourceTrigger: &armcontainerregistry.PipelineSourceTriggerProperties{
	// 					Status: to.Ptr(armcontainerregistry.TriggerStatusEnabled),
	// 				},
	// 			},
	// 		},
	// 	}
}
