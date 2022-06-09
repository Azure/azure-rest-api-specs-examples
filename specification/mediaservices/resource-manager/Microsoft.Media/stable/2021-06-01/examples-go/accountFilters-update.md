```go
package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/accountFilters-update.json
func ExampleAccountFiltersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewAccountFiltersClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<filter-name>",
		armmediaservices.AccountFilter{
			Properties: &armmediaservices.MediaFilterProperties{
				FirstQuality: &armmediaservices.FirstQuality{
					Bitrate: to.Int32Ptr(128000),
				},
				PresentationTimeRange: &armmediaservices.PresentationTimeRange{
					EndTimestamp:               to.Int64Ptr(170000000),
					ForceEndTimestamp:          to.BoolPtr(false),
					LiveBackoffDuration:        to.Int64Ptr(0),
					PresentationWindowDuration: to.Int64Ptr(9223372036854775000),
					StartTimestamp:             to.Int64Ptr(10),
					Timescale:                  to.Int64Ptr(10000000),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountFiltersClientUpdateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmediaservices%2Farmmediaservices%2Fv0.3.1/sdk/resourcemanager/mediaservices/armmediaservices/README.md) on how to add the SDK to your project and authenticate.
