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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactProfileCreate.json
func ExampleContactProfilesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armorbital.NewContactProfilesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"AQUA_DIRECTPLAYBACK_WITH_UPLINK",
		armorbital.ContactProfile{
			Location: to.Ptr("westus"),
			Properties: &armorbital.ContactProfileProperties{
				AutoTrackingConfiguration: to.Ptr(armorbital.AutoTrackingConfigurationXBand),
				EventHubURI:               to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.EventHub/namespaces/orbitalppewestus2-ns/eventhubs/telemetry-hub/"),
				Links: []*armorbital.ContactProfileLink{
					{
						Name: to.Ptr("RHCP_UL"),
						Channels: []*armorbital.ContactProfileLinkChannel{
							{
								Name:                      to.Ptr("channel1"),
								BandwidthMHz:              to.Ptr[float32](0.036),
								CenterFrequencyMHz:        to.Ptr[float32](2106.4063),
								DecodingConfiguration:     to.Ptr("na"),
								DemodulationConfiguration: to.Ptr("na"),
								EncodingConfiguration:     to.Ptr("AQUA_CMD_CCSDS"),
								EndPoint: &armorbital.EndPoint{
									EndPointName: to.Ptr("AQUA_command"),
									IPAddress:    to.Ptr("10.0.1.0"),
									Port:         to.Ptr("4000"),
									Protocol:     to.Ptr(armorbital.ProtocolTCP),
								},
								ModulationConfiguration: to.Ptr("AQUA_UPLINK_BPSK"),
							}},
						Direction:           to.Ptr(armorbital.DirectionUplink),
						EirpdBW:             to.Ptr[float32](45),
						GainOverTemperature: to.Ptr[float32](0),
						Polarization:        to.Ptr(armorbital.PolarizationRHCP),
					},
					{
						Name: to.Ptr("RHCP_DL"),
						Channels: []*armorbital.ContactProfileLinkChannel{
							{
								Name:                      to.Ptr("channel1"),
								BandwidthMHz:              to.Ptr[float32](150),
								CenterFrequencyMHz:        to.Ptr[float32](8160),
								DecodingConfiguration:     to.Ptr("AQUA_DIRECTPLAYBACK_CCSDS"),
								DemodulationConfiguration: to.Ptr("AQUA_DOWNLINK_QPSK"),
								EncodingConfiguration:     to.Ptr("na"),
								EndPoint: &armorbital.EndPoint{
									EndPointName: to.Ptr("AQUA_directplayback"),
									IPAddress:    to.Ptr("10.0.2.0"),
									Port:         to.Ptr("4000"),
									Protocol:     to.Ptr(armorbital.ProtocolTCP),
								},
								ModulationConfiguration: to.Ptr("na"),
							}},
						Direction:           to.Ptr(armorbital.DirectionDownlink),
						EirpdBW:             to.Ptr[float32](0),
						GainOverTemperature: to.Ptr[float32](25),
						Polarization:        to.Ptr(armorbital.PolarizationRHCP),
					}},
				MinimumElevationDegrees:      to.Ptr[float32](10),
				MinimumViableContactDuration: to.Ptr("PT1M"),
				NetworkConfiguration: &armorbital.ContactProfilesPropertiesNetworkConfiguration{
					SubnetID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnetName"),
				},
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
