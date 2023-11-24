package armsphere_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sphere/armsphere"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ed9bde6a3db71b84fdba076ba0546213bcce56ee/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostListDeviceInsightsCatalog.json
func ExampleCatalogsClient_NewListDeviceInsightsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCatalogsClient().NewListDeviceInsightsPager("MyResourceGroup1", "MyCatalog1", &armsphere.CatalogsClientListDeviceInsightsOptions{Filter: nil,
		Top:         to.Ptr[int32](10),
		Skip:        nil,
		Maxpagesize: nil,
	})
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
		// page.PagedDeviceInsight = armsphere.PagedDeviceInsight{
		// 	Value: []*armsphere.DeviceInsight{
		// 		{
		// 			Description: to.Ptr("eventDescription1"),
		// 			DeviceID: to.Ptr("eventIdentifier1"),
		// 			EndTimestampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-30T23:54:21.960Z"); return t}()),
		// 			EventCategory: to.Ptr("eventCategory1"),
		// 			EventClass: to.Ptr("eventClass1"),
		// 			EventCount: to.Ptr[int32](1),
		// 			EventType: to.Ptr("eventType1"),
		// 			StartTimestampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-30T21:51:39.260Z"); return t}()),
		// 		},
		// 		{
		// 			Description: to.Ptr("eventDescription2"),
		// 			DeviceID: to.Ptr("eventIdentifier2"),
		// 			EndTimestampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-07T17:34:12.500Z"); return t}()),
		// 			EventCategory: to.Ptr("eventCategory2"),
		// 			EventClass: to.Ptr("eventClass2"),
		// 			EventCount: to.Ptr[int32](1),
		// 			EventType: to.Ptr("eventType2"),
		// 			StartTimestampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-06T12:41:39.260Z"); return t}()),
		// 	}},
		// }
	}
}
