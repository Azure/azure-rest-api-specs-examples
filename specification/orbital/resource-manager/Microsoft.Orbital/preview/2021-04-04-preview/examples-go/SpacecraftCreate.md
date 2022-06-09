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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/orbital/resource-manager/Microsoft.Orbital/preview/2021-04-04-preview/examples/SpacecraftCreate.json
func ExampleSpacecraftsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armorbital.NewSpacecraftsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<spacecraft-name>",
		armorbital.Spacecraft{
			Location: to.Ptr("<location>"),
			Properties: &armorbital.SpacecraftsProperties{
				Links: []*armorbital.SpacecraftLink{
					{
						BandwidthMHz:       to.Ptr[float32](0.036),
						CenterFrequencyMHz: to.Ptr[float32](2106.4063),
						Direction:          to.Ptr(armorbital.DirectionUplink),
						Polarization:       to.Ptr(armorbital.PolarizationRHCP),
					},
					{
						BandwidthMHz:       to.Ptr[float32](150),
						CenterFrequencyMHz: to.Ptr[float32](8125),
						Direction:          to.Ptr(armorbital.DirectionDownlink),
						Polarization:       to.Ptr(armorbital.PolarizationRHCP),
					}},
				NoradID:   to.Ptr("<norad-id>"),
				TitleLine: to.Ptr("<title-line>"),
				TleLine1:  to.Ptr("<tle-line1>"),
				TleLine2:  to.Ptr("<tle-line2>"),
			},
		},
		&armorbital.SpacecraftsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Forbital%2Farmorbital%2Fv0.4.0/sdk/resourcemanager/orbital/armorbital/README.md) on how to add the SDK to your project and authenticate.
