package armiotoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
)

// Generated from example definition: 2026-03-01/AkriConnectorTemplate_CreateOrUpdate_MaximumSet_Gen.json
func ExampleAkriConnectorTemplateClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAkriConnectorTemplateClient().BeginCreateOrUpdate(ctx, "rgiotoperations", "resource-name123", "resource-name123", armiotoperations.AkriConnectorTemplateResource{
		Properties: &armiotoperations.AkriConnectorTemplateProperties{
			AioMetadata: &armiotoperations.AkriConnectorTemplateAioMetadata{
				AioMinVersion: to.Ptr("1.2.0"),
				AioMaxVersion: to.Ptr("1.4.0"),
			},
			RuntimeConfiguration: &armiotoperations.AkriConnectorTemplateManagedConfiguration{
				RuntimeConfigurationType: to.Ptr(armiotoperations.AkriConnectorTemplateRuntimeConfigurationTypeManagedConfiguration),
				ManagedConfigurationSettings: &armiotoperations.AkriConnectorTemplateRuntimeImageConfiguration{
					ManagedConfigurationType: to.Ptr(armiotoperations.AkriConnectorTemplateManagedConfigurationTypeImageConfiguration),
					ImageConfigurationSettings: &armiotoperations.AkriConnectorTemplateRuntimeImageConfigurationSettings{
						RegistrySettings: &armiotoperations.AkriConnectorsContainerRegistry{
							RegistrySettingsType: to.Ptr(armiotoperations.AkriConnectorsRegistrySettingsTypeContainerRegistry),
							ContainerRegistrySettings: &armiotoperations.AkriConnectorsContainerRegistrySettings{
								Registry: to.Ptr("akribuilds.azurecr.io"),
							},
						},
						ImageName: to.Ptr("akri-connectors/rest"),
						TagDigestSettings: &armiotoperations.AkriConnectorsTag{
							TagDigestType: to.Ptr(armiotoperations.AkriConnectorsTagDigestTypeTag),
							Tag:           to.Ptr("0.5.0-20250825.4"),
						},
					},
				},
			},
			Diagnostics: &armiotoperations.AkriConnectorTemplateDiagnostics{
				Logs: &armiotoperations.AkriConnectorsDiagnosticsLogs{
					Level: to.Ptr("info"),
				},
			},
			DeviceInboundEndpointTypes: []*armiotoperations.AkriConnectorTemplateDeviceInboundEndpointType{
				{
					EndpointType: to.Ptr("Microsoft.Rest"),
					Version:      to.Ptr("0.0.1"),
				},
			},
			MqttConnectionConfiguration: &armiotoperations.AkriConnectorsMqttConnectionConfiguration{
				Authentication: &armiotoperations.AkriConnectorsServiceAccountAuthentication{
					Method: to.Ptr(armiotoperations.AkriConnectorsMqttAuthenticationMethodServiceAccountToken),
					ServiceAccountTokenSettings: &armiotoperations.AkriConnectorsServiceAccountTokenSettings{
						Audience: to.Ptr("MQ-SAT"),
					},
				},
				Host:                 to.Ptr("aio-broker:18883"),
				Protocol:             to.Ptr(armiotoperations.AkriConnectorsMqttProtocolTypeMqtt),
				KeepAliveSeconds:     to.Ptr[int32](10),
				MaxInflightMessages:  to.Ptr[int32](10),
				SessionExpirySeconds: to.Ptr[int32](60),
				TLS: &armiotoperations.TLSProperties{
					Mode:                             to.Ptr(armiotoperations.OperationalModeEnabled),
					TrustedCaCertificateConfigMapRef: to.Ptr("azure-iot-operations-aio-ca-trust-bundle"),
				},
			},
		},
		ExtendedLocation: &armiotoperations.ExtendedLocation{
			Name: to.Ptr("/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123"),
			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
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
	// res = armiotoperations.AkriConnectorTemplateClientCreateOrUpdateResponse{
	// 	AkriConnectorTemplateResource: &armiotoperations.AkriConnectorTemplateResource{
	// 		Properties: &armiotoperations.AkriConnectorTemplateProperties{
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 			AioMetadata: &armiotoperations.AkriConnectorTemplateAioMetadata{
	// 				AioMinVersion: to.Ptr("1.2.0"),
	// 				AioMaxVersion: to.Ptr("1.4.0"),
	// 			},
	// 			RuntimeConfiguration: &armiotoperations.AkriConnectorTemplateManagedConfiguration{
	// 				RuntimeConfigurationType: to.Ptr(armiotoperations.AkriConnectorTemplateRuntimeConfigurationTypeManagedConfiguration),
	// 				ManagedConfigurationSettings: &armiotoperations.AkriConnectorTemplateRuntimeImageConfiguration{
	// 					ManagedConfigurationType: to.Ptr(armiotoperations.AkriConnectorTemplateManagedConfigurationTypeImageConfiguration),
	// 					ImageConfigurationSettings: &armiotoperations.AkriConnectorTemplateRuntimeImageConfigurationSettings{
	// 						RegistrySettings: &armiotoperations.AkriConnectorsContainerRegistry{
	// 							RegistrySettingsType: to.Ptr(armiotoperations.AkriConnectorsRegistrySettingsTypeContainerRegistry),
	// 							ContainerRegistrySettings: &armiotoperations.AkriConnectorsContainerRegistrySettings{
	// 								Registry: to.Ptr("akribuilds.azurecr.io"),
	// 							},
	// 						},
	// 						ImageName: to.Ptr("akri-connectors/rest"),
	// 						TagDigestSettings: &armiotoperations.AkriConnectorsTag{
	// 							TagDigestType: to.Ptr(armiotoperations.AkriConnectorsTagDigestTypeTag),
	// 							Tag: to.Ptr("0.5.0-20250825.4"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Diagnostics: &armiotoperations.AkriConnectorTemplateDiagnostics{
	// 				Logs: &armiotoperations.AkriConnectorsDiagnosticsLogs{
	// 					Level: to.Ptr("info"),
	// 				},
	// 			},
	// 			DeviceInboundEndpointTypes: []*armiotoperations.AkriConnectorTemplateDeviceInboundEndpointType{
	// 				{
	// 					EndpointType: to.Ptr("Microsoft.Rest"),
	// 					Version: to.Ptr("0.0.1"),
	// 				},
	// 			},
	// 			MqttConnectionConfiguration: &armiotoperations.AkriConnectorsMqttConnectionConfiguration{
	// 				Authentication: &armiotoperations.AkriConnectorsServiceAccountAuthentication{
	// 					Method: to.Ptr(armiotoperations.AkriConnectorsMqttAuthenticationMethodServiceAccountToken),
	// 					ServiceAccountTokenSettings: &armiotoperations.AkriConnectorsServiceAccountTokenSettings{
	// 						Audience: to.Ptr("MQ-SAT"),
	// 					},
	// 				},
	// 				Host: to.Ptr("aio-broker:18883"),
	// 				Protocol: to.Ptr(armiotoperations.AkriConnectorsMqttProtocolTypeMqtt),
	// 				KeepAliveSeconds: to.Ptr[int32](10),
	// 				MaxInflightMessages: to.Ptr[int32](10),
	// 				SessionExpirySeconds: to.Ptr[int32](60),
	// 				TLS: &armiotoperations.TLSProperties{
	// 					Mode: to.Ptr(armiotoperations.OperationalModeEnabled),
	// 					TrustedCaCertificateConfigMapRef: to.Ptr("azure-iot-operations-aio-ca-trust-bundle"),
	// 				},
	// 			},
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/akriConnectorTemplates/resource-name123"),
	// 		Name: to.Ptr("bfimycofjtzxduufwanuxwoudsh"),
	// 		Type: to.Ptr("wnabnudmydrgpccqusxkmsmgcwzwh"),
	// 		SystemData: &armiotoperations.SystemData{
	// 			CreatedBy: to.Ptr("contosouser"),
	// 			CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("contosouser"),
	// 			LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 		},
	// 	},
	// }
}
