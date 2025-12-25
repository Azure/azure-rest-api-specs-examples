package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be46becafeb29aa993898709e35759d3643b2809/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/QueryPackQueriesList.json
func ExampleQueriesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewQueriesClient().NewListPager("my-resource-group", "my-querypack", &armoperationalinsights.QueriesClientListOptions{Top: nil,
		IncludeBody: to.Ptr(true),
		SkipToken:   nil,
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
		// page.LogAnalyticsQueryPackQueryListResult = armoperationalinsights.LogAnalyticsQueryPackQueryListResult{
		// 	Value: []*armoperationalinsights.LogAnalyticsQueryPackQuery{
		// 		{
		// 			Name: to.Ptr("4337bb16-d6fe-4ff7-97cf-59df25941476"),
		// 			Type: to.Ptr("microsoft.operationalinsights/queryPacks/queries"),
		// 			ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4918/resourceGroups/my-resource-group/providers/microsoft.operationalinsights/queryPacks/my-querypack/queries/4337bb16-d6fe-4ff7-97cf-59df25941476"),
		// 			SystemData: &armoperationalinsights.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armoperationalinsights.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armoperationalinsights.CreatedByTypeApplication),
		// 			},
		// 			Properties: &armoperationalinsights.LogAnalyticsQueryPackQueryProperties{
		// 				Description: to.Ptr("Thie query takes 10 entries of heartbeat"),
		// 				Author: to.Ptr("1809f206-263a-46f7-942d-4572c156b7e7"),
		// 				Body: to.Ptr("heartbeat | take 10"),
		// 				DisplayName: to.Ptr("Heartbeat_1"),
		// 				ID: to.Ptr("4337bb16-d6fe-4ff7-97cf-59df25941476"),
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:29:56.103Z"); return t}()),
		// 				TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:29:56.103Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("bf015bf7-be70-49c2-8d52-4cce85c42ef1"),
		// 			Type: to.Ptr("microsoft.operationalinsights/queryPacks/queries"),
		// 			ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4918/resourceGroups/my-resource-group/providers/microsoft.operationalinsights/queryPacks/my-querypack/queries/bf015bf7-be70-49c2-8d52-4cce85c42ef1"),
		// 			SystemData: &armoperationalinsights.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armoperationalinsights.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armoperationalinsights.CreatedByTypeApplication),
		// 			},
		// 			Properties: &armoperationalinsights.LogAnalyticsQueryPackQueryProperties{
		// 				Description: to.Ptr("Thie query takes 10 entries of heartbeat"),
		// 				Author: to.Ptr("1809f206-263a-46f7-942d-4572c156b7e7"),
		// 				Body: to.Ptr("heartbeat | take 10"),
		// 				DisplayName: to.Ptr("Heartbeat_2"),
		// 				ID: to.Ptr("bf015bf7-be70-49c2-8d52-4cce85c42ef1"),
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:30:26.794Z"); return t}()),
		// 				TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:30:26.794Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("8d91c6ca-9c56-49c6-b3ae-112a68871acd"),
		// 			Type: to.Ptr("microsoft.operationalinsights/queryPacks/queries"),
		// 			ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4918/resourceGroups/my-resource-group/providers/microsoft.operationalinsights/queryPacks/my-querypack/queries/8d91c6ca-9c56-49c6-b3ae-112a68871acd"),
		// 			SystemData: &armoperationalinsights.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armoperationalinsights.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armoperationalinsights.CreatedByTypeApplication),
		// 			},
		// 			Properties: &armoperationalinsights.LogAnalyticsQueryPackQueryProperties{
		// 				Description: to.Ptr("Thie query takes 10 entries of heartbeat"),
		// 				Author: to.Ptr("1809f206-263a-46f7-942d-4572c156b7e7"),
		// 				Body: to.Ptr("heartbeat | take 10"),
		// 				DisplayName: to.Ptr("Heartbeat_3"),
		// 				ID: to.Ptr("8d91c6ca-9c56-49c6-b3ae-112a68871acd"),
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:30:29.450Z"); return t}()),
		// 				TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:30:29.450Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("a5a9010e-e4b7-45ad-8b14-09d7e6082819"),
		// 			Type: to.Ptr("microsoft.operationalinsights/queryPacks/queries"),
		// 			ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4918/resourceGroups/my-resource-group/providers/microsoft.operationalinsights/queryPacks/my-querypack/queries/a5a9010e-e4b7-45ad-8b14-09d7e6082819"),
		// 			SystemData: &armoperationalinsights.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armoperationalinsights.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armoperationalinsights.CreatedByTypeApplication),
		// 			},
		// 			Properties: &armoperationalinsights.LogAnalyticsQueryPackQueryProperties{
		// 				Description: to.Ptr("Thie query takes 10 entries of heartbeat"),
		// 				Author: to.Ptr("1809f206-263a-46f7-942d-4572c156b7e7"),
		// 				Body: to.Ptr("heartbeat | take 10"),
		// 				DisplayName: to.Ptr("Heartbeat_4"),
		// 				ID: to.Ptr("a5a9010e-e4b7-45ad-8b14-09d7e6082819"),
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:30:32.574Z"); return t}()),
		// 				TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:30:32.574Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
