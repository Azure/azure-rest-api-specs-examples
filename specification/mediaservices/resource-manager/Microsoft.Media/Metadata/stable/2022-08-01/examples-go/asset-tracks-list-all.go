package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/asset-tracks-list-all.json
func ExampleTracksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTracksClient().NewListPager("contoso", "contosomedia", "ClimbingMountRainer", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AssetTrackCollection = armmediaservices.AssetTrackCollection{
		// 	Value: []*armmediaservices.AssetTrack{
		// 		{
		// 			Name: to.Ptr("video"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/assets/tracks"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/assets/ClimbingMountRainer/tracks/video"),
		// 			Properties: &armmediaservices.AssetTrackProperties{
		// 				ProvisioningState: to.Ptr(armmediaservices.ProvisioningStateSucceeded),
		// 				Track: &armmediaservices.VideoTrack{
		// 					ODataType: to.Ptr("#Microsoft.Media.VideoTrack"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("audio"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/assets/tracks"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/assets/ClimbingMountRainer/tracks/audio"),
		// 			Properties: &armmediaservices.AssetTrackProperties{
		// 				ProvisioningState: to.Ptr(armmediaservices.ProvisioningStateSucceeded),
		// 				Track: &armmediaservices.AudioTrack{
		// 					ODataType: to.Ptr("#Microsoft.Media.AudioTrack"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("text1"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/assets/tracks"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/assets/ClimbingMountRainer/tracks/text1"),
		// 			Properties: &armmediaservices.AssetTrackProperties{
		// 				ProvisioningState: to.Ptr(armmediaservices.ProvisioningStateSucceeded),
		// 				Track: &armmediaservices.TextTrack{
		// 					ODataType: to.Ptr("#Microsoft.Media.TextTrack"),
		// 					DisplayName: to.Ptr("Auto generated"),
		// 					FileName: to.Ptr("auto_generated.ttml"),
		// 					LanguageCode: to.Ptr("en-us"),
		// 					PlayerVisibility: to.Ptr(armmediaservices.VisibilityVisible),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("text2"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/assets/tracks"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/assets/ClimbingMountRainer/tracks/text2"),
		// 			Properties: &armmediaservices.AssetTrackProperties{
		// 				ProvisioningState: to.Ptr(armmediaservices.ProvisioningStateSucceeded),
		// 				Track: &armmediaservices.TextTrack{
		// 					ODataType: to.Ptr("#Microsoft.Media.TextTrack"),
		// 					DisplayName: to.Ptr("user uploaded text track"),
		// 					FileName: to.Ptr("text2.vtt"),
		// 					HlsSettings: &armmediaservices.HlsSettings{
		// 						Default: to.Ptr(true),
		// 						Characteristics: to.Ptr("public.accessibility.transcribes-spoken-dialog,public.easy-to-read"),
		// 						Forced: to.Ptr(true),
		// 					},
		// 					LanguageCode: to.Ptr("en-us"),
		// 					PlayerVisibility: to.Ptr(armmediaservices.VisibilityHidden),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
