Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Forbital%2Farmorbital%2Fv0.1.0/sdk/resourcemanager/orbital/armorbital/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/preview/2021-04-04-preview/examples/ContactProfileCreate.json
func ExampleContactProfilesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armorbital.NewContactProfilesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<contact-profile-name>",
		armorbital.ContactProfile{
			TrackedResource: armorbital.TrackedResource{
				Location: to.StringPtr("<location>"),
			},
			Properties: &armorbital.ContactProfilesProperties{
				AutoTrackingConfiguration: armorbital.AutoTrackingConfigurationXBand.ToPtr(),
				Links: []*armorbital.ContactProfileLink{
					{
						Channels: []*armorbital.ContactProfileLinkChannel{
							{
								BandwidthMHz:              to.Float32Ptr(0.036),
								CenterFrequencyMHz:        to.Float32Ptr(2106.4063),
								DecodingConfiguration:     to.StringPtr("<decoding-configuration>"),
								DemodulationConfiguration: to.StringPtr("<demodulation-configuration>"),
								EncodingConfiguration:     to.StringPtr("<encoding-configuration>"),
								EndPoint: &armorbital.EndPoint{
									EndPointName: to.StringPtr("<end-point-name>"),
									IPAddress:    to.StringPtr("<ipaddress>"),
									Port:         to.StringPtr("<port>"),
									Protocol:     armorbital.ProtocolTCP.ToPtr(),
								},
								ModulationConfiguration: to.StringPtr("<modulation-configuration>"),
							}},
						Direction:           armorbital.DirectionUplink.ToPtr(),
						EirpdBW:             to.Float32Ptr(45),
						GainOverTemperature: to.Float32Ptr(0),
						Polarization:        armorbital.PolarizationRHCP.ToPtr(),
					},
					{
						Channels: []*armorbital.ContactProfileLinkChannel{
							{
								BandwidthMHz:              to.Float32Ptr(150),
								CenterFrequencyMHz:        to.Float32Ptr(8160),
								DecodingConfiguration:     to.StringPtr("<decoding-configuration>"),
								DemodulationConfiguration: to.StringPtr("<demodulation-configuration>"),
								EncodingConfiguration:     to.StringPtr("<encoding-configuration>"),
								EndPoint: &armorbital.EndPoint{
									EndPointName: to.StringPtr("<end-point-name>"),
									IPAddress:    to.StringPtr("<ipaddress>"),
									Port:         to.StringPtr("<port>"),
									Protocol:     armorbital.ProtocolTCP.ToPtr(),
								},
								ModulationConfiguration: to.StringPtr("<modulation-configuration>"),
							}},
						Direction:           armorbital.DirectionDownlink.ToPtr(),
						EirpdBW:             to.Float32Ptr(0),
						GainOverTemperature: to.Float32Ptr(25),
						Polarization:        armorbital.PolarizationRHCP.ToPtr(),
					}},
				MinimumElevationDegrees:      to.Float32Ptr(10),
				MinimumViableContactDuration: to.StringPtr("<minimum-viable-contact-duration>"),
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
	log.Printf("ContactProfile.ID: %s\n", *res.ID)
}
```
