package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetReportsByGeo.json
func ExampleReportsClient_NewListByGeoPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReportsClient().NewListByGeoPager("rg1", "apimService1", "timestamp ge datetime'2017-06-01T00:00:00' and timestamp le datetime'2017-06-04T00:00:00'", &armapimanagement.ReportsClientListByGeoOptions{Top: nil,
		Skip: nil,
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
		// 	Value: []*armapimanagement.ReportRecordContract{
		// 		{
		// 			APITimeAvg: to.Ptr[float64](1015.7607923076923),
		// 			APITimeMax: to.Ptr[float64](1819.2173),
		// 			APITimeMin: to.Ptr[float64](330.3206),
		// 			Bandwidth: to.Ptr[int64](11019),
		// 			CacheHitCount: to.Ptr[int32](0),
		// 			CacheMissCount: to.Ptr[int32](0),
		// 			CallCountBlocked: to.Ptr[int32](1),
		// 			CallCountFailed: to.Ptr[int32](0),
		// 			CallCountOther: to.Ptr[int32](0),
		// 			CallCountSuccess: to.Ptr[int32](13),
		// 			CallCountTotal: to.Ptr[int32](14),
		// 			Country: to.Ptr("US"),
		// 			Region: to.Ptr("WA"),
		// 			ServiceTimeAvg: to.Ptr[float64](957.094776923077),
		// 			ServiceTimeMax: to.Ptr[float64](1697.3612),
		// 			ServiceTimeMin: to.Ptr[float64](215.24),
		// 			Zip: to.Ptr("98052"),
		// 	}},
		// }
	}
}
