package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-08-01-preview/examples/PipelineRunList.json
func ExamplePipelineRunsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.PipelineRunListResult = armcontainerregistry.PipelineRunListResult{
		// 	Value: []*armcontainerregistry.PipelineRun{
		// 		{
		// 			Name: to.Ptr("myPipelineRun"),
		// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/pipelineRuns"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/pipelineRuns/myPipelineRun"),
		// 			Properties: &armcontainerregistry.PipelineRunProperties{
		// 				ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 				Response: &armcontainerregistry.PipelineRunResponse{
		// 					CatalogDigest: to.Ptr("sha256@"),
		// 					Progress: &armcontainerregistry.ProgressProperties{
		// 						Percentage: to.Ptr("20"),
		// 					},
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-04T17:23:21.9261521+00:00"); return t}()),
		// 					Status: to.Ptr("Running"),
		// 					Target: &armcontainerregistry.ExportPipelineTargetProperties{
		// 						Type: to.Ptr("AzureStorageBlob"),
		// 						KeyVaultURI: to.Ptr("https://myvault.vault.azure.net/secrets/acrexportsas"),
		// 						URI: to.Ptr("https://accountname.blob.core.windows.net/containername/myblob.tar.gz"),
		// 					},
		// 				},
		// 				Request: &armcontainerregistry.PipelineRunRequest{
		// 					Artifacts: []*string{
		// 						to.Ptr("sourceRepository/hello-world"),
		// 						to.Ptr("sourceRepository2@sha256:00000000000000000000000000000000000")},
		// 						PipelineResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/exportPipelines/myExportPipeline"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myPipelineRun"),
		// 				Type: to.Ptr("Microsoft.ContainerRegistry/registries/pipelineRuns"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/pipelineRuns/myPipelineRun"),
		// 				Properties: &armcontainerregistry.PipelineRunProperties{
		// 					ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 					Response: &armcontainerregistry.PipelineRunResponse{
		// 						CatalogDigest: to.Ptr("sha256@"),
		// 						ImportedArtifacts: []*string{
		// 							to.Ptr("sourceRepository/hello-world"),
		// 							to.Ptr("sourceRepository2@sha256:00000000000000000000000000000000000")},
		// 							Progress: &armcontainerregistry.ProgressProperties{
		// 								Percentage: to.Ptr("100"),
		// 							},
		// 							Source: &armcontainerregistry.ImportPipelineSourceProperties{
		// 								Type: to.Ptr(armcontainerregistry.PipelineSourceType("AzureStorageBlob")),
		// 								KeyVaultURI: to.Ptr("https://myvault.vault.azure.net/secrets/acrimportsas"),
		// 								URI: to.Ptr("https://accountname.blob.core.windows.net/containername/myblob.tar.gz"),
		// 							},
		// 							StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-03T17:23:21.9261521+00:00"); return t}()),
		// 							Status: to.Ptr("Succeeded"),
		// 						},
		// 						Request: &armcontainerregistry.PipelineRunRequest{
		// 							PipelineResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/importPipelines/myImportPipeline"),
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
