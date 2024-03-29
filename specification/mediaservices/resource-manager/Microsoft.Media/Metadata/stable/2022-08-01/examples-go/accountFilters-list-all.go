package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/accountFilters-list-all.json
func ExampleAccountFiltersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountFiltersClient().NewListPager("contoso", "contosomedia", nil)
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
		// page.AccountFilterCollection = armmediaservices.AccountFilterCollection{
		// 	Value: []*armmediaservices.AccountFilter{
		// 		{
		// 			Name: to.Ptr("accountFilterWithTimeWindowAndTrack"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/accountFilters"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/accountFilters/accountFilterWithTimeWindowAndTrack"),
		// 			Properties: &armmediaservices.MediaFilterProperties{
		// 				FirstQuality: &armmediaservices.FirstQuality{
		// 					Bitrate: to.Ptr[int32](128000),
		// 				},
		// 				PresentationTimeRange: &armmediaservices.PresentationTimeRange{
		// 					EndTimestamp: to.Ptr[int64](170000000),
		// 					ForceEndTimestamp: to.Ptr(false),
		// 					LiveBackoffDuration: to.Ptr[int64](0),
		// 					PresentationWindowDuration: to.Ptr[int64](9223372036854775000),
		// 					StartTimestamp: to.Ptr[int64](0),
		// 					Timescale: to.Ptr[int64](10000000),
		// 				},
		// 				Tracks: []*armmediaservices.FilterTrackSelection{
		// 					{
		// 						TrackSelections: []*armmediaservices.FilterTrackPropertyCondition{
		// 							{
		// 								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationEqual),
		// 								Property: to.Ptr(armmediaservices.FilterTrackPropertyTypeType),
		// 								Value: to.Ptr("Audio"),
		// 							},
		// 							{
		// 								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationNotEqual),
		// 								Property: to.Ptr(armmediaservices.FilterTrackPropertyTypeLanguage),
		// 								Value: to.Ptr("en"),
		// 							},
		// 							{
		// 								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationNotEqual),
		// 								Property: to.Ptr(armmediaservices.FilterTrackPropertyTypeFourCC),
		// 								Value: to.Ptr("EC-3"),
		// 						}},
		// 					},
		// 					{
		// 						TrackSelections: []*armmediaservices.FilterTrackPropertyCondition{
		// 							{
		// 								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationEqual),
		// 								Property: to.Ptr(armmediaservices.FilterTrackPropertyTypeType),
		// 								Value: to.Ptr("Video"),
		// 							},
		// 							{
		// 								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationEqual),
		// 								Property: to.Ptr(armmediaservices.FilterTrackPropertyTypeBitrate),
		// 								Value: to.Ptr("3000000-5000000"),
		// 						}},
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("accountFilterWithTrack"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/accountFilters"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/accountFilters/accountFilterWithTrack"),
		// 			Properties: &armmediaservices.MediaFilterProperties{
		// 				Tracks: []*armmediaservices.FilterTrackSelection{
		// 					{
		// 						TrackSelections: []*armmediaservices.FilterTrackPropertyCondition{
		// 							{
		// 								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationEqual),
		// 								Property: to.Ptr(armmediaservices.FilterTrackPropertyTypeType),
		// 								Value: to.Ptr("Audio"),
		// 							},
		// 							{
		// 								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationNotEqual),
		// 								Property: to.Ptr(armmediaservices.FilterTrackPropertyTypeLanguage),
		// 								Value: to.Ptr("en"),
		// 							},
		// 							{
		// 								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationNotEqual),
		// 								Property: to.Ptr(armmediaservices.FilterTrackPropertyTypeFourCC),
		// 								Value: to.Ptr("EC-3"),
		// 						}},
		// 					},
		// 					{
		// 						TrackSelections: []*armmediaservices.FilterTrackPropertyCondition{
		// 							{
		// 								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationEqual),
		// 								Property: to.Ptr(armmediaservices.FilterTrackPropertyTypeType),
		// 								Value: to.Ptr("Video"),
		// 							},
		// 							{
		// 								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationEqual),
		// 								Property: to.Ptr(armmediaservices.FilterTrackPropertyTypeBitrate),
		// 								Value: to.Ptr("3000000-5000000"),
		// 						}},
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
