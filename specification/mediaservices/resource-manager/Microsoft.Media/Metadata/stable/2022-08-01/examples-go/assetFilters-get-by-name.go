package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assetFilters-get-by-name.json
func ExampleAssetFiltersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssetFiltersClient().Get(ctx, "contoso", "contosomedia", "ClimbingMountRainer", "assetFilterWithTimeWindowAndTrack", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssetFilter = armmediaservices.AssetFilter{
	// 	Name: to.Ptr("assetFilterWithTimeWindowAndTrack"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/assets/assetFilters"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/assets/ClimbingMountRainer/assetFilters/assetFilterWithTimeWindowAndTrack"),
	// 	Properties: &armmediaservices.MediaFilterProperties{
	// 		FirstQuality: &armmediaservices.FirstQuality{
	// 			Bitrate: to.Ptr[int32](128000),
	// 		},
	// 		PresentationTimeRange: &armmediaservices.PresentationTimeRange{
	// 			EndTimestamp: to.Ptr[int64](170000000),
	// 			ForceEndTimestamp: to.Ptr(false),
	// 			LiveBackoffDuration: to.Ptr[int64](0),
	// 			PresentationWindowDuration: to.Ptr[int64](9223372036854775000),
	// 			StartTimestamp: to.Ptr[int64](0),
	// 			Timescale: to.Ptr[int64](10000000),
	// 		},
	// 		Tracks: []*armmediaservices.FilterTrackSelection{
	// 			{
	// 				TrackSelections: []*armmediaservices.FilterTrackPropertyCondition{
	// 					{
	// 						Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationEqual),
	// 						Property: to.Ptr(armmediaservices.FilterTrackPropertyTypeType),
	// 						Value: to.Ptr("Audio"),
	// 					},
	// 					{
	// 						Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationNotEqual),
	// 						Property: to.Ptr(armmediaservices.FilterTrackPropertyTypeLanguage),
	// 						Value: to.Ptr("en"),
	// 					},
	// 					{
	// 						Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationNotEqual),
	// 						Property: to.Ptr(armmediaservices.FilterTrackPropertyTypeFourCC),
	// 						Value: to.Ptr("EC-3"),
	// 				}},
	// 			},
	// 			{
	// 				TrackSelections: []*armmediaservices.FilterTrackPropertyCondition{
	// 					{
	// 						Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationEqual),
	// 						Property: to.Ptr(armmediaservices.FilterTrackPropertyTypeType),
	// 						Value: to.Ptr("Video"),
	// 					},
	// 					{
	// 						Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationEqual),
	// 						Property: to.Ptr(armmediaservices.FilterTrackPropertyTypeBitrate),
	// 						Value: to.Ptr("3000000-5000000"),
	// 				}},
	// 		}},
	// 	},
	// }
}
