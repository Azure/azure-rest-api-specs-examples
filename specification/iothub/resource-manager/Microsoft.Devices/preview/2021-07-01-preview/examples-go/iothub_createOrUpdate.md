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

// x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2021-07-01-preview/examples/iothub_createOrUpdate.json
func ExampleResourceClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiothub.NewResourceClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armiothub.Description{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Etag:     to.StringPtr("<etag>"),
			Properties: &armiothub.Properties{
				CloudToDevice: &armiothub.CloudToDeviceProperties{
					DefaultTTLAsIso8601: to.StringPtr("<default-ttlas-iso8601>"),
					Feedback: &armiothub.FeedbackProperties{
						LockDurationAsIso8601: to.StringPtr("<lock-duration-as-iso8601>"),
						MaxDeliveryCount:      to.Int32Ptr(10),
						TTLAsIso8601:          to.StringPtr("<ttlas-iso8601>"),
					},
					MaxDeliveryCount: to.Int32Ptr(10),
				},
				EnableFileUploadNotifications: to.BoolPtr(false),
				EventHubEndpoints: map[string]*armiothub.EventHubProperties{
					"events": {
						PartitionCount:      to.Int32Ptr(2),
						RetentionTimeInDays: to.Int64Ptr(1),
					},
				},
				Features:      armiothub.Capabilities("None").ToPtr(),
				IPFilterRules: []*armiothub.IPFilterRule{},
				MessagingEndpoints: map[string]*armiothub.MessagingEndpointProperties{
					"fileNotifications": {
						LockDurationAsIso8601: to.StringPtr("<lock-duration-as-iso8601>"),
						MaxDeliveryCount:      to.Int32Ptr(10),
						TTLAsIso8601:          to.StringPtr("<ttlas-iso8601>"),
					},
				},
				MinTLSVersion: to.StringPtr("<min-tlsversion>"),
				NetworkRuleSets: &armiothub.NetworkRuleSetProperties{
					ApplyToBuiltInEventHubEndpoint: to.BoolPtr(true),
					DefaultAction:                  armiothub.DefaultAction("Deny").ToPtr(),
					IPRules: []*armiothub.NetworkRuleSetIPRule{
						{
							Action:     armiothub.NetworkRuleIPAction("Allow").ToPtr(),
							FilterName: to.StringPtr("<filter-name>"),
							IPMask:     to.StringPtr("<ipmask>"),
						},
						{
							Action:     armiothub.NetworkRuleIPAction("Allow").ToPtr(),
							FilterName: to.StringPtr("<filter-name>"),
							IPMask:     to.StringPtr("<ipmask>"),
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
						Name:      to.StringPtr("<name>"),
						Condition: to.StringPtr("<condition>"),
						EndpointNames: []*string{
							to.StringPtr("events")},
						IsEnabled: to.BoolPtr(true),
						Source:    armiothub.RoutingSource("DeviceMessages").ToPtr(),
					},
					Routes: []*armiothub.RouteProperties{},
				},
				StorageEndpoints: map[string]*armiothub.StorageEndpointProperties{
					"$default": {
						ConnectionString: to.StringPtr("<connection-string>"),
						ContainerName:    to.StringPtr("<container-name>"),
						SasTTLAsIso8601:  to.StringPtr("<sas-ttlas-iso8601>"),
					},
				},
			},
			SKU: &armiothub.SKUInfo{
				Name:     armiothub.IotHubSKU("S1").ToPtr(),
				Capacity: to.Int64Ptr(1),
			},
		},
		&armiothub.ResourceClientBeginCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ResourceClientCreateOrUpdateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fiothub%2Farmiothub%2Fv0.3.1/sdk/resourcemanager/iothub/armiothub/README.md) on how to add the SDK to your project and authenticate.
