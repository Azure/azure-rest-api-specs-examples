package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86c6306649b02e542117adb46c61e8019dbd78e9/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/BlobServicesPutAllowPermanentDelete.json
func ExampleBlobServicesClient_SetServiceProperties_blobServicesPutAllowPermanentDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBlobServicesClient().SetServiceProperties(ctx, "res4410", "sto8607", armstorage.BlobServiceProperties{
		BlobServiceProperties: &armstorage.BlobServicePropertiesProperties{
			DeleteRetentionPolicy: &armstorage.DeleteRetentionPolicy{
				AllowPermanentDelete: to.Ptr(true),
				Days:                 to.Ptr[int32](300),
				Enabled:              to.Ptr(true),
			},
			IsVersioningEnabled: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BlobServiceProperties = armstorage.BlobServiceProperties{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/blobServices"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res4410/providers/Microsoft.Storage/storageAccounts/sto8607/blobServices/default"),
	// 	BlobServiceProperties: &armstorage.BlobServicePropertiesProperties{
	// 		DeleteRetentionPolicy: &armstorage.DeleteRetentionPolicy{
	// 			AllowPermanentDelete: to.Ptr(true),
	// 			Days: to.Ptr[int32](300),
	// 			Enabled: to.Ptr(true),
	// 		},
	// 		IsVersioningEnabled: to.Ptr(true),
	// 	},
	// }
}
