package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: 2026-03-01-preview/PipelineRunList.json
func ExamplePipelineRunsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPipelineRunsClient().NewListPager("myResourceGroup", "myRegistry", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcontainerregistry.PipelineRunsClientListResponse{
		// 	PipelineRunListResult: armcontainerregistry.PipelineRunListResult{
		// 		Value: []*armcontainerregistry.PipelineRun{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/pipelineRuns/myPipelineRun"),
		// 				Name: to.Ptr("myPipelineRun"),
		// 				Type: to.Ptr("Microsoft.ContainerRegistry/registries/pipelineRuns"),
		// 				Properties: &armcontainerregistry.PipelineRunProperties{
		// 					Request: &armcontainerregistry.PipelineRunRequest{
		// 						PipelineResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/exportPipelines/myExportPipeline"),
		// 						Artifacts: []*string{
		// 							to.Ptr("sourceRepository/hello-world"),
		// 							to.Ptr("sourceRepository2@sha256:00000000000000000000000000000000000"),
		// 						},
		// 					},
		// 					Response: &armcontainerregistry.PipelineRunResponse{
		// 						Status: to.Ptr("Running"),
		// 						CatalogDigest: to.Ptr("sha256@"),
		// 						Progress: &armcontainerregistry.ProgressProperties{
		// 							Percentage: to.Ptr("20"),
		// 						},
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-04T17:23:21.9261521+00:00"); return t}()),
		// 						Target: &armcontainerregistry.ExportPipelineTargetProperties{
		// 							Type: to.Ptr("AzureStorageBlob"),
		// 							URI: to.Ptr("https://accountname.blob.core.windows.net/containername/myblob.tar.gz"),
		// 							KeyVaultURI: to.Ptr("https://myvault.vault.azure.net/secrets/acrexportsas"),
		// 							StorageAccessMode: to.Ptr(armcontainerregistry.StorageAccessModeSasToken),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/pipelineRuns/myPipelineRun"),
		// 				Name: to.Ptr("myPipelineRun"),
		// 				Type: to.Ptr("Microsoft.ContainerRegistry/registries/pipelineRuns"),
		// 				Properties: &armcontainerregistry.PipelineRunProperties{
		// 					Request: &armcontainerregistry.PipelineRunRequest{
		// 						PipelineResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/importPipelines/myImportPipeline"),
		// 					},
		// 					Response: &armcontainerregistry.PipelineRunResponse{
		// 						Status: to.Ptr("Succeeded"),
		// 						Progress: &armcontainerregistry.ProgressProperties{
		// 							Percentage: to.Ptr("100"),
		// 						},
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-03T17:23:21.9261521+00:00"); return t}()),
		// 						ImportedArtifacts: []*string{
		// 							to.Ptr("sourceRepository/hello-world"),
		// 							to.Ptr("sourceRepository2@sha256:00000000000000000000000000000000000"),
		// 						},
		// 						Source: &armcontainerregistry.ImportPipelineSourceProperties{
		// 							Type: to.Ptr(armcontainerregistry.PipelineSourceType("AzureStorageBlob")),
		// 							URI: to.Ptr("https://accountname.blob.core.windows.net/containername/myblob.tar.gz"),
		// 							KeyVaultURI: to.Ptr("https://myvault.vault.azure.net/secrets/acrimportsas"),
		// 							StorageAccessMode: to.Ptr(armcontainerregistry.StorageAccessModeSasToken),
		// 						},
		// 						CatalogDigest: to.Ptr("sha256@"),
		// 					},
		// 					ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
