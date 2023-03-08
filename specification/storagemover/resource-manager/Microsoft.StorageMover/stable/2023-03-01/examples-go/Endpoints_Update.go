package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/241e964afe675a7be98aa6a2e171a3c5f830816c/specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-03-01/examples/Endpoints_Update.json
func ExampleEndpointsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragemover.NewEndpointsClient("11111111-2222-3333-4444-555555555555", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "examples-rg", "examples-storageMoverName", "examples-endpointName", armstoragemover.EndpointBaseUpdateParameters{
		Properties: &armstoragemover.EndpointBaseUpdateProperties{
			Description: to.Ptr("Updated Endpoint Description"),
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
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-endpointName"),
	// 	Properties: &armstoragemover.AzureStorageBlobContainerEndpointProperties{
	// 		Description: to.Ptr("Updated Endpoint Description"),
	// 		EndpointType: to.Ptr(armstoragemover.EndpointTypeAzureStorageBlobContainer),
	// 		BlobContainerName: to.Ptr("examples-blobContainerName"),
	// 		StorageAccountResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.Storage/storageAccounts/examples-storageAccountName/"),
	// 	},
	// }
}
