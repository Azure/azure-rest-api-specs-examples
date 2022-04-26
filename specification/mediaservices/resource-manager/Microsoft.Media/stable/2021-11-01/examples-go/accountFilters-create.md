Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmediaservices%2Farmmediaservices%2Fv0.6.0/sdk/resourcemanager/mediaservices/armmediaservices/README.md) on how to add the SDK to your project and authenticate.

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
		return
	}
	ctx := context.Background()
	client, err := armmediaservices.NewAccountFiltersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<filter-name>",
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
								Value:     to.Ptr("<value>"),
							},
							{
								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationNotEqual),
								Property:  to.Ptr(armmediaservices.FilterTrackPropertyTypeLanguage),
								Value:     to.Ptr("<value>"),
							},
							{
								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationNotEqual),
								Property:  to.Ptr(armmediaservices.FilterTrackPropertyTypeFourCC),
								Value:     to.Ptr("<value>"),
							}},
					},
					{
						TrackSelections: []*armmediaservices.FilterTrackPropertyCondition{
							{
								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationEqual),
								Property:  to.Ptr(armmediaservices.FilterTrackPropertyTypeType),
								Value:     to.Ptr("<value>"),
							},
							{
								Operation: to.Ptr(armmediaservices.FilterTrackPropertyCompareOperationEqual),
								Property:  to.Ptr(armmediaservices.FilterTrackPropertyTypeBitrate),
								Value:     to.Ptr("<value>"),
							}},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
