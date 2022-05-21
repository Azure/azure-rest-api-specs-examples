Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Forbital%2Farmorbital%2Fv1.0.0/sdk/resourcemanager/orbital/armorbital/README.md) on how to add the SDK to your project and authenticate.

```go
package armorbital_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/SpacecraftCreate.json
func ExampleSpacecraftsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armorbital.NewSpacecraftsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"AQUA",
		armorbital.Spacecraft{
			Location: to.Ptr("westus"),
			Properties: &armorbital.SpacecraftsProperties{
				Links: []*armorbital.SpacecraftLink{
					{
						Name:               to.Ptr("S_RHCP_UL"),
						BandwidthMHz:       to.Ptr[float32](0.036),
						CenterFrequencyMHz: to.Ptr[float32](2106.4063),
						Direction:          to.Ptr(armorbital.DirectionUplink),
						Polarization:       to.Ptr(armorbital.PolarizationRHCP),
					},
					{
						Name:               to.Ptr("X_RHCP_DL"),
						BandwidthMHz:       to.Ptr[float32](150),
						CenterFrequencyMHz: to.Ptr[float32](8125),
						Direction:          to.Ptr(armorbital.DirectionDownlink),
						Polarization:       to.Ptr(armorbital.PolarizationRHCP),
					}},
				NoradID:   to.Ptr("27424"),
				TitleLine: to.Ptr("(AQUA)"),
				TleLine1:  to.Ptr("1 27424U 02022A   20195.59202355  .00000039  00000-0  18634-4 0  9991"),
				TleLine2:  to.Ptr("2 27424  98.2098 135.8486 0000176  28.4050 144.5909 14.57108832967671"),
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
```
