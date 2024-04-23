package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c77bbf822be2deaac1b690270c6cd03a52df0e37/specification/deviceregistry/resource-manager/Microsoft.DeviceRegistry/preview/2023-11-01-preview/examples/Create_Asset_Without_DisplayName.json
func ExampleAssetsClient_BeginCreateOrReplace_createAnAssetWithoutDisplayName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAssetsClient().BeginCreateOrReplace(ctx, "myResourceGroup", "my-asset", armdeviceregistry.Asset{
		Location: to.Ptr("West Europe"),
		Tags: map[string]*string{
			"site": to.Ptr("building-1"),
		},
		ExtendedLocation: &armdeviceregistry.ExtendedLocation{
			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
			Type: to.Ptr("CustomLocation"),
		},
		Properties: &armdeviceregistry.AssetProperties{
			Description:             to.Ptr("This is a sample Asset"),
			AssetEndpointProfileURI: to.Ptr("https://www.example.com/myAssetEndpointProfile"),
			AssetType:               to.Ptr("MyAssetType"),
			DataPoints: []*armdeviceregistry.DataPoint{
				{
					CapabilityID:           to.Ptr("dtmi:com:example:Thermostat:__temperature;1"),
					DataPointConfiguration: to.Ptr("{\"publishingInterval\":8,\"samplingInterval\":8,\"queueSize\":4}"),
					DataSource:             to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt1"),
					ObservabilityMode:      to.Ptr(armdeviceregistry.DataPointsObservabilityModeCounter),
				},
				{
					CapabilityID:           to.Ptr("dtmi:com:example:Thermostat:__pressure;1"),
					DataPointConfiguration: to.Ptr("{\"publishingInterval\":4,\"samplingInterval\":4,\"queueSize\":7}"),
					DataSource:             to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt2"),
					ObservabilityMode:      to.Ptr(armdeviceregistry.DataPointsObservabilityModeNone),
				}},
			DefaultDataPointsConfiguration: to.Ptr("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
			DefaultEventsConfiguration:     to.Ptr("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
			DocumentationURI:               to.Ptr("https://www.example.com/manual"),
			Enabled:                        to.Ptr(true),
			Events: []*armdeviceregistry.Event{
				{
					CapabilityID:       to.Ptr("dtmi:com:example:Thermostat:__temperature;1"),
					EventConfiguration: to.Ptr("{\"publishingInterval\":7,\"samplingInterval\":1,\"queueSize\":8}"),
					EventNotifier:      to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt3"),
					ObservabilityMode:  to.Ptr(armdeviceregistry.EventsObservabilityModeNone),
				},
				{
					CapabilityID:       to.Ptr("dtmi:com:example:Thermostat:__pressure;1"),
					EventConfiguration: to.Ptr("{\"publishingInterval\":7,\"samplingInterval\":8,\"queueSize\":4}"),
					EventNotifier:      to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt4"),
					ObservabilityMode:  to.Ptr(armdeviceregistry.EventsObservabilityModeLog),
				}},
			ExternalAssetID:  to.Ptr("8ZBA6LRHU0A458969"),
			HardwareRevision: to.Ptr("1.0"),
			Manufacturer:     to.Ptr("Contoso"),
			ManufacturerURI:  to.Ptr("https://www.contoso.com/manufacturerUri"),
			Model:            to.Ptr("ContosoModel"),
			ProductCode:      to.Ptr("SA34VDG"),
			SerialNumber:     to.Ptr("64-103816-519918-8"),
			SoftwareRevision: to.Ptr("2.0"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Asset = armdeviceregistry.Asset{
	// 	Name: to.Ptr("my-asset"),
	// 	Type: to.Ptr("Microsoft.DeviceRegistry/assets"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assets/my-asset"),
	// 	SystemData: &armdeviceregistry.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.251Z"); return t}()),
	// 		CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 		CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.092Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 		LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("West Europe"),
	// 	Tags: map[string]*string{
	// 		"site": to.Ptr("building-1"),
	// 	},
	// 	ExtendedLocation: &armdeviceregistry.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armdeviceregistry.AssetProperties{
	// 		Description: to.Ptr("This is a sample Asset"),
	// 		AssetEndpointProfileURI: to.Ptr("https://www.example.com/myAssetEndpointProfile"),
	// 		AssetType: to.Ptr("MyAssetType"),
	// 		DataPoints: []*armdeviceregistry.DataPoint{
	// 			{
	// 				CapabilityID: to.Ptr("dtmi:com:example:Thermostat:__temperature;1"),
	// 				DataPointConfiguration: to.Ptr("{\"publishingInterval\":8,\"samplingInterval\":8,\"queueSize\":4}"),
	// 				DataSource: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt1"),
	// 				ObservabilityMode: to.Ptr(armdeviceregistry.DataPointsObservabilityModeCounter),
	// 			},
	// 			{
	// 				CapabilityID: to.Ptr("dtmi:com:example:Thermostat:__pressure;1"),
	// 				DataPointConfiguration: to.Ptr("{\"publishingInterval\":4,\"samplingInterval\":4,\"queueSize\":7}"),
	// 				DataSource: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt2"),
	// 				ObservabilityMode: to.Ptr(armdeviceregistry.DataPointsObservabilityModeNone),
	// 		}},
	// 		DefaultDataPointsConfiguration: to.Ptr("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
	// 		DefaultEventsConfiguration: to.Ptr("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
	// 		DisplayName: to.Ptr("myAsset"),
	// 		DocumentationURI: to.Ptr("https://www.example.com/manual"),
	// 		Enabled: to.Ptr(true),
	// 		Events: []*armdeviceregistry.Event{
	// 			{
	// 				CapabilityID: to.Ptr("dtmi:com:example:Thermostat:__temperature;1"),
	// 				EventConfiguration: to.Ptr("{\"publishingInterval\":7,\"samplingInterval\":1,\"queueSize\":8}"),
	// 				EventNotifier: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt3"),
	// 				ObservabilityMode: to.Ptr(armdeviceregistry.EventsObservabilityModeNone),
	// 			},
	// 			{
	// 				CapabilityID: to.Ptr("dtmi:com:example:Thermostat:__pressure;1"),
	// 				EventConfiguration: to.Ptr("{\"publishingInterval\":7,\"samplingInterval\":8,\"queueSize\":4}"),
	// 				EventNotifier: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt4"),
	// 				ObservabilityMode: to.Ptr(armdeviceregistry.EventsObservabilityModeLog),
	// 		}},
	// 		ExternalAssetID: to.Ptr("8ZBA6LRHU0A458969"),
	// 		HardwareRevision: to.Ptr("1.0"),
	// 		Manufacturer: to.Ptr("Contoso"),
	// 		ManufacturerURI: to.Ptr("https://www.contoso.com/manufacturerUri"),
	// 		Model: to.Ptr("ContosoModel"),
	// 		ProductCode: to.Ptr("SA34VDG"),
	// 		ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		SerialNumber: to.Ptr("64-103816-519918-8"),
	// 		SoftwareRevision: to.Ptr("2.0"),
	// 		UUID: to.Ptr("0796f7c1-f2c8-44d7-9f5b-9a6f9522a85d"),
	// 		Version: to.Ptr[int32](73766),
	// 	},
	// }
}
