package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/galleryExamples/GalleryImageVersion_Update.json
func ExampleGalleryImageVersionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewGalleryImageVersionsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		"1.0.0",
		armcompute.GalleryImageVersionUpdate{
			Properties: &armcompute.GalleryImageVersionProperties{
				PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
					TargetRegions: []*armcompute.TargetRegion{
						{
							Name:                 to.Ptr("West US"),
							RegionalReplicaCount: to.Ptr[int32](1),
						},
						{
							Name:                 to.Ptr("East US"),
							RegionalReplicaCount: to.Ptr[int32](2),
							StorageAccountType:   to.Ptr(armcompute.StorageAccountTypeStandardZRS),
						}},
				},
				StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
					Source: &armcompute.GalleryArtifactVersionSource{
						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}"),
					},
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
