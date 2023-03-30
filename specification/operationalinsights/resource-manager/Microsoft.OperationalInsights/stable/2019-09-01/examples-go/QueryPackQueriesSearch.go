package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf3574813e15bb33b3cb610f44edfcbebd8b1b23/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPackQueriesSearch.json
func ExampleQueriesClient_NewSearchPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewQueriesClient().NewSearchPager("my-resource-group", "my-querypack", armoperationalinsights.LogAnalyticsQueryPackQuerySearchProperties{
		Related: &armoperationalinsights.LogAnalyticsQueryPackQuerySearchPropertiesRelated{
			Categories: []*string{
				to.Ptr("other"),
				to.Ptr("analytics")},
		},
		Tags: map[string][]*string{
			"my-label": {
				to.Ptr("label1")},
		},
	}, &armoperationalinsights.QueriesClientSearchOptions{Top: to.Ptr[int64](3),
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
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.1075056Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armoperationalinsights.IdentityTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.1974346Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armoperationalinsights.IdentityTypeApplication),
		// 			},
		// 			Properties: &armoperationalinsights.LogAnalyticsQueryPackQueryProperties{
		// 				Description: to.Ptr("Thie query takes 10 entries of heartbeat 0"),
		// 				Author: to.Ptr("1809f206-263a-46f7-942d-4572c156b7e7"),
		// 				Body: to.Ptr("Heartbeat | take 1"),
		// 				DisplayName: to.Ptr("Heartbeat_1"),
		// 				ID: to.Ptr("4337bb16-d6fe-4ff7-97cf-59df25941476"),
		// 				Related: &armoperationalinsights.LogAnalyticsQueryPackQueryPropertiesRelated{
		// 					Categories: []*string{
		// 						to.Ptr("other")},
		// 					},
		// 					Tags: map[string][]*string{
		// 						"my-label": []*string{
		// 							to.Ptr("label1")},
		// 							"my-other-label": []*string{
		// 								to.Ptr("label2")},
		// 							},
		// 							TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:29:56.1030254Z"); return t}()),
		// 							TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:29:56.1030254Z"); return t}()),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("bf015bf7-be70-49c2-8d52-4cce85c42ef1"),
		// 						Type: to.Ptr("microsoft.operationalinsights/queryPacks/queries"),
		// 						ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4918/resourceGroups/my-resource-group/providers/microsoft.operationalinsights/queryPacks/my-querypack/queries/bf015bf7-be70-49c2-8d52-4cce85c42ef1"),
		// 						SystemData: &armoperationalinsights.SystemData{
		// 							CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.1075056Z"); return t}()),
		// 							CreatedBy: to.Ptr("string"),
		// 							CreatedByType: to.Ptr(armoperationalinsights.IdentityTypeApplication),
		// 							LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.1974346Z"); return t}()),
		// 							LastModifiedBy: to.Ptr("string"),
		// 							LastModifiedByType: to.Ptr(armoperationalinsights.IdentityTypeApplication),
		// 						},
		// 						Properties: &armoperationalinsights.LogAnalyticsQueryPackQueryProperties{
		// 							Description: to.Ptr("Thie query takes 10 entries of heartbeat 1"),
		// 							Author: to.Ptr("1809f206-263a-46f7-942d-4572c156b7e7"),
		// 							Body: to.Ptr("Heartbeat | take 1"),
		// 							DisplayName: to.Ptr("Heartbeat_2"),
		// 							ID: to.Ptr("bf015bf7-be70-49c2-8d52-4cce85c42ef1"),
		// 							Related: &armoperationalinsights.LogAnalyticsQueryPackQueryPropertiesRelated{
		// 								Categories: []*string{
		// 									to.Ptr("analytics")},
		// 								},
		// 								Tags: map[string][]*string{
		// 									"my-label": []*string{
		// 										to.Ptr("label1")},
		// 										"my-other-label": []*string{
		// 											to.Ptr("label2")},
		// 										},
		// 										TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:30:26.7943629Z"); return t}()),
		// 										TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:30:26.7943629Z"); return t}()),
		// 									},
		// 								},
		// 								{
		// 									Name: to.Ptr("8d91c6ca-9c56-49c6-b3ae-112a68871acd"),
		// 									Type: to.Ptr("microsoft.operationalinsights/queryPacks/queries"),
		// 									ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4918/resourceGroups/my-resource-group/providers/microsoft.operationalinsights/queryPacks/my-querypack/queries/8d91c6ca-9c56-49c6-b3ae-112a68871acd"),
		// 									SystemData: &armoperationalinsights.SystemData{
		// 										CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.1075056Z"); return t}()),
		// 										CreatedBy: to.Ptr("string"),
		// 										CreatedByType: to.Ptr(armoperationalinsights.IdentityTypeApplication),
		// 										LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.1974346Z"); return t}()),
		// 										LastModifiedBy: to.Ptr("string"),
		// 										LastModifiedByType: to.Ptr(armoperationalinsights.IdentityTypeApplication),
		// 									},
		// 									Properties: &armoperationalinsights.LogAnalyticsQueryPackQueryProperties{
		// 										Description: to.Ptr("Thie query takes 10 entries of heartbeat 2"),
		// 										Author: to.Ptr("1809f206-263a-46f7-942d-4572c156b7e7"),
		// 										Body: to.Ptr("Heartbeat | take 1"),
		// 										DisplayName: to.Ptr("Heartbeat_3"),
		// 										ID: to.Ptr("8d91c6ca-9c56-49c6-b3ae-112a68871acd"),
		// 										Related: &armoperationalinsights.LogAnalyticsQueryPackQueryPropertiesRelated{
		// 											Categories: []*string{
		// 												to.Ptr("other"),
		// 												to.Ptr("analytics")},
		// 											},
		// 											Tags: map[string][]*string{
		// 												"my-label": []*string{
		// 													to.Ptr("label1")},
		// 													"my-other-label": []*string{
		// 														to.Ptr("label2")},
		// 													},
		// 													TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:30:29.4505584Z"); return t}()),
		// 													TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:30:29.4505584Z"); return t}()),
		// 												},
		// 										}},
		// 									}
	}
}
