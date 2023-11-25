package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/streamingendpoint-operation-location.json
func ExampleStreamingEndpointsClient_OperationLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStreamingEndpointsClient().OperationLocation(ctx, "mediaresources", "slitestmedia10", "myStreamingEndpoint1", "62e4d893-d233-4005-988e-a428d9f77076", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StreamingEndpoint = armmediaservices.StreamingEndpoint{
	// 	Name: to.Ptr("myStreamingEndpoint1"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/streamingEndpoints"),
	// 	ID: to.Ptr("/subscriptions/0a6ec948-5a62-437d-b9df-934dc7c1b722/resourceGroups/mediaresources/providers/Microsoft.Media/mediaservices/slitestmedia10/streamingendpoints/myStreamingEndpoint1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armmediaservices.StreamingEndpointProperties{
	// 		Description: to.Ptr("test event 1"),
	// 		AvailabilitySetName: to.Ptr("availableset"),
	// 		CdnEnabled: to.Ptr(false),
	// 		CdnProfile: to.Ptr(""),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:09.350Z"); return t}()),
	// 		CustomHostNames: []*string{
	// 		},
	// 		FreeTrialEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
	// 		HostName: to.Ptr("mystreamingendpoint1-slitestmedia10.streaming.mediaservices.windows.net"),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:09.350Z"); return t}()),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ResourceState: to.Ptr(armmediaservices.StreamingEndpointResourceStateStopped),
	// 		ScaleUnits: to.Ptr[int32](1),
	// 	},
	// 	SKU: &armmediaservices.ArmStreamingEndpointCurrentSKU{
	// 		Name: to.Ptr("Premium"),
	// 		Capacity: to.Ptr[int32](1),
	// 	},
	// 	SystemData: &armmediaservices.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:09.350Z"); return t}()),
	// 		CreatedBy: to.Ptr("example@microsoft.com"),
	// 		CreatedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:09.350Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("example@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
	// 	},
	// }
}
