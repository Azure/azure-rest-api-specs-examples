package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/galleryExamples/CommunityGallery_Create.json
func ExampleGalleriesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewGalleriesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		armcompute.Gallery{
			Location: to.Ptr("West US"),
			Properties: &armcompute.GalleryProperties{
				Description: to.Ptr("This is the gallery description."),
				SharingProfile: &armcompute.SharingProfile{
					CommunityGalleryInfo: map[string]interface{}{
						"eula":             "eula",
						"publicNamePrefix": "PirPublic",
						"publisherContact": "pir@microsoft.com",
						"publisherUri":     "uri",
					},
					Permissions: to.Ptr(armcompute.GallerySharingPermissionTypesCommunity),
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
