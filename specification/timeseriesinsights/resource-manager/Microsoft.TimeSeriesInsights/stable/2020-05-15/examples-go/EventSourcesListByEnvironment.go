package armtimeseriesinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/EventSourcesListByEnvironment.json
func ExampleEventSourcesClient_ListByEnvironment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtimeseriesinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEventSourcesClient().ListByEnvironment(ctx, "rg1", "env1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EventSourceListResponse = armtimeseriesinsights.EventSourceListResponse{
	// 	Value: []armtimeseriesinsights.EventSourceResourceClassification{
	// 		&armtimeseriesinsights.EventHubEventSourceResource{
	// 			Name: to.Ptr("es1"),
	// 			Type: to.Ptr("Microsoft.TimeSeriesInsights/Environments/EventSources"),
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.TimeSeriesInsights/Environments/env1/eventSources/es1"),
	// 			Location: to.Ptr("West US"),
	// 			Tags: map[string]*string{
	// 			},
	// 			Kind: to.Ptr(armtimeseriesinsights.EventSourceResourceKindMicrosoftEventHub),
	// 			Properties: &armtimeseriesinsights.EventHubEventSourceResourceProperties{
	// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-18T19:20:33.228Z"); return t}()),
	// 				ProvisioningState: to.Ptr(armtimeseriesinsights.ProvisioningStateSucceeded),
	// 				IngressStartAt: &armtimeseriesinsights.IngressStartAtProperties{
	// 					Type: to.Ptr(armtimeseriesinsights.IngressStartAtTypeEarliestAvailable),
	// 				},
	// 				LocalTimestamp: &armtimeseriesinsights.LocalTimestamp{
	// 					Format: to.Ptr(armtimeseriesinsights.LocalTimestampFormat("TimeSpan")),
	// 					TimeZoneOffset: &armtimeseriesinsights.LocalTimestampTimeZoneOffset{
	// 						PropertyName: to.Ptr("someEventPropertyName"),
	// 					},
	// 				},
	// 				EventSourceResourceID: to.Ptr("somePathInArm"),
	// 				ConsumerGroupName: to.Ptr("cgn"),
	// 				EventHubName: to.Ptr("ehn"),
	// 				KeyName: to.Ptr("managementKey"),
	// 				ServiceBusNamespace: to.Ptr("sbn"),
	// 			},
	// 	}},
	// }
}
