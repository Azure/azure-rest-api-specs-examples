package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub/v2"
)

// Generated from example definition: 2026-10-01-preview/CreateOrReplace_IotHub.json
func ExampleResourceClient_BeginCreateOrUpdate_createOrReplaceIotHub() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewResourceClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "testHub", armiothub.Description{
		Location: to.Ptr("centraluseuap"),
		Tags:     map[string]*string{},
		Etag:     to.Ptr("AAAAAAFD6M4="),
		Properties: &armiothub.Properties{
			IPFilterRules: []*armiothub.IPFilterRule{},
			NetworkRuleSets: &armiothub.NetworkRuleSetProperties{
				DefaultAction:                  to.Ptr(armiothub.DefaultActionDeny),
				ApplyToBuiltInEventHubEndpoint: to.Ptr(true),
				IPRules: []*armiothub.NetworkRuleSetIPRule{
					{
						FilterName: to.Ptr("rule1"),
						Action:     to.Ptr(armiothub.NetworkRuleIPActionAllow),
						IPMask:     to.Ptr("131.117.159.53"),
					},
					{
						FilterName: to.Ptr("rule2"),
						Action:     to.Ptr(armiothub.NetworkRuleIPActionAllow),
						IPMask:     to.Ptr("157.55.59.128/25"),
					},
				},
			},
			EventHubEndpoints: map[string]*armiothub.EventHubProperties{
				"events": {
					RetentionTimeInDays: to.Ptr[int64](1),
					PartitionCount:      to.Ptr[int32](2),
				},
			},
			Routing: &armiothub.RoutingProperties{
				Endpoints: &armiothub.RoutingEndpoints{
					ServiceBusQueues:  []*armiothub.RoutingServiceBusQueueEndpointProperties{},
					ServiceBusTopics:  []*armiothub.RoutingServiceBusTopicEndpointProperties{},
					EventHubs:         []*armiothub.RoutingEventHubProperties{},
					StorageContainers: []*armiothub.RoutingStorageContainerProperties{},
				},
				Routes: []*armiothub.RouteProperties{
					{
						Name:       to.Ptr("Routeid"),
						Source:     to.Ptr(armiothub.RoutingSourceDeviceMessages),
						Condition:  to.Ptr("true"),
						DataSchema: to.Ptr("aio-sr://aiosaalkopkedev/62a24af1d7db61cd44b2ad6b6c3f4ab7312be447f89ff3401d18357d0d05ce3a:1"),
						EndpointNames: []*string{
							to.Ptr("events"),
						},
						IsEnabled: to.Ptr(true),
					},
				},
				FallbackRoute: &armiothub.FallbackRouteProperties{
					Name:      to.Ptr("$fallback"),
					Source:    to.Ptr(armiothub.RoutingSourceDeviceMessages),
					Condition: to.Ptr("true"),
					EndpointNames: []*string{
						to.Ptr("events"),
					},
					IsEnabled: to.Ptr(true),
				},
			},
			StorageEndpoints: map[string]*armiothub.StorageEndpointProperties{
				"$default": {
					SasTTLAsIso8601:  to.Ptr("PT1H"),
					ConnectionString: to.Ptr(""),
					ContainerName:    to.Ptr(""),
				},
			},
			MessagingEndpoints: map[string]*armiothub.MessagingEndpointProperties{
				"fileNotifications": {
					LockDurationAsIso8601: to.Ptr("PT1M"),
					TTLAsIso8601:          to.Ptr("PT1H"),
					MaxDeliveryCount:      to.Ptr[int32](10),
				},
			},
			EnableFileUploadNotifications: to.Ptr(false),
			CloudToDevice: &armiothub.CloudToDeviceProperties{
				MaxDeliveryCount:    to.Ptr[int32](10),
				DefaultTTLAsIso8601: to.Ptr("PT1H"),
				Feedback: &armiothub.FeedbackProperties{
					LockDurationAsIso8601: to.Ptr("PT1M"),
					TTLAsIso8601:          to.Ptr("PT1H"),
					MaxDeliveryCount:      to.Ptr[int32](10),
				},
			},
			Features:            to.Ptr(armiothub.CapabilitiesNone),
			MinTLSVersion:       to.Ptr("1.2"),
			EnableDataResidency: to.Ptr(true),
			RootCertificate: &armiothub.RootCertificateProperties{
				EnableRootCertificateV2: to.Ptr(true),
			},
			IPVersion: to.Ptr(armiothub.IPVersionIPv4IPv6),
		},
		SKU: &armiothub.SKUInfo{
			Name:     to.Ptr(armiothub.IotHubSKUS1),
			Capacity: to.Ptr[int64](1),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiothub.ResourceClientCreateOrUpdateResponse{
	// 	Description: armiothub.Description{
	// 		ID: to.Ptr("/subscriptions/ae24ff83-d2ca-4fc8-9717-05dae4bba489/resourceGroups/myResourceGroup/providers/Microsoft.Devices/IotHubs/testHub"),
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
	// 					{
	// 						Name: to.Ptr("Routeid"),
	// 						Source: to.Ptr(armiothub.RoutingSourceDeviceMessages),
	// 						Condition: to.Ptr("true"),
	// 						DataSchema: to.Ptr("aio-sr://aiosaalkopkedev/62a24af1d7db61cd44b2ad6b6c3f4ab7312be447f89ff3401d18357d0d05ce3a:1"),
	// 						EndpointNames: []*string{
	// 							to.Ptr("events"),
	// 						},
	// 						IsEnabled: to.Ptr(true),
	// 					},
	// 				},
	// 				FallbackRoute: &armiothub.FallbackRouteProperties{
	// 					Name: to.Ptr("$fallback"),
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
	// 			MinTLSVersion: to.Ptr("1.2"),
	// 			EnableDataResidency: to.Ptr(true),
	// 			RootCertificate: &armiothub.RootCertificateProperties{
	// 				EnableRootCertificateV2: to.Ptr(true),
	// 				LastUpdatedTimeUTC: to.Ptr(time.Date(2022, time.April, 8, 18, 16, 49, 74895800, time.UTC)),
	// 			},
	// 			IPVersion: to.Ptr(armiothub.IPVersionIPv4IPv6),
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
