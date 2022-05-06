Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fiothub%2Farmiothub%2Fv0.5.0/sdk/resourcemanager/iothub/armiothub/README.md) on how to add the SDK to your project and authenticate.

```go
package armiothub_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_createOrUpdate.json
func ExampleResourceClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armiothub.Description{
			Location: to.Ptr("<location>"),
			Tags:     map[string]*string{},
			Etag:     to.Ptr("<etag>"),
			Properties: &armiothub.Properties{
				CloudToDevice: &armiothub.CloudToDeviceProperties{
					DefaultTTLAsIso8601: to.Ptr("<default-ttlas-iso8601>"),
					Feedback: &armiothub.FeedbackProperties{
						LockDurationAsIso8601: to.Ptr("<lock-duration-as-iso8601>"),
						MaxDeliveryCount:      to.Ptr[int32](10),
						TTLAsIso8601:          to.Ptr("<ttlas-iso8601>"),
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
						LockDurationAsIso8601: to.Ptr("<lock-duration-as-iso8601>"),
						MaxDeliveryCount:      to.Ptr[int32](10),
						TTLAsIso8601:          to.Ptr("<ttlas-iso8601>"),
					},
				},
				MinTLSVersion: to.Ptr("<min-tlsversion>"),
				NetworkRuleSets: &armiothub.NetworkRuleSetProperties{
					ApplyToBuiltInEventHubEndpoint: to.Ptr(true),
					DefaultAction:                  to.Ptr(armiothub.DefaultActionDeny),
					IPRules: []*armiothub.NetworkRuleSetIPRule{
						{
							Action:     to.Ptr(armiothub.NetworkRuleIPActionAllow),
							FilterName: to.Ptr("<filter-name>"),
							IPMask:     to.Ptr("<ipmask>"),
						},
						{
							Action:     to.Ptr(armiothub.NetworkRuleIPActionAllow),
							FilterName: to.Ptr("<filter-name>"),
							IPMask:     to.Ptr("<ipmask>"),
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
						Name:      to.Ptr("<name>"),
						Condition: to.Ptr("<condition>"),
						EndpointNames: []*string{
							to.Ptr("events")},
						IsEnabled: to.Ptr(true),
						Source:    to.Ptr(armiothub.RoutingSourceDeviceMessages),
					},
					Routes: []*armiothub.RouteProperties{},
				},
				StorageEndpoints: map[string]*armiothub.StorageEndpointProperties{
					"$default": {
						ConnectionString: to.Ptr("<connection-string>"),
						ContainerName:    to.Ptr("<container-name>"),
						SasTTLAsIso8601:  to.Ptr("<sas-ttlas-iso8601>"),
					},
				},
			},
			SKU: &armiothub.SKUInfo{
				Name:     to.Ptr(armiothub.IotHubSKUS1),
				Capacity: to.Ptr[int64](1),
			},
		},
		&armiothub.ResourceClientBeginCreateOrUpdateOptions{IfMatch: nil,
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
