package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub/v2"
)

// Generated from example definition: 2026-10-01-preview/Get_IotHub_With_DeviceRegistry.json
func ExampleResourceClient_Get_getIotHubWithDeviceRegistry() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceClient().Get(ctx, "myResourceGroup", "testHub", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiothub.ResourceClientGetResponse{
	// 	Description: armiothub.Description{
	// 		ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/IotHubs/testHub"),
	// 		Name: to.Ptr("testHub"),
	// 		Type: to.Ptr("Microsoft.Devices/IotHubs"),
	// 		Location: to.Ptr("centraluseuap"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Etag: to.Ptr("AAAAAAFD6M4="),
	// 		Properties: &armiothub.Properties{
	// 			State: to.Ptr("Active"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			IPFilterRules: []*armiothub.IPFilterRule{
	// 				{
	// 					FilterName: to.Ptr("rule1"),
	// 					Action: to.Ptr(armiothub.IPFilterActionTypeAccept),
	// 					IPMask: to.Ptr("131.117.159.53"),
	// 				},
	// 				{
	// 					FilterName: to.Ptr("rule2"),
	// 					Action: to.Ptr(armiothub.IPFilterActionTypeAccept),
	// 					IPMask: to.Ptr("157.55.59.128/25"),
	// 				},
	// 			},
	// 			NetworkRuleSets: &armiothub.NetworkRuleSetProperties{
	// 				DefaultAction: to.Ptr(armiothub.DefaultActionDeny),
	// 				ApplyToBuiltInEventHubEndpoint: to.Ptr(true),
	// 				IPRules: []*armiothub.NetworkRuleSetIPRule{
	// 					{
	// 						FilterName: to.Ptr("rule1"),
	// 						Action: to.Ptr(armiothub.NetworkRuleIPActionAllow),
	// 						IPMask: to.Ptr("131.117.159.53"),
	// 					},
	// 					{
	// 						FilterName: to.Ptr("rule2"),
	// 						Action: to.Ptr(armiothub.NetworkRuleIPActionAllow),
	// 						IPMask: to.Ptr("157.55.59.128/25"),
	// 					},
	// 				},
	// 			},
	// 			HostName: to.Ptr("iot-dps-cit-hub-1.azure-devices.net"),
	// 			DeviceHostName: to.Ptr("iot-dps-cit-hub-1.device.azure-devices.net"),
	// 			ServiceHostName: to.Ptr("iot-dps-cit-hub-1.service.azure-devices.net"),
	// 			EventHubEndpoints: map[string]*armiothub.EventHubProperties{
	// 				"events": &armiothub.EventHubProperties{
	// 					RetentionTimeInDays: to.Ptr[int64](1),
	// 					PartitionCount: to.Ptr[int32](2),
	// 					PartitionIDs: []*string{
	// 						to.Ptr("0"),
	// 						to.Ptr("1"),
	// 					},
	// 					Path: to.Ptr("iot-dps-cit-hub-1"),
	// 					Endpoint: to.Ptr("sb://iothub-ns-iot-dps-ci-245306-76aca8e13b.servicebus.windows.net/"),
	// 				},
	// 			},
	// 			Routing: &armiothub.RoutingProperties{
	// 				Endpoints: &armiothub.RoutingEndpoints{
	// 					ServiceBusQueues: []*armiothub.RoutingServiceBusQueueEndpointProperties{
	// 					},
	// 					ServiceBusTopics: []*armiothub.RoutingServiceBusTopicEndpointProperties{
	// 					},
	// 					EventHubs: []*armiothub.RoutingEventHubProperties{
	// 					},
	// 					StorageContainers: []*armiothub.RoutingStorageContainerProperties{
	// 					},
	// 				},
	// 				Routes: []*armiothub.RouteProperties{
	// 				},
	// 				FallbackRoute: &armiothub.FallbackRouteProperties{
	// 					Source: to.Ptr(armiothub.RoutingSourceDeviceMessages),
	// 					Condition: to.Ptr("true"),
	// 					EndpointNames: []*string{
	// 						to.Ptr("events"),
	// 					},
	// 					IsEnabled: to.Ptr(true),
	// 				},
	// 			},
	// 			StorageEndpoints: map[string]*armiothub.StorageEndpointProperties{
	// 				"$default": &armiothub.StorageEndpointProperties{
	// 					SasTTLAsIso8601: to.Ptr("PT1H"),
	// 					ConnectionString: to.Ptr(""),
	// 					ContainerName: to.Ptr(""),
	// 				},
	// 			},
	// 			MessagingEndpoints: map[string]*armiothub.MessagingEndpointProperties{
	// 				"fileNotifications": &armiothub.MessagingEndpointProperties{
	// 					LockDurationAsIso8601: to.Ptr("PT1M"),
	// 					TTLAsIso8601: to.Ptr("PT1H"),
	// 					MaxDeliveryCount: to.Ptr[int32](10),
	// 				},
	// 			},
	// 			EnableFileUploadNotifications: to.Ptr(false),
	// 			CloudToDevice: &armiothub.CloudToDeviceProperties{
	// 				MaxDeliveryCount: to.Ptr[int32](10),
	// 				DefaultTTLAsIso8601: to.Ptr("PT1H"),
	// 				Feedback: &armiothub.FeedbackProperties{
	// 					LockDurationAsIso8601: to.Ptr("PT1M"),
	// 					TTLAsIso8601: to.Ptr("PT1H"),
	// 					MaxDeliveryCount: to.Ptr[int32](10),
	// 				},
	// 			},
	// 			Features: to.Ptr(armiothub.CapabilitiesNone),
	// 			DeviceStreams: &armiothub.PropertiesDeviceStreams{
	// 				StreamingEndpoints: []*string{
	// 					to.Ptr("https://streams.azure-devices-int.net:9443"),
	// 				},
	// 			},
	// 			Locations: []*armiothub.LocationDescription{
	// 				{
	// 					Location: to.Ptr("West US"),
	// 					Role: to.Ptr(armiothub.IotHubReplicaRoleTypePrimary),
	// 				},
	// 				{
	// 					Location: to.Ptr("East US"),
	// 					Role: to.Ptr(armiothub.IotHubReplicaRoleTypeSecondary),
	// 				},
	// 			},
	// 			MinTLSVersion: to.Ptr("1.2"),
	// 			IPVersion: to.Ptr(armiothub.IPVersionIPv4IPv6),
	// 			DeviceRegistry: &armiothub.DeviceRegistry{
	// 				NamespaceResourceID: to.Ptr("/subscriptions/ae24ff83-d2ca-4fc8-9717-05dae4bba489/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/namespaces/testNamespace"),
	// 				NamespaceUUID: to.Ptr("0dd781aa-f7a9-4f7e-8105-0579ce78f407"),
	// 				DataPlaneHostName: to.Ptr("centraluseuap.api.deviceregistry.com"),
	// 				Identity: &armiothub.DeviceRegistryIdentity{
	// 					Type: to.Ptr(armiothub.DeviceRegistryIdentityTypeUserAssigned),
	// 					UserAssignedIdentity: to.Ptr("/subscriptions/ae24ff83-d2ca-4fc8-9717-05dae4bba489/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testIdentity"),
	// 				},
	// 				LinkingProperties: &armiothub.DeviceRegistryLinkingProperties{
	// 					State: to.Ptr(armiothub.DeviceRegistryLinkingStateSuccess),
	// 				},
	// 			},
	// 			IotHubDetails: &armiothub.Details{
	// 				GatewayVersion: to.Ptr(armiothub.GatewayVersionV2),
	// 			},
	// 		},
	// 		SKU: &armiothub.SKUInfo{
	// 			Name: to.Ptr(armiothub.IotHubSKUS1),
	// 			Tier: to.Ptr(armiothub.IotHubSKUTierStandard),
	// 			Capacity: to.Ptr[int64](1),
	// 		},
	// 		SystemData: &armiothub.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2021, time.January, 30, 0, 28, 38, 963000000, time.UTC)),
	// 		},
	// 	},
	// }
}
