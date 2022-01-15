Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Forbital%2Farmorbital%2Fv0.2.0/sdk/resourcemanager/orbital/armorbital/README.md) on how to add the SDK to your project and authenticate.

```go
package armorbital_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital"
)

// x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/preview/2021-04-04-preview/examples/SpacecraftCreate.json
func ExampleSpacecraftsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armorbital.NewSpacecraftsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<spacecraft-name>",
		armorbital.Spacecraft{
			Location: to.StringPtr("<location>"),
			Properties: &armorbital.SpacecraftsProperties{
				Links: []*armorbital.SpacecraftLink{
					{
						BandwidthMHz:       to.Float32Ptr(0.036),
						CenterFrequencyMHz: to.Float32Ptr(2106.4063),
						Direction:          armorbital.Direction("uplink").ToPtr(),
						Polarization:       armorbital.Polarization("RHCP").ToPtr(),
					},
					{
						BandwidthMHz:       to.Float32Ptr(150),
						CenterFrequencyMHz: to.Float32Ptr(8125),
						Direction:          armorbital.Direction("downlink").ToPtr(),
						Polarization:       armorbital.Polarization("RHCP").ToPtr(),
					}},
				NoradID:   to.StringPtr("<norad-id>"),
				TitleLine: to.StringPtr("<title-line>"),
				TleLine1:  to.StringPtr("<tle-line1>"),
				TleLine2:  to.StringPtr("<tle-line2>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SpacecraftsClientCreateOrUpdateResult)
}
```
