package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry"
)

// Generated from example definition: 2024-11-01/Create_Asset_Without_DisplayName.json
func ExampleAssetsClient_BeginCreateOrReplace_createAssetWithoutDisplayName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAssetsClient().BeginCreateOrReplace(ctx, "myResourceGroup", "my-asset", armdeviceregistry.Asset{
		Location: to.Ptr("West Europe"),
		ExtendedLocation: &armdeviceregistry.ExtendedLocation{
			Type: to.Ptr("CustomLocation"),
			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		},
		Tags: map[string]*string{
			"site": to.Ptr("building-1"),
		},
		Properties: &armdeviceregistry.AssetProperties{
			Enabled:                      to.Ptr(true),
			ExternalAssetID:              to.Ptr("8ZBA6LRHU0A458969"),
			Description:                  to.Ptr("This is a sample Asset"),
			AssetEndpointProfileRef:      to.Ptr("myAssetEndpointProfile"),
			Manufacturer:                 to.Ptr("Contoso"),
			ManufacturerURI:              to.Ptr("https://www.contoso.com/manufacturerUri"),
			Model:                        to.Ptr("ContosoModel"),
			ProductCode:                  to.Ptr("SA34VDG"),
			HardwareRevision:             to.Ptr("1.0"),
			SoftwareRevision:             to.Ptr("2.0"),
			DocumentationURI:             to.Ptr("https://www.example.com/manual"),
			SerialNumber:                 to.Ptr("64-103816-519918-8"),
			DefaultDatasetsConfiguration: to.Ptr("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
			DefaultEventsConfiguration:   to.Ptr("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
			DefaultTopic: &armdeviceregistry.Topic{
				Path:   to.Ptr("/path/defaultTopic"),
				Retain: to.Ptr(armdeviceregistry.TopicRetainTypeKeep),
			},
			Datasets: []*armdeviceregistry.Dataset{
				{
					Name:                 to.Ptr("dataset1"),
					DatasetConfiguration: to.Ptr("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
					Topic: &armdeviceregistry.Topic{
						Path:   to.Ptr("/path/dataset1"),
						Retain: to.Ptr(armdeviceregistry.TopicRetainTypeKeep),
					},
					DataPoints: []*armdeviceregistry.DataPoint{
						{
							Name:                   to.Ptr("dataPoint1"),
							DataSource:             to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt1"),
							ObservabilityMode:      to.Ptr(armdeviceregistry.DataPointObservabilityModeCounter),
							DataPointConfiguration: to.Ptr("{\"publishingInterval\":8,\"samplingInterval\":8,\"queueSize\":4}"),
						},
						{
							Name:                   to.Ptr("dataPoint2"),
							DataSource:             to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt2"),
							ObservabilityMode:      to.Ptr(armdeviceregistry.DataPointObservabilityModeNone),
							DataPointConfiguration: to.Ptr("{\"publishingInterval\":4,\"samplingInterval\":4,\"queueSize\":7}"),
						},
					},
				},
			},
			Events: []*armdeviceregistry.Event{
				{
					Name:               to.Ptr("event1"),
					EventNotifier:      to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt3"),
					ObservabilityMode:  to.Ptr(armdeviceregistry.EventObservabilityModeNone),
					EventConfiguration: to.Ptr("{\"publishingInterval\":7,\"samplingInterval\":1,\"queueSize\":8}"),
					Topic: &armdeviceregistry.Topic{
						Path:   to.Ptr("/path/event1"),
						Retain: to.Ptr(armdeviceregistry.TopicRetainTypeKeep),
					},
				},
				{
					Name:               to.Ptr("event2"),
					EventNotifier:      to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt4"),
					ObservabilityMode:  to.Ptr(armdeviceregistry.EventObservabilityModeLog),
					EventConfiguration: to.Ptr("{\"publishingInterval\":7,\"samplingInterval\":8,\"queueSize\":4}"),
				},
			},
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
	// res = armdeviceregistry.AssetsClientCreateOrReplaceResponse{
	// 	Asset: &armdeviceregistry.Asset{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assets/my-asset"),
	// 		Name: to.Ptr("my-asset"),
	// 		Type: to.Ptr("Microsoft.DeviceRegistry/assets"),
	// 		Location: to.Ptr("West Europe"),
	// 		ExtendedLocation: &armdeviceregistry.ExtendedLocation{
	// 			Type: to.Ptr("CustomLocation"),
	// 			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"site": to.Ptr("building-1"),
	// 		},
	// 		SystemData: &armdeviceregistry.SystemData{
	// 			CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 			CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.2516899Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 			LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.0922793Z"); return t}()),
	// 		},
	// 		Properties: &armdeviceregistry.AssetProperties{
	// 			UUID: to.Ptr("0796f7c1-f2c8-44d7-9f5b-9a6f9522a85d"),
	// 			Enabled: to.Ptr(true),
	// 			ExternalAssetID: to.Ptr("8ZBA6LRHU0A458969"),
	// 			DisplayName: to.Ptr("my-asset"),
	// 			Description: to.Ptr("This is a sample Asset"),
	// 			AssetEndpointProfileRef: to.Ptr("myAssetEndpointProfile"),
	// 			Version: to.Ptr[int64](73766),
	// 			Manufacturer: to.Ptr("Contoso"),
	// 			ManufacturerURI: to.Ptr("https://www.contoso.com/manufacturerUri"),
	// 			Model: to.Ptr("ContosoModel"),
	// 			ProductCode: to.Ptr("SA34VDG"),
	// 			HardwareRevision: to.Ptr("1.0"),
	// 			SoftwareRevision: to.Ptr("2.0"),
	// 			DocumentationURI: to.Ptr("https://www.example.com/manual"),
	// 			SerialNumber: to.Ptr("64-103816-519918-8"),
	// 			DefaultDatasetsConfiguration: to.Ptr("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
	// 			DefaultEventsConfiguration: to.Ptr("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
	// 			DefaultTopic: &armdeviceregistry.Topic{
	// 				Path: to.Ptr("/path/defaultTopic"),
	// 				Retain: to.Ptr(armdeviceregistry.TopicRetainTypeKeep),
	// 			},
	// 			Datasets: []*armdeviceregistry.Dataset{
	// 				{
	// 					Name: to.Ptr("dataset1"),
	// 					DatasetConfiguration: to.Ptr("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
	// 					Topic: &armdeviceregistry.Topic{
	// 						Path: to.Ptr("/path/dataset1"),
	// 						Retain: to.Ptr(armdeviceregistry.TopicRetainTypeKeep),
	// 					},
	// 					DataPoints: []*armdeviceregistry.DataPoint{
	// 						{
	// 							Name: to.Ptr("dataPoint1"),
	// 							DataSource: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt1"),
	// 							ObservabilityMode: to.Ptr(armdeviceregistry.DataPointObservabilityModeCounter),
	// 							DataPointConfiguration: to.Ptr("{\"publishingInterval\":8,\"samplingInterval\":8,\"queueSize\":4}"),
	// 						},
	// 						{
	// 							Name: to.Ptr("dataPoint2"),
	// 							DataSource: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt2"),
	// 							ObservabilityMode: to.Ptr(armdeviceregistry.DataPointObservabilityModeNone),
	// 							DataPointConfiguration: to.Ptr("{\"publishingInterval\":4,\"samplingInterval\":4,\"queueSize\":7}"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Events: []*armdeviceregistry.Event{
	// 				{
	// 					Name: to.Ptr("event1"),
	// 					EventNotifier: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt3"),
	// 					ObservabilityMode: to.Ptr(armdeviceregistry.EventObservabilityModeNone),
	// 					EventConfiguration: to.Ptr("{\"publishingInterval\":7,\"samplingInterval\":1,\"queueSize\":8}"),
	// 					Topic: &armdeviceregistry.Topic{
	// 						Path: to.Ptr("/path/event1"),
	// 						Retain: to.Ptr(armdeviceregistry.TopicRetainTypeKeep),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("event2"),
	// 					EventNotifier: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt4"),
	// 					ObservabilityMode: to.Ptr(armdeviceregistry.EventObservabilityModeLog),
	// 					EventConfiguration: to.Ptr("{\"publishingInterval\":7,\"samplingInterval\":8,\"queueSize\":4}"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
