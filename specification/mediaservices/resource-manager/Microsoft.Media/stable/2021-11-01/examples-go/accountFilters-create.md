```go
package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/accountFilters-create.json
func ExampleAccountFiltersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewAccountFiltersClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"contoso",
		"contosomedia",
		"newAccountFilter",
		armmediaservices.AccountFilter{
			Properties: &armmediaservices.MediaFilterProperties{
				FirstQuality: &armmediaservices.FirstQuality{
					Bitrate: to.Ptr[int32](128000),
				},
				PresentationTimeRange: &armmediaservices.PresentationTimeRange{
					EndTimestamp:               to.Ptr[int64](170000000),
					ForceEndTimestamp:          to.Ptr(false),
					LiveBackoffDuration:        to.Ptr[int64](0),
					PresentationWindowDuration: to.Ptr[int64](9223372036854775000),
					StartTimestamp:             to.Ptr[int64](0),
					Timescale:                  to.Ptr[int64](10000000),
				},
				Tracks: []*armmediaservices.FilterTrackSelection{
					{
						TrackSelections: []*armmediaservices.FilterTrackPropertyCondition{
							{
								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationEqual),
								Property:  to.Ptr(armmediaservices.FilterTrackPropertyTypeType),
								Value:     to.Ptr("Audio"),
							},
							{
								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationNotEqual),
								Property:  to.Ptr(armmediaservices.FilterTrackPropertyTypeLanguage),
								Value:     to.Ptr("en"),
							},
							{
								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationNotEqual),
								Property:  to.Ptr(armmediaservices.FilterTrackPropertyTypeFourCC),
								Value:     to.Ptr("EC-3"),
							}},
					},
					{
						TrackSelections: []*armmediaservices.FilterTrackPropertyCondition{
							{
								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationEqual),
								Property:  to.Ptr(armmediaservices.FilterTrackPropertyTypeType),
								Value:     to.Ptr("Video"),
							},
							{
								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationEqual),
								Property:  to.Ptr(armmediaservices.FilterTrackPropertyTypeBitrate),
								Value:     to.Ptr("3000000-5000000"),
							}},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmediaservices%2Farmmediaservices%2Fv1.0.0/sdk/resourcemanager/mediaservices/armmediaservices/README.md) on how to add the SDK to your project and authenticate.
