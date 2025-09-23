package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/260ed6a52537921f53a18ffaf4020e3b4d510367/specification/storage/resource-manager/Microsoft.Storage/stable/2025-01-01/examples/BlobContainersPatch.json
func ExampleBlobContainersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBlobContainersClient().Update(ctx, "res3376", "sto328", "container6185", armstorage.BlobContainer{
		ContainerProperties: &armstorage.ContainerProperties{
			Metadata: map[string]*string{
				"metadata": to.Ptr("true"),
			},
			PublicAccess: to.Ptr(armstorage.PublicAccessContainer),
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
	// 		HasImmutabilityPolicy: to.Ptr(false),
	// 		HasLegalHold: to.Ptr(false),
	// 		Metadata: map[string]*string{
	// 			"metadata": to.Ptr("true"),
	// 		},
	// 		PublicAccess: to.Ptr(armstorage.PublicAccessContainer),
	// 	},
	// }
}
