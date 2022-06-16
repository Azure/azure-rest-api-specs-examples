package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/galleryExamples/GalleryApplicationVersion_Update.json
func ExampleGalleryApplicationVersionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewGalleryApplicationVersionsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryApplicationName",
		"1.0.0",
		armcompute.GalleryApplicationVersionUpdate{
			Properties: &armcompute.GalleryApplicationVersionProperties{
				PublishingProfile: &armcompute.GalleryApplicationVersionPublishingProfile{
					EndOfLifeDate:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-01T07:00:00Z"); return t }()),
					ReplicaCount:       to.Ptr[int32](1),
					StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
					TargetRegions: []*armcompute.TargetRegion{
						{
							Name:                 to.Ptr("West US"),
							RegionalReplicaCount: to.Ptr[int32](1),
							StorageAccountType:   to.Ptr(armcompute.StorageAccountTypeStandardLRS),
						}},
					ManageActions: &armcompute.UserArtifactManage{
						Install: to.Ptr("powershell -command \"Expand-Archive -Path package.zip -DestinationPath C:\\package\""),
						Remove:  to.Ptr("del C:\\package "),
					},
					Source: &armcompute.UserArtifactSource{
						MediaLink: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/package.zip?{sasKey}"),
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
