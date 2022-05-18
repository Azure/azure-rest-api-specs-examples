Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerregistry%2Farmcontainerregistry%2Fv0.6.0/sdk/resourcemanager/containerregistry/armcontainerregistry/README.md) on how to add the SDK to your project and authenticate.

```go
package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/ImportPipelineCreate.json
func ExampleImportPipelinesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewImportPipelinesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"myResourceGroup",
		"myRegistry",
		"myImportPipeline",
		armcontainerregistry.ImportPipeline{
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
```
