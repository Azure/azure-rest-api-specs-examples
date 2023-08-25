package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetReportsByTime.json
func ExampleReportsClient_NewListByTimePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReportsClient().NewListByTimePager("rg1", "apimService1", "timestamp ge datetime'2017-06-01T00:00:00' and timestamp le datetime'2017-06-04T00:00:00'", "PT15M", &armapimanagement.ReportsClientListByTimeOptions{Top: nil,
		Skip:    nil,
		Orderby: nil,
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
		// page.ReportCollection = armapimanagement.ReportCollection{
		// 	Count: to.Ptr[int64](2),
		// 	Value: []*armapimanagement.ReportRecordContract{
		// 		{
		// 			APITimeAvg: to.Ptr[float64](1337.46335),
		// 			APITimeMax: to.Ptr[float64](1819.2173),
		// 			APITimeMin: to.Ptr[float64](885.0839000000001),
		// 			Bandwidth: to.Ptr[int64](3243),
		// 			CacheHitCount: to.Ptr[int32](0),
		// 			CacheMissCount: to.Ptr[int32](0),
		// 			CallCountBlocked: to.Ptr[int32](0),
		// 			CallCountFailed: to.Ptr[int32](0),
		// 			CallCountOther: to.Ptr[int32](0),
		// 			CallCountSuccess: to.Ptr[int32](4),
		// 			CallCountTotal: to.Ptr[int32](4),
		// 			Interval: to.Ptr("PT15M"),
		// 			ServiceTimeAvg: to.Ptr[float64](1255.917425),
		// 			ServiceTimeMax: to.Ptr[float64](1697.3612),
		// 			ServiceTimeMin: to.Ptr[float64](882.8264),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-03T00:15:00Z"); return t}()),
		// 		},
		// 		{
		// 			APITimeAvg: to.Ptr[float64](872.7818777777778),
		// 			APITimeMax: to.Ptr[float64](1093.8407),
		// 			APITimeMin: to.Ptr[float64](330.3206),
		// 			Bandwidth: to.Ptr[int64](7776),
		// 			CacheHitCount: to.Ptr[int32](0),
		// 			CacheMissCount: to.Ptr[int32](0),
		// 			CallCountBlocked: to.Ptr[int32](1),
		// 			CallCountFailed: to.Ptr[int32](0),
		// 			CallCountOther: to.Ptr[int32](0),
		// 			CallCountSuccess: to.Ptr[int32](9),
		// 			CallCountTotal: to.Ptr[int32](10),
		// 			Interval: to.Ptr("PT15M"),
		// 			ServiceTimeAvg: to.Ptr[float64](824.2847111111112),
		// 			ServiceTimeMax: to.Ptr[float64](973.2262000000001),
		// 			ServiceTimeMin: to.Ptr[float64](215.24),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-03T00:30:00Z"); return t}()),
		// 	}},
		// }
	}
}
