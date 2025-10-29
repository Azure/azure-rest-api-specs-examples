package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry/v2"
)

// Generated from example definition: 2025-10-01/Get_NamespaceDiscoveredAsset.json
func ExampleNamespaceDiscoveredAssetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespaceDiscoveredAssetsClient().Get(ctx, "myResourceGroup", "my-namespace-1", "my-discoveredasset-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdeviceregistry.NamespaceDiscoveredAssetsClientGetResponse{
	// 	NamespaceDiscoveredAsset: &armdeviceregistry.NamespaceDiscoveredAsset{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/namespaces/my-namespace-1/discoveredAssets/my-discoveredasset-1"),
	// 		Name: to.Ptr("my-discoveredasset-1"),
	// 		Type: to.Ptr("Microsoft.DeviceRegistry/namespaces/discoveredAssets"),
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
	// 		Properties: &armdeviceregistry.NamespaceDiscoveredAssetProperties{
	// 			DeviceRef: &armdeviceregistry.DeviceRef{
	// 				DeviceName: to.Ptr("myDevice"),
	// 				EndpointName: to.Ptr("opcuaendpointname"),
	// 			},
	// 			AssetTypeRefs: []*string{
	// 				to.Ptr("myAssetTypeRef1"),
	// 				to.Ptr("myAssetTypeRef2"),
	// 			},
	// 			DiscoveryID: to.Ptr("11111111-1111-1111-1111-111111111111"),
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
	// 			DefaultManagementGroupsConfiguration: to.Ptr("{\"retryCount\":10,\"retryBackoffInterval\":15}"),
	// 			Datasets: []*armdeviceregistry.NamespaceDiscoveredDataset{
	// 				{
	// 					Name: to.Ptr("dataset1"),
	// 					DataSource: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/Boiler;i=5"),
	// 					TypeRef: to.Ptr("dataset1TypeRef"),
	// 					DatasetConfiguration: to.Ptr("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
	// 					LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-09T14:20:00.52Z"); return t}()),
	// 					DataPoints: []*armdeviceregistry.NamespaceDiscoveredDatasetDataPoint{
	// 						{
	// 							Name: to.Ptr("dataPoint1"),
	// 							DataSource: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt1"),
	// 							TypeRef: to.Ptr("dataset1DataPoint1TypeRef"),
	// 							LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-09T14:20:00.52Z"); return t}()),
	// 						},
	// 						{
	// 							Name: to.Ptr("dataPoint2"),
	// 							DataSource: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt2"),
	// 							LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-09T14:20:00.52Z"); return t}()),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			EventGroups: []*armdeviceregistry.NamespaceDiscoveredEventGroup{
	// 				{
	// 					Name: to.Ptr("default"),
	// 					Events: []*armdeviceregistry.NamespaceDiscoveredEvent{
	// 						{
	// 							Name: to.Ptr("event1"),
	// 							DataSource: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt3"),
	// 							LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-09T14:20:00.52Z"); return t}()),
	// 						},
	// 						{
	// 							Name: to.Ptr("event2"),
	// 							DataSource: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt5"),
	// 							LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-09T14:20:00.52Z"); return t}()),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Streams: []*armdeviceregistry.NamespaceDiscoveredStream{
	// 				{
	// 					Name: to.Ptr("stream1"),
	// 					StreamConfiguration: to.Ptr("{\"publishingInterval\":8,\"samplingInterval\":8,\"queueSize\":4}"),
	// 					TypeRef: to.Ptr("stream1TypeRef"),
	// 					LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-09T14:20:00.52Z"); return t}()),
	// 					Destinations: []armdeviceregistry.StreamDestinationClassification{
	// 						&armdeviceregistry.StreamStorageDestination{
	// 							Target: to.Ptr(armdeviceregistry.StreamDestinationTargetStorage),
	// 							Configuration: &armdeviceregistry.StorageDestinationConfiguration{
	// 								Path: to.Ptr("/tmp/stream1"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("stream2"),
	// 					StreamConfiguration: to.Ptr("{\"publishingInterval\":8,\"samplingInterval\":8,\"queueSize\":4}"),
	// 					TypeRef: to.Ptr("stream2TypeRef"),
	// 					LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-09T14:20:00.52Z"); return t}()),
	// 					Destinations: []armdeviceregistry.StreamDestinationClassification{
	// 						&armdeviceregistry.StreamMqttDestination{
	// 							Target: to.Ptr(armdeviceregistry.StreamDestinationTargetMqtt),
	// 							Configuration: &armdeviceregistry.MqttDestinationConfiguration{
	// 								Topic: to.Ptr("/contoso/testStream2"),
	// 								Retain: to.Ptr(armdeviceregistry.TopicRetainTypeNever),
	// 								Qos: to.Ptr(armdeviceregistry.MqttDestinationQosQos0),
	// 								TTL: to.Ptr[int64](7200),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ManagementGroups: []*armdeviceregistry.NamespaceDiscoveredManagementGroup{
	// 				{
	// 					Name: to.Ptr("managementGroup1"),
	// 					ManagementGroupConfiguration: to.Ptr("{\"retryCount\":10,\"retryBackoffInterval\":15}"),
	// 					TypeRef: to.Ptr("managementGroup1TypeRef"),
	// 					DefaultTopic: to.Ptr("/contoso/managementGroup1"),
	// 					DefaultTimeoutInSeconds: to.Ptr[int32](100),
	// 					LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-09T14:20:00.52Z"); return t}()),
	// 					Actions: []*armdeviceregistry.NamespaceDiscoveredManagementAction{
	// 						{
	// 							Name: to.Ptr("action1"),
	// 							ActionConfiguration: to.Ptr("{\"retryCount\":5,\"retryBackoffInterval\":5}"),
	// 							TargetURI: to.Ptr("/onvif/device_service?ONVIFProfile=Profile1"),
	// 							Topic: to.Ptr("/contoso/managementGroup1/action1"),
	// 							TypeRef: to.Ptr("action1TypeRef"),
	// 							ActionType: to.Ptr(armdeviceregistry.NamespaceDiscoveredManagementActionTypeCall),
	// 							TimeoutInSeconds: to.Ptr[int32](60),
	// 							LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-09T14:20:00.52Z"); return t}()),
	// 						},
	// 						{
	// 							Name: to.Ptr("action2"),
	// 							ActionConfiguration: to.Ptr("{\"retryCount\":5,\"retryBackoffInterval\":5}"),
	// 							TargetURI: to.Ptr("/onvif/device_service?ONVIFProfile=Profile2"),
	// 							Topic: to.Ptr("/contoso/managementGroup1/action2"),
	// 							TypeRef: to.Ptr("action2TypeRef"),
	// 							ActionType: to.Ptr(armdeviceregistry.NamespaceDiscoveredManagementActionTypeCall),
	// 							TimeoutInSeconds: to.Ptr[int32](60),
	// 							LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-09T14:20:00.52Z"); return t}()),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
