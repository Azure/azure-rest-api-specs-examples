package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ee6d9fd7687d4b67117c5a167c191a7e7e70b53/specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/Endpoints_CreateOrUpdate_AzureStorageBlobContainer.json
func ExampleEndpointsClient_CreateOrUpdate_endpointsCreateOrUpdateAzureStorageBlobContainer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().CreateOrUpdate(ctx, "examples-rg", "examples-storageMoverName", "examples-endpointName", armstoragemover.Endpoint{
		Properties: &armstoragemover.AzureStorageBlobContainerEndpointProperties{
			Description:              to.Ptr("Example Storage Blob Container Endpoint Description"),
			EndpointType:             to.Ptr(armstoragemover.EndpointTypeAzureStorageBlobContainer),
			BlobContainerName:        to.Ptr("examples-blobcontainer"),
			StorageAccountResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.Storage/storageAccounts/examplesa"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Endpoint = armstoragemover.Endpoint{
	// 	Name: to.Ptr("examples-endpointName"),
	// 	Type: to.Ptr("Microsoft.StorageMover/storageMovers/endpoints"),
	// 	ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-endpointName"),
	// 	Properties: &armstoragemover.AzureStorageBlobContainerEndpointProperties{
	// 		Description: to.Ptr("Example Storage Blob Container Endpoint Description"),
	// 		EndpointType: to.Ptr(armstoragemover.EndpointTypeAzureStorageBlobContainer),
	// 		ProvisioningState: to.Ptr(armstoragemover.ProvisioningStateSucceeded),
	// 		BlobContainerName: to.Ptr("examples-blobcontainer"),
	// 		StorageAccountResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.Storage/storageAccounts/examplesa"),
	// 	},
	// }
}
