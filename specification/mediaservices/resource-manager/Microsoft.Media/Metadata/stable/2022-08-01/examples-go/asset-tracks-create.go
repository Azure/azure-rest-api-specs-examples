package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/asset-tracks-create.json
func ExampleTracksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTracksClient().BeginCreateOrUpdate(ctx, "contoso", "contosomedia", "ClimbingMountRainer", "text3", armmediaservices.AssetTrack{
		Properties: &armmediaservices.AssetTrackProperties{
			Track: &armmediaservices.TextTrack{
				ODataType:        to.Ptr("#Microsoft.Media.TextTrack"),
				DisplayName:      to.Ptr("A new track"),
				FileName:         to.Ptr("text3.ttml"),
				PlayerVisibility: to.Ptr(armmediaservices.VisibilityVisible),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssetTrack = armmediaservices.AssetTrack{
	// 	Name: to.Ptr("text3"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/assets/tracks"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/assets/ClimbingMountRainer/tracks/text3"),
	// 	Properties: &armmediaservices.AssetTrackProperties{
	// 		ProvisioningState: to.Ptr(armmediaservices.ProvisioningStateSucceeded),
	// 		Track: &armmediaservices.TextTrack{
	// 			ODataType: to.Ptr("#Microsoft.Media.TextTrack"),
	// 			DisplayName: to.Ptr("A new track"),
	// 			FileName: to.Ptr("text3.ttml"),
	// 			PlayerVisibility: to.Ptr(armmediaservices.VisibilityVisible),
	// 		},
	// 	},
	// }
}
