package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetReportsByUser.json
func ExampleReportsClient_NewListByUserPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReportsClient().NewListByUserPager("rg1", "apimService1", "timestamp ge datetime'2017-06-01T00:00:00' and timestamp le datetime'2017-06-04T00:00:00'", &armapimanagement.ReportsClientListByUserOptions{Top: nil,
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
		// 	Count: to.Ptr[int64](3),
		// 	Value: []*armapimanagement.ReportRecordContract{
		// 		{
		// 			Name: to.Ptr("Administrator"),
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
		// 			ServiceTimeAvg: to.Ptr[float64](957.094776923077),
		// 			ServiceTimeMax: to.Ptr[float64](1697.3612),
		// 			ServiceTimeMin: to.Ptr[float64](215.24),
		// 			UserID: to.Ptr("/users/1"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Samir Solanki"),
		// 			APITimeAvg: to.Ptr[float64](0),
		// 			APITimeMax: to.Ptr[float64](0),
		// 			APITimeMin: to.Ptr[float64](0),
		// 			Bandwidth: to.Ptr[int64](0),
		// 			CacheHitCount: to.Ptr[int32](0),
		// 			CacheMissCount: to.Ptr[int32](0),
		// 			CallCountBlocked: to.Ptr[int32](0),
		// 			CallCountFailed: to.Ptr[int32](0),
		// 			CallCountOther: to.Ptr[int32](0),
		// 			CallCountSuccess: to.Ptr[int32](0),
		// 			CallCountTotal: to.Ptr[int32](0),
		// 			ServiceTimeAvg: to.Ptr[float64](0),
		// 			ServiceTimeMax: to.Ptr[float64](0),
		// 			ServiceTimeMin: to.Ptr[float64](0),
		// 			UserID: to.Ptr("/users/56eaec62baf08b06e46d27fd"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Anonymous"),
		// 			APITimeAvg: to.Ptr[float64](0),
		// 			APITimeMax: to.Ptr[float64](0),
		// 			APITimeMin: to.Ptr[float64](0),
		// 			Bandwidth: to.Ptr[int64](0),
		// 			CacheHitCount: to.Ptr[int32](0),
		// 			CacheMissCount: to.Ptr[int32](0),
		// 			CallCountBlocked: to.Ptr[int32](0),
		// 			CallCountFailed: to.Ptr[int32](0),
		// 			CallCountOther: to.Ptr[int32](0),
		// 			CallCountSuccess: to.Ptr[int32](0),
		// 			CallCountTotal: to.Ptr[int32](0),
		// 			ServiceTimeAvg: to.Ptr[float64](0),
		// 			ServiceTimeMax: to.Ptr[float64](0),
		// 			ServiceTimeMin: to.Ptr[float64](0),
		// 			UserID: to.Ptr("/users/54c800b332965a0035030000"),
		// 	}},
		// }
	}
}
