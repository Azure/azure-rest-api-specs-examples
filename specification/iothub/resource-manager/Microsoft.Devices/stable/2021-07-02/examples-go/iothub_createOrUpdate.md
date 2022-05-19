Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fiothub%2Farmiothub%2Fv1.0.0/sdk/resourcemanager/iothub/armiothub/README.md) on how to add the SDK to your project and authenticate.

```go
package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_createOrUpdate.json
func ExampleResourceClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"testHub",
		armiothub.Description{
			Location: to.Ptr("centraluseuap"),
			Tags:     map[string]*string{},
			Etag:     to.Ptr("AAAAAAFD6M4="),
			Properties: &armiothub.Properties{
				CloudToDevice: &armiothub.CloudToDeviceProperties{
					DefaultTTLAsIso8601: to.Ptr("PT1H"),
					Feedback: &armiothub.FeedbackProperties{
						LockDurationAsIso8601: to.Ptr("PT1M"),
						MaxDeliveryCount:      to.Ptr[int32](10),
						TTLAsIso8601:          to.Ptr("PT1H"),
					},
					MaxDeliveryCount: to.Ptr[int32](10),
				},
				EnableDataResidency:           to.Ptr(true),
				EnableFileUploadNotifications: to.Ptr(false),
				EventHubEndpoints: map[string]*armiothub.EventHubProperties{
					"events": {
						PartitionCount:      to.Ptr[int32](2),
						RetentionTimeInDays: to.Ptr[int64](1),
					},
				},
				Features:      to.Ptr(armiothub.CapabilitiesNone),
				IPFilterRules: []*armiothub.IPFilterRule{},
				MessagingEndpoints: map[string]*armiothub.MessagingEndpointProperties{
					"fileNotifications": {
						LockDurationAsIso8601: to.Ptr("PT1M"),
						MaxDeliveryCount:      to.Ptr[int32](10),
						TTLAsIso8601:          to.Ptr("PT1H"),
					},
				},
				MinTLSVersion: to.Ptr("1.2"),
				NetworkRuleSets: &armiothub.NetworkRuleSetProperties{
					ApplyToBuiltInEventHubEndpoint: to.Ptr(true),
					DefaultAction:                  to.Ptr(armiothub.DefaultActionDeny),
					IPRules: []*armiothub.NetworkRuleSetIPRule{
						{
							Action:     to.Ptr(armiothub.NetworkRuleIPActionAllow),
							FilterName: to.Ptr("rule1"),
							IPMask:     to.Ptr("131.117.159.53"),
						},
						{
							Action:     to.Ptr(armiothub.NetworkRuleIPActionAllow),
							FilterName: to.Ptr("rule2"),
							IPMask:     to.Ptr("157.55.59.128/25"),
						}},
				},
				Routing: &armiothub.RoutingProperties{
					Endpoints: &armiothub.RoutingEndpoints{
						EventHubs:         []*armiothub.RoutingEventHubProperties{},
						ServiceBusQueues:  []*armiothub.RoutingServiceBusQueueEndpointProperties{},
						ServiceBusTopics:  []*armiothub.RoutingServiceBusTopicEndpointProperties{},
						StorageContainers: []*armiothub.RoutingStorageContainerProperties{},
					},
					FallbackRoute: &armiothub.FallbackRouteProperties{
						Name:      to.Ptr("$fallback"),
						Condition: to.Ptr("true"),
						EndpointNames: []*string{
							to.Ptr("events")},
						IsEnabled: to.Ptr(true),
						Source:    to.Ptr(armiothub.RoutingSourceDeviceMessages),
					},
					Routes: []*armiothub.RouteProperties{},
				},
				StorageEndpoints: map[string]*armiothub.StorageEndpointProperties{
					"$default": {
						ConnectionString: to.Ptr(""),
						ContainerName:    to.Ptr(""),
						SasTTLAsIso8601:  to.Ptr("PT1H"),
					},
				},
			},
			SKU: &armiothub.SKUInfo{
				Name:     to.Ptr(armiothub.IotHubSKUS1),
				Capacity: to.Ptr[int64](1),
			},
		},
		&armiothub.ResourceClientBeginCreateOrUpdateOptions{IfMatch: nil})
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
