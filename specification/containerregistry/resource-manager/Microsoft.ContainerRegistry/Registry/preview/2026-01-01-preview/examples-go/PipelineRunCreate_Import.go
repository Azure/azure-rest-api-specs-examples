package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: 2026-01-01-preview/PipelineRunCreate_Import.json
func ExamplePipelineRunsClient_BeginCreate_pipelineRunCreateImport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPipelineRunsClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", "myPipelineRun", armcontainerregistry.PipelineRun{
		Properties: &armcontainerregistry.PipelineRunProperties{
			Request: &armcontainerregistry.PipelineRunRequest{
				PipelineResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/importPipelines/myImportPipeline"),
				Source: &armcontainerregistry.PipelineRunSourceProperties{
					Type: to.Ptr(armcontainerregistry.PipelineRunSourceTypeAzureStorageBlob),
					Name: to.Ptr("myblob.tar.gz"),
				},
				CatalogDigest: to.Ptr("sha256@"),
			},
			ForceUpdateTag: to.Ptr("2020-03-04T17:23:21.9261521+00:00"),
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
	// res = armcontainerregistry.PipelineRunsClientCreateResponse{
	// 	PipelineRun: &armcontainerregistry.PipelineRun{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/pipelineRuns/myPipelineRun"),
	// 		Name: to.Ptr("myPipelineRun"),
	// 		Type: to.Ptr("Microsoft.ContainerRegistry/registries/pipelineRuns"),
	// 		Properties: &armcontainerregistry.PipelineRunProperties{
	// 			Request: &armcontainerregistry.PipelineRunRequest{
	// 				PipelineResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/importPipelines/myImportPipeline"),
	// 			},
	// 			Response: &armcontainerregistry.PipelineRunResponse{
	// 				Status: to.Ptr("Succeeded"),
	// 				Progress: &armcontainerregistry.ProgressProperties{
	// 					Percentage: to.Ptr("100"),
	// 				},
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-04T17:23:21.9261521+00:00"); return t}()),
	// 				ImportedArtifacts: []*string{
	// 					to.Ptr("sourceRepository/hello-world"),
	// 					to.Ptr("sourceRepository2@sha256:00000000000000000000000000000000000"),
	// 				},
	// 				Source: &armcontainerregistry.ImportPipelineSourceProperties{
	// 					Type: to.Ptr(armcontainerregistry.PipelineSourceType("AzureStorageBlob")),
	// 					URI: to.Ptr("https://accountname.blob.core.windows.net/containername/myblob.tar.gz"),
	// 					KeyVaultURI: to.Ptr("https://myvault.vault.azure.net/secrets/acrimportsas"),
	// 					StorageAccessMode: to.Ptr(armcontainerregistry.StorageAccessModeSasToken),
	// 				},
	// 				CatalogDigest: to.Ptr("sha256@"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
