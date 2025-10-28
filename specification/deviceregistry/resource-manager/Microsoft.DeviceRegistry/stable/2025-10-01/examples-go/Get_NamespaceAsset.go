package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry/v2"
)

// Generated from example definition: 2025-10-01/Get_NamespaceAsset.json
func ExampleNamespaceAssetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespaceAssetsClient().Get(ctx, "myResourceGroup", "my-namespace-1", "my-asset-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdeviceregistry.NamespaceAssetsClientGetResponse{
	// 	NamespaceAsset: &armdeviceregistry.NamespaceAsset{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/namespaces/my-namespace-1/assets/my-asset-1"),
	// 		Name: to.Ptr("my-asset-1"),
	// 		Type: to.Ptr("Microsoft.DeviceRegistry/namespaces/assets"),
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
	// 		Properties: &armdeviceregistry.NamespaceAssetProperties{
	// 			UUID: to.Ptr("1824a74f-21e1-4458-ae07-604d3a241d2e"),
	// 			DeviceRef: &armdeviceregistry.DeviceRef{
	// 				DeviceName: to.Ptr("device1"),
	// 				EndpointName: to.Ptr("8ZBA6LRHU0A458969"),
	// 			},
	// 			AssetTypeRefs: []*string{
	// 				to.Ptr("myAssetTypeRef1"),
	// 				to.Ptr("myAssetTypeRef2"),
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			ExternalAssetID: to.Ptr("8ZBA6LRHU0A458969"),
	// 			DisplayName: to.Ptr("AssetDisplayName"),
	// 			Description: to.Ptr("This is a sample Asset"),
	// 			Version: to.Ptr[int64](12),
	// 			LastTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-11T20:14:21.949Z"); return t}()),
	// 			Manufacturer: to.Ptr("Contoso"),
	// 			ManufacturerURI: to.Ptr("https://www.contoso.com/manufacturerUri"),
	// 			Model: to.Ptr("ContosoModel"),
	// 			ProductCode: to.Ptr("SA34VDG"),
	// 			HardwareRevision: to.Ptr("1.0"),
	// 			SoftwareRevision: to.Ptr("2.0"),
	// 			DocumentationURI: to.Ptr("https://www.example.com/manual"),
	// 			SerialNumber: to.Ptr("64-103816-519918-8"),
	// 			Attributes: map[string]any{
	// 				"floor": "1",
	// 			},
	// 			DiscoveredAssetRefs: []*string{
	// 				to.Ptr("discoveredAsset1"),
	// 			},
	// 			DefaultDatasetsConfiguration: to.Ptr("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
	// 			DefaultEventsConfiguration: to.Ptr("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
	// 			DefaultStreamsConfiguration: to.Ptr("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
	// 			DefaultManagementGroupsConfiguration: to.Ptr("{\"retryCount\":10,\"retryBackoffInterval\":15}"),
	// 			DefaultDatasetsDestinations: []armdeviceregistry.DatasetDestinationClassification{
	// 				&armdeviceregistry.DatasetBrokerStateStoreDestination{
	// 					Target: to.Ptr(armdeviceregistry.DatasetDestinationTargetBrokerStateStore),
	// 					Configuration: &armdeviceregistry.BrokerStateStoreDestinationConfiguration{
	// 						Key: to.Ptr("defaultValue"),
	// 					},
	// 				},
	// 			},
	// 			DefaultEventsDestinations: []armdeviceregistry.EventDestinationClassification{
	// 				&armdeviceregistry.EventStorageDestination{
	// 					Target: to.Ptr(armdeviceregistry.EventDestinationTargetStorage),
	// 					Configuration: &armdeviceregistry.StorageDestinationConfiguration{
	// 						Path: to.Ptr("/tmp"),
	// 					},
	// 				},
	// 			},
	// 			DefaultStreamsDestinations: []armdeviceregistry.StreamDestinationClassification{
	// 				&armdeviceregistry.StreamMqttDestination{
	// 					Target: to.Ptr(armdeviceregistry.StreamDestinationTargetMqtt),
	// 					Configuration: &armdeviceregistry.MqttDestinationConfiguration{
	// 						Topic: to.Ptr("/contoso/test"),
	// 						Retain: to.Ptr(armdeviceregistry.TopicRetainTypeNever),
	// 						Qos: to.Ptr(armdeviceregistry.MqttDestinationQosQos0),
	// 						TTL: to.Ptr[int64](3600),
	// 					},
	// 				},
	// 			},
	// 			Datasets: []*armdeviceregistry.NamespaceDataset{
	// 				{
	// 					Name: to.Ptr("dataset1"),
	// 					DataSource: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/Oven;i=5"),
	// 					TypeRef: to.Ptr("dataset1TypeRef"),
	// 					DatasetConfiguration: to.Ptr("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
	// 					Destinations: []armdeviceregistry.DatasetDestinationClassification{
	// 						&armdeviceregistry.DatasetBrokerStateStoreDestination{
	// 							Target: to.Ptr(armdeviceregistry.DatasetDestinationTargetBrokerStateStore),
	// 							Configuration: &armdeviceregistry.BrokerStateStoreDestinationConfiguration{
	// 								Key: to.Ptr("dataset1"),
	// 							},
	// 						},
	// 					},
	// 					DataPoints: []*armdeviceregistry.NamespaceDatasetDataPoint{
	// 						{
	// 							Name: to.Ptr("dataset1DataPoint1"),
	// 							DataSource: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt3"),
	// 							DataPointConfiguration: to.Ptr("{\"publishingInterval\":8,\"samplingInterval\":8,\"queueSize\":4}"),
	// 							TypeRef: to.Ptr("dataset1DataPoint1TypeRef"),
	// 						},
	// 						{
	// 							Name: to.Ptr("dataset1DataPoint2"),
	// 							DataSource: to.Ptr("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt4"),
	// 							DataPointConfiguration: to.Ptr("{\"publishingInterval\":8,\"samplingInterval\":8,\"queueSize\":4}"),
	// 							TypeRef: to.Ptr("dataset1DataPoint2TypeRef"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Streams: []*armdeviceregistry.NamespaceStream{
	// 				{
	// 					Name: to.Ptr("stream1"),
	// 					StreamConfiguration: to.Ptr("{\"publishingInterval\":8,\"samplingInterval\":8,\"queueSize\":4}"),
	// 					TypeRef: to.Ptr("stream1TypeRef"),
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
	// 			ManagementGroups: []*armdeviceregistry.ManagementGroup{
	// 				{
	// 					Name: to.Ptr("managementGroup1"),
	// 					ManagementGroupConfiguration: to.Ptr("{\"retryCount\":10,\"retryBackoffInterval\":15}"),
	// 					TypeRef: to.Ptr("managementGroup1TypeRef"),
	// 					DefaultTopic: to.Ptr("/contoso/managementGroup1"),
	// 					DefaultTimeoutInSeconds: to.Ptr[int32](100),
	// 					Actions: []*armdeviceregistry.ManagementAction{
	// 						{
	// 							Name: to.Ptr("action1"),
	// 							ActionConfiguration: to.Ptr("{\"retryCount\":5,\"retryBackoffInterval\":5}"),
	// 							TargetURI: to.Ptr("/onvif/device_service?ONVIFProfile=Profile1"),
	// 							Topic: to.Ptr("/contoso/managementGroup1/action1"),
	// 							TypeRef: to.Ptr("action1TypeRef"),
	// 							ActionType: to.Ptr(armdeviceregistry.ManagementActionTypeCall),
	// 							TimeoutInSeconds: to.Ptr[int32](60),
	// 						},
	// 						{
	// 							Name: to.Ptr("action2"),
	// 							ActionConfiguration: to.Ptr("{\"retryCount\":5,\"retryBackoffInterval\":5}"),
	// 							TargetURI: to.Ptr("/onvif/device_service?ONVIFProfile=Profile2"),
	// 							Topic: to.Ptr("/contoso/managementGroup1/action2"),
	// 							TypeRef: to.Ptr("action2TypeRef"),
	// 							ActionType: to.Ptr(armdeviceregistry.ManagementActionTypeCall),
	// 							TimeoutInSeconds: to.Ptr[int32](60),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Status: &armdeviceregistry.NamespaceAssetStatus{
	// 				Config: &armdeviceregistry.StatusConfig{
	// 					Version: to.Ptr[int64](9),
	// 					LastTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-11T02:19:16.489Z"); return t}()),
	// 					Error: &armdeviceregistry.StatusError{
	// 						Code: to.Ptr("400"),
	// 						Message: to.Ptr("u"),
	// 						Details: []*armdeviceregistry.ErrorDetails{
	// 							{
	// 								Code: to.Ptr("400.123.456.789"),
	// 								Message: to.Ptr("Validation error"),
	// 								Info: to.Ptr("Property is not valid."),
	// 								CorrelationID: to.Ptr("xqoettlcdlxchoscv"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 				Datasets: []*armdeviceregistry.NamespaceAssetStatusDataset{
	// 					{
	// 						Name: to.Ptr("dataset1"),
	// 						MessageSchemaReference: &armdeviceregistry.NamespaceMessageSchemaReference{
	// 							SchemaRegistryNamespace: to.Ptr("liagdwhnvlhcptvmufws"),
	// 							SchemaName: to.Ptr("lytgdlsvivtcrtuvje"),
	// 							SchemaVersion: to.Ptr("1"),
	// 						},
	// 						Error: &armdeviceregistry.StatusError{
	// 							Code: to.Ptr("400"),
	// 							Message: to.Ptr("Error"),
	// 							Details: []*armdeviceregistry.ErrorDetails{
	// 								{
	// 									Code: to.Ptr("400.123.456.789"),
	// 									Message: to.Ptr("Validation error"),
	// 									Info: to.Ptr("Property is not valid."),
	// 									CorrelationID: to.Ptr("xqoettlcdlxchoscv"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				Streams: []*armdeviceregistry.NamespaceAssetStatusStream{
	// 					{
	// 						Name: to.Ptr("stream1"),
	// 						MessageSchemaReference: &armdeviceregistry.NamespaceMessageSchemaReference{
	// 							SchemaRegistryNamespace: to.Ptr("liagdwhnvlhcptvmufws"),
	// 							SchemaName: to.Ptr("lytgdlsvivtcrtuvje"),
	// 							SchemaVersion: to.Ptr("1"),
	// 						},
	// 						Error: &armdeviceregistry.StatusError{
	// 							Code: to.Ptr("400"),
	// 							Message: to.Ptr("Error"),
	// 							Details: []*armdeviceregistry.ErrorDetails{
	// 								{
	// 									Code: to.Ptr("400.123.456.789"),
	// 									Message: to.Ptr("Validation error"),
	// 									Info: to.Ptr("Property is not valid."),
	// 									CorrelationID: to.Ptr("xqoettlcdlxchoscv"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				ManagementGroups: []*armdeviceregistry.NamespaceAssetStatusManagementGroup{
	// 					{
	// 						Name: to.Ptr("managementGroup1"),
	// 						Actions: []*armdeviceregistry.NamespaceAssetStatusManagementAction{
	// 							{
	// 								Name: to.Ptr("action1"),
	// 								RequestMessageSchemaReference: &armdeviceregistry.NamespaceMessageSchemaReference{
	// 									SchemaRegistryNamespace: to.Ptr("liagdwhnvlhcptvmufws"),
	// 									SchemaName: to.Ptr("lytgdlsvivtcrtuvje"),
	// 									SchemaVersion: to.Ptr("1"),
	// 								},
	// 								ResponseMessageSchemaReference: &armdeviceregistry.NamespaceMessageSchemaReference{
	// 									SchemaRegistryNamespace: to.Ptr("liagdwhnvlhcptvmufws"),
	// 									SchemaName: to.Ptr("lytgdlsvivtcrtuvje"),
	// 									SchemaVersion: to.Ptr("1"),
	// 								},
	// 								Error: &armdeviceregistry.StatusError{
	// 									Code: to.Ptr("400"),
	// 									Message: to.Ptr("Error"),
	// 									Details: []*armdeviceregistry.ErrorDetails{
	// 										{
	// 											Code: to.Ptr("400.123.456.789"),
	// 											Message: to.Ptr("Validation error"),
	// 											Info: to.Ptr("Property is not valid."),
	// 											CorrelationID: to.Ptr("xqoettlcdlxchoscv"),
	// 										},
	// 									},
	// 								},
	// 							},
	// 							{
	// 								Name: to.Ptr("action2"),
	// 								RequestMessageSchemaReference: &armdeviceregistry.NamespaceMessageSchemaReference{
	// 									SchemaRegistryNamespace: to.Ptr("liagdwhnvlhcptvmufws"),
	// 									SchemaName: to.Ptr("lytgdlsvivtcrtuvje"),
	// 									SchemaVersion: to.Ptr("1"),
	// 								},
	// 								ResponseMessageSchemaReference: &armdeviceregistry.NamespaceMessageSchemaReference{
	// 									SchemaRegistryNamespace: to.Ptr("liagdwhnvlhcptvmufws"),
	// 									SchemaName: to.Ptr("lytgdlsvivtcrtuvje"),
	// 									SchemaVersion: to.Ptr("1"),
	// 								},
	// 								Error: &armdeviceregistry.StatusError{
	// 									Code: to.Ptr("400"),
	// 									Message: to.Ptr("Error"),
	// 									Details: []*armdeviceregistry.ErrorDetails{
	// 										{
	// 											Code: to.Ptr("400.123.456.789"),
	// 											Message: to.Ptr("Validation error"),
	// 											Info: to.Ptr("Property is not valid."),
	// 											CorrelationID: to.Ptr("xqoettlcdlxchoscv"),
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
