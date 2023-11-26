package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveoutput-create.json
func ExampleLiveOutputsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLiveOutputsClient().BeginCreate(ctx, "mediaresources", "slitestmedia10", "myLiveEvent1", "myLiveOutput1", armmediaservices.LiveOutput{
		Properties: &armmediaservices.LiveOutputProperties{
			Description:         to.Ptr("test live output 1"),
			ArchiveWindowLength: to.Ptr("PT5M"),
			AssetName:           to.Ptr("6f3264f5-a189-48b4-a29a-a40f22575212"),
			Hls: &armmediaservices.Hls{
				FragmentsPerTsSegment: to.Ptr[int32](5),
			},
			ManifestName:       to.Ptr("testmanifest"),
			RewindWindowLength: to.Ptr("PT4M"),
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
	// res.LiveOutput = armmediaservices.LiveOutput{
	// 	Name: to.Ptr("myLiveOutput1"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/liveevents/liveoutputs"),
	// 	ID: to.Ptr("/subscriptions/0a6ec948-5a62-437d-b9df-934dc7c1b722/resourceGroups/mediaresources/providers/Microsoft.Media/mediaservices/slitestmedia10/liveevents/myLiveEvent1/liveoutputs/myLiveOutput1"),
	// 	Properties: &armmediaservices.LiveOutputProperties{
	// 		Description: to.Ptr("test live output 1"),
	// 		ArchiveWindowLength: to.Ptr("PT5M"),
	// 		AssetName: to.Ptr("6f3264f5-a189-48b4-a29a-a40f22575212"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:09.943Z"); return t}()),
	// 		Hls: &armmediaservices.Hls{
	// 			FragmentsPerTsSegment: to.Ptr[int32](5),
	// 		},
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:09.943Z"); return t}()),
	// 		ManifestName: to.Ptr("testmanifest"),
	// 		OutputSnapTime: to.Ptr[int64](0),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ResourceState: to.Ptr(armmediaservices.LiveOutputResourceState("Stopped")),
	// 		RewindWindowLength: to.Ptr("PT4M"),
	// 	},
	// 	SystemData: &armmediaservices.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:09.943Z"); return t}()),
	// 		CreatedBy: to.Ptr("example@microsoft.com"),
	// 		CreatedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:09.943Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("example@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
	// 	},
	// }
}
