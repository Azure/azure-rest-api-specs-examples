package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobServicesPutLastAccessTimeBasedTracking.json
func ExampleBlobServicesClient_SetServiceProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobServicesClient("<subscription-id>", cred, nil)
	res, err := client.SetServiceProperties(ctx,
		"<resource-group-name>",
		"<account-name>",
		armstorage.BlobServiceProperties{
			BlobServiceProperties: &armstorage.BlobServicePropertiesProperties{
				LastAccessTimeTrackingPolicy: &armstorage.LastAccessTimeTrackingPolicy{
					Name: armstorage.Name("AccessTimeTracking").ToPtr(),
					BlobType: []*string{
						to.StringPtr("blockBlob")},
					Enable:                    to.BoolPtr(true),
					TrackingGranularityInDays: to.Int32Ptr(1),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BlobServicesClientSetServicePropertiesResult)
}
