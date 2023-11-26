package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveoutput-list-all.json
func ExampleLiveOutputsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLiveOutputsClient().NewListPager("mediaresources", "slitestmedia10", "myLiveEvent1", nil)
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
		// page.LiveOutputListResult = armmediaservices.LiveOutputListResult{
		// 	Value: []*armmediaservices.LiveOutput{
		// 		{
		// 			Name: to.Ptr("liveoutput1"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/liveevents/liveoutputs"),
		// 			ID: to.Ptr("/subscriptions/0a6ec948-5a62-437d-b9df-934dc7c1b722/resourceGroups/mediaresources/providers/Microsoft.Media/mediaservices/slitestmedia10/liveevents/myLiveEvent1/liveoutputs/"),
		// 			Properties: &armmediaservices.LiveOutputProperties{
		// 				ArchiveWindowLength: to.Ptr("PT5M"),
		// 				AssetName: to.Ptr("95dafce4-5320-464c-8597-909373854119"),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
		// 				Hls: &armmediaservices.Hls{
		// 					FragmentsPerTsSegment: to.Ptr[int32](5),
		// 				},
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
		// 				ManifestName: to.Ptr("c3a23d4b-02a6-4937-a1ad-6416f463fdca"),
		// 				OutputSnapTime: to.Ptr[int64](0),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ResourceState: to.Ptr(armmediaservices.LiveOutputResourceStateRunning),
		// 				RewindWindowLength: to.Ptr("PT4M"),
		// 			},
		// 			SystemData: &armmediaservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("example@microsoft.com"),
		// 				CreatedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("example@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
