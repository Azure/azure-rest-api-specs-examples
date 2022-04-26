Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Forbital%2Farmorbital%2Fv0.4.0/sdk/resourcemanager/orbital/armorbital/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/orbital/resource-manager/Microsoft.Orbital/preview/2021-04-04-preview/examples/ContactProfileCreate.json
func ExampleContactProfilesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armorbital.NewContactProfilesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<contact-profile-name>",
		armorbital.ContactProfile{
			Location: to.Ptr("<location>"),
			Properties: &armorbital.ContactProfilesProperties{
				AutoTrackingConfiguration: to.Ptr(armorbital.AutoTrackingConfigurationXBand),
				EventHubURI:               to.Ptr("<event-hub-uri>"),
				Links: []*armorbital.ContactProfileLink{
					{
						Channels: []*armorbital.ContactProfileLinkChannel{
							{
								BandwidthMHz:              to.Ptr[float32](0.036),
								CenterFrequencyMHz:        to.Ptr[float32](2106.4063),
								DecodingConfiguration:     to.Ptr("<decoding-configuration>"),
								DemodulationConfiguration: to.Ptr("<demodulation-configuration>"),
								EncodingConfiguration:     to.Ptr("<encoding-configuration>"),
								EndPoint: &armorbital.EndPoint{
									EndPointName: to.Ptr("<end-point-name>"),
									IPAddress:    to.Ptr("<ipaddress>"),
									Port:         to.Ptr("<port>"),
									Protocol:     to.Ptr(armorbital.ProtocolTCP),
								},
								ModulationConfiguration: to.Ptr("<modulation-configuration>"),
							}},
						Direction:           to.Ptr(armorbital.DirectionUplink),
						EirpdBW:             to.Ptr[float32](45),
						GainOverTemperature: to.Ptr[float32](0),
						Polarization:        to.Ptr(armorbital.PolarizationRHCP),
					},
					{
						Channels: []*armorbital.ContactProfileLinkChannel{
							{
								BandwidthMHz:              to.Ptr[float32](150),
								CenterFrequencyMHz:        to.Ptr[float32](8160),
								DecodingConfiguration:     to.Ptr("<decoding-configuration>"),
								DemodulationConfiguration: to.Ptr("<demodulation-configuration>"),
								EncodingConfiguration:     to.Ptr("<encoding-configuration>"),
								EndPoint: &armorbital.EndPoint{
									EndPointName: to.Ptr("<end-point-name>"),
									IPAddress:    to.Ptr("<ipaddress>"),
									Port:         to.Ptr("<port>"),
									Protocol:     to.Ptr(armorbital.ProtocolTCP),
								},
								ModulationConfiguration: to.Ptr("<modulation-configuration>"),
							}},
						Direction:           to.Ptr(armorbital.DirectionDownlink),
						EirpdBW:             to.Ptr[float32](0),
						GainOverTemperature: to.Ptr[float32](25),
						Polarization:        to.Ptr(armorbital.PolarizationRHCP),
					}},
				MinimumElevationDegrees:      to.Ptr[float32](10),
				MinimumViableContactDuration: to.Ptr("<minimum-viable-contact-duration>"),
			},
		},
		&armorbital.ContactProfilesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
