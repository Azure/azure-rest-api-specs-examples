package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/BlobServicesPutLastAccessTimeBasedTracking.json
func ExampleBlobServicesClient_SetServiceProperties_blobServicesPutLastAccessTimeBasedTracking() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobServicesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.SetServiceProperties(ctx, "res4410", "sto8607", armstorage.BlobServiceProperties{
		BlobServiceProperties: &armstorage.BlobServicePropertiesProperties{
			LastAccessTimeTrackingPolicy: &armstorage.LastAccessTimeTrackingPolicy{
				Name: to.Ptr(armstorage.NameAccessTimeTracking),
				BlobType: []*string{
					to.Ptr("blockBlob")},
				Enable:                    to.Ptr(true),
				TrackingGranularityInDays: to.Ptr[int32](1),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
