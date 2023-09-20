package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c178f2467f792a322f56174be244135d2c907f/specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_delete.json
func ExampleResourceClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewResourceClient().BeginDelete(ctx, "myResourceGroup", "testHub", nil)
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
	// res.Description = armiothub.Description{
	// 	Name: to.Ptr("testHub"),
	// 	Type: to.Ptr("Microsoft.Devices/IotHubs"),
	// 	ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/IotHubs/testHub"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Etag: to.Ptr("AAAAAAFD6M4="),
	// 	Properties: &armiothub.Properties{
	// 		CloudToDevice: &armiothub.CloudToDeviceProperties{
	// 			DefaultTTLAsIso8601: to.Ptr("PT1H"),
	// 			Feedback: &armiothub.FeedbackProperties{
	// 				LockDurationAsIso8601: to.Ptr("PT1M"),
	// 				MaxDeliveryCount: to.Ptr[int32](10),
	// 				TTLAsIso8601: to.Ptr("PT1H"),
	// 			},
	// 			MaxDeliveryCount: to.Ptr[int32](10),
	// 		},
	// 		EnableFileUploadNotifications: to.Ptr(false),
	// 		EventHubEndpoints: map[string]*armiothub.EventHubProperties{
	// 			"events": &armiothub.EventHubProperties{
	// 				Path: to.Ptr("iot-dps-cit-hub-1"),
	// 				Endpoint: to.Ptr("sb://iothub-ns-iot-dps-ci-245306-76aca8e13b.servicebus.windows.net/"),
	// 				PartitionCount: to.Ptr[int32](2),
	// 				PartitionIDs: []*string{
	// 					to.Ptr("0"),
	// 					to.Ptr("1")},
	// 					RetentionTimeInDays: to.Ptr[int64](1),
	// 				},
	// 			},
	// 			Features: to.Ptr(armiothub.CapabilitiesNone),
	// 			HostName: to.Ptr("iot-dps-cit-hub-1.azure-devices.net"),
	// 			IPFilterRules: []*armiothub.IPFilterRule{
	// 				{
	// 					Action: to.Ptr(armiothub.IPFilterActionTypeAccept),
	// 					FilterName: to.Ptr("rule1"),
	// 					IPMask: to.Ptr("131.117.159.53"),
	// 				},
	// 				{
	// 					Action: to.Ptr(armiothub.IPFilterActionTypeAccept),
	// 					FilterName: to.Ptr("rule2"),
	// 					IPMask: to.Ptr("157.55.59.128/25"),
	// 			}},
	// 			MessagingEndpoints: map[string]*armiothub.MessagingEndpointProperties{
	// 				"fileNotifications": &armiothub.MessagingEndpointProperties{
	// 					LockDurationAsIso8601: to.Ptr("PT1M"),
	// 					MaxDeliveryCount: to.Ptr[int32](10),
	// 					TTLAsIso8601: to.Ptr("PT1H"),
	// 				},
	// 			},
	// 			NetworkRuleSets: &armiothub.NetworkRuleSetProperties{
	// 				ApplyToBuiltInEventHubEndpoint: to.Ptr(true),
	// 				DefaultAction: to.Ptr(armiothub.DefaultActionDeny),
	// 				IPRules: []*armiothub.NetworkRuleSetIPRule{
	// 					{
	// 						Action: to.Ptr(armiothub.NetworkRuleIPActionAllow),
	// 						FilterName: to.Ptr("rule1"),
	// 						IPMask: to.Ptr("131.117.159.53"),
	// 					},
	// 					{
	// 						Action: to.Ptr(armiothub.NetworkRuleIPActionAllow),
	// 						FilterName: to.Ptr("rule2"),
	// 						IPMask: to.Ptr("157.55.59.128/25"),
	// 				}},
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			Routing: &armiothub.RoutingProperties{
	// 				Endpoints: &armiothub.RoutingEndpoints{
	// 					EventHubs: []*armiothub.RoutingEventHubProperties{
	// 					},
	// 					ServiceBusQueues: []*armiothub.RoutingServiceBusQueueEndpointProperties{
	// 					},
	// 					ServiceBusTopics: []*armiothub.RoutingServiceBusTopicEndpointProperties{
	// 					},
	// 					StorageContainers: []*armiothub.RoutingStorageContainerProperties{
	// 					},
	// 				},
	// 				FallbackRoute: &armiothub.FallbackRouteProperties{
	// 					Condition: to.Ptr("true"),
	// 					EndpointNames: []*string{
	// 						to.Ptr("events")},
	// 						IsEnabled: to.Ptr(true),
	// 						Source: to.Ptr(armiothub.RoutingSourceDeviceMessages),
	// 					},
	// 					Routes: []*armiothub.RouteProperties{
	// 					},
	// 				},
	// 				State: to.Ptr("Active"),
	// 				StorageEndpoints: map[string]*armiothub.StorageEndpointProperties{
	// 					"$default": &armiothub.StorageEndpointProperties{
	// 						ConnectionString: to.Ptr(""),
	// 						ContainerName: to.Ptr(""),
	// 						SasTTLAsIso8601: to.Ptr("PT1H"),
	// 					},
	// 				},
	// 			},
	// 			SKU: &armiothub.SKUInfo{
	// 				Name: to.Ptr(armiothub.IotHubSKUS1),
	// 				Capacity: to.Ptr[int64](1),
	// 				Tier: to.Ptr(armiothub.IotHubSKUTierStandard),
	// 			},
	// 			SystemData: &armiothub.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-30T00:28:38.963Z"); return t}()),
	// 			},
	// 		}
}
