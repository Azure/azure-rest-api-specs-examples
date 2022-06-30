package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/galleryExamples/Gallery_AddToSharingProfile.json
func ExampleGallerySharingProfileClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewGallerySharingProfileClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		armcompute.SharingUpdate{
			Groups: []*armcompute.SharingProfileGroup{
				{
					Type: to.Ptr(armcompute.SharingProfileGroupTypesSubscriptions),
					IDs: []*string{
						to.Ptr("34a4ab42-0d72-47d9-bd1a-aed207386dac"),
						to.Ptr("380fd389-260b-41aa-bad9-0a83108c370b")},
				},
				{
					Type: to.Ptr(armcompute.SharingProfileGroupTypesAADTenants),
					IDs: []*string{
						to.Ptr("c24c76aa-8897-4027-9b03-8f7928b54ff6")},
				}},
			OperationType: to.Ptr(armcompute.SharingUpdateOperationTypesAdd),
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
