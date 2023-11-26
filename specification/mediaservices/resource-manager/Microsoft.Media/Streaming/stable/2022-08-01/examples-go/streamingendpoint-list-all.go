package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/streamingendpoint-list-all.json
func ExampleStreamingEndpointsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStreamingEndpointsClient().NewListPager("mediaresources", "slitestmedia10", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.StreamingEndpointListResult = armmediaservices.StreamingEndpointListResult{
		// 	Value: []*armmediaservices.StreamingEndpoint{
		// 		{
		// 			Name: to.Ptr("myStreamingEndpoint1"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/streamingEndpoints"),
		// 			ID: to.Ptr("/subscriptions/0a6ec948-5a62-437d-b9df-934dc7c1b722/resourceGroups/mediaresources/providers/Microsoft.Media/mediaservices/slitestmedia10/streamingendpoints/myStreamingEndpoint1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("value1"),
		// 				"tag2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armmediaservices.StreamingEndpointProperties{
		// 				Description: to.Ptr("test event 1"),
		// 				AvailabilitySetName: to.Ptr("availableset"),
		// 				CdnEnabled: to.Ptr(false),
		// 				CdnProfile: to.Ptr(""),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:09.202Z"); return t}()),
		// 				CustomHostNames: []*string{
		// 				},
		// 				FreeTrialEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
		// 				HostName: to.Ptr("mystreamingendpoint1-slitestmedia10.streaming.mediaservices.windows.net"),
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:09.202Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ResourceState: to.Ptr(armmediaservices.StreamingEndpointResourceStateStopped),
		// 				ScaleUnits: to.Ptr[int32](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/streamingEndpoints"),
		// 			ID: to.Ptr("/subscriptions/0a6ec948-5a62-437d-b9df-934dc7c1b722/resourceGroups/mediaresources/providers/Microsoft.Media/mediaservices/slitestmedia10/streamingendpoints/default"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armmediaservices.StreamingEndpointProperties{
		// 				Description: to.Ptr(""),
		// 				CdnEnabled: to.Ptr(true),
		// 				CdnProfile: to.Ptr("AzureMediaStreamingPlatformCdnProfile-StandardVerizon"),
		// 				CdnProvider: to.Ptr("StandardVerizon"),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:09.231Z"); return t}()),
		// 				CustomHostNames: []*string{
		// 				},
		// 				FreeTrialEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
		// 				HostName: to.Ptr("slitestmedia10.streaming.mediaservices.windows.net"),
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:09.231Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ResourceState: to.Ptr(armmediaservices.StreamingEndpointResourceStateStarting),
		// 				ScaleUnits: to.Ptr[int32](0),
		// 			},
		// 	}},
		// }
	}
}
