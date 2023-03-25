package armorbital_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/ContactProfilesByResourceGroupList.json
func ExampleContactProfilesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armorbital.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContactProfilesClient().NewListPager("contoso-Rgp", &armorbital.ContactProfilesClientListOptions{Skiptoken: nil})
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
		// page.ContactProfileListResult = armorbital.ContactProfileListResult{
		// 	Value: []*armorbital.ContactProfile{
		// 		{
		// 			Name: to.Ptr("CONTOSO-CP"),
		// 			Type: to.Ptr("Microsoft.Orbital/contactProfiles"),
		// 			ID: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/contactProfiles/CONTOSO-CP"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armorbital.ContactProfileProperties{
		// 				AutoTrackingConfiguration: to.Ptr(armorbital.AutoTrackingConfigurationDisabled),
		// 				EventHubURI: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.EventHub/namespaces/contosoHub/eventhubs/contosoHub"),
		// 				Links: []*armorbital.ContactProfileLink{
		// 					{
		// 						Name: to.Ptr("contoso-uplink"),
		// 						Channels: []*armorbital.ContactProfileLinkChannel{
		// 							{
		// 								Name: to.Ptr("contoso-uplink-channel"),
		// 								BandwidthMHz: to.Ptr[float32](2),
		// 								CenterFrequencyMHz: to.Ptr[float32](2250),
		// 								EndPoint: &armorbital.EndPoint{
		// 									EndPointName: to.Ptr("ContosoTest_Uplink"),
		// 									IPAddress: to.Ptr("10.1.0.4"),
		// 									Port: to.Ptr("50000"),
		// 									Protocol: to.Ptr(armorbital.ProtocolTCP),
		// 								},
		// 						}},
		// 						Direction: to.Ptr(armorbital.DirectionUplink),
		// 						EirpdBW: to.Ptr[float32](45),
		// 						GainOverTemperature: to.Ptr[float32](0),
		// 						Polarization: to.Ptr(armorbital.PolarizationLHCP),
		// 					},
		// 					{
		// 						Name: to.Ptr("contoso-downlink"),
		// 						Channels: []*armorbital.ContactProfileLinkChannel{
		// 							{
		// 								Name: to.Ptr("contoso-downlink-channel"),
		// 								BandwidthMHz: to.Ptr[float32](15),
		// 								CenterFrequencyMHz: to.Ptr[float32](8160),
		// 								EndPoint: &armorbital.EndPoint{
		// 									EndPointName: to.Ptr("ContosoTest_Downlink"),
		// 									IPAddress: to.Ptr("10.1.0.5"),
		// 									Port: to.Ptr("50001"),
		// 									Protocol: to.Ptr(armorbital.ProtocolUDP),
		// 								},
		// 						}},
		// 						Direction: to.Ptr(armorbital.DirectionDownlink),
		// 						EirpdBW: to.Ptr[float32](0),
		// 						GainOverTemperature: to.Ptr[float32](25),
		// 						Polarization: to.Ptr(armorbital.PolarizationRHCP),
		// 				}},
		// 				MinimumElevationDegrees: to.Ptr[float32](5),
		// 				MinimumViableContactDuration: to.Ptr("PT1M"),
		// 				NetworkConfiguration: &armorbital.ContactProfilesPropertiesNetworkConfiguration{
		// 					SubnetID: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Network/virtualNetworks/contoso-vnet/subnets/orbital-delegated-subnet"),
		// 				},
		// 				ProvisioningState: to.Ptr(armorbital.ContactProfilesPropertiesProvisioningState("Succeeded")),
		// 				ThirdPartyConfigurations: []*armorbital.ContactProfileThirdPartyConfiguration{
		// 					{
		// 						MissionConfiguration: to.Ptr("Ksat_MissionConfiguration"),
		// 						ProviderName: to.Ptr("KSAT"),
		// 					},
		// 					{
		// 						MissionConfiguration: to.Ptr("Viasat_Configuration"),
		// 						ProviderName: to.Ptr("VIASAT"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
