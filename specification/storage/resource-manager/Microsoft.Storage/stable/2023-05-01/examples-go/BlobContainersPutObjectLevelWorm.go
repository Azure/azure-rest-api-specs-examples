package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/220ad9c6554fc7d6d10a89bdb441c1e3b36e3285/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/BlobContainersPutObjectLevelWorm.json
func ExampleBlobContainersClient_Create_putContainerWithObjectLevelWorm() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBlobContainersClient().Create(ctx, "res3376", "sto328", "container6185", armstorage.BlobContainer{
		ContainerProperties: &armstorage.ContainerProperties{
			ImmutableStorageWithVersioning: &armstorage.ImmutableStorageWithVersioning{
				Enabled: to.Ptr(true),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BlobContainer = armstorage.BlobContainer{
	// 	Name: to.Ptr("container6185"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/blobServices/containers"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/blobServices/default/containers/container6185"),
	// 	ContainerProperties: &armstorage.ContainerProperties{
	// 		ImmutableStorageWithVersioning: &armstorage.ImmutableStorageWithVersioning{
	// 			Enabled: to.Ptr(true),
	// 		},
	// 	},
	// }
}
