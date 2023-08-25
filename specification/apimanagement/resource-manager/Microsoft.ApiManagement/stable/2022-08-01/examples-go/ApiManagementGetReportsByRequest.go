package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetReportsByRequest.json
func ExampleReportsClient_NewListByRequestPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReportsClient().NewListByRequestPager("rg1", "apimService1", "timestamp ge datetime'2017-06-01T00:00:00' and timestamp le datetime'2017-06-04T00:00:00'", &armapimanagement.ReportsClientListByRequestOptions{Top: nil,
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
		// page.RequestReportCollection = armapimanagement.RequestReportCollection{
		// 	Count: to.Ptr[int64](2),
		// 	Value: []*armapimanagement.RequestReportRecordContract{
		// 		{
		// 			OperationID: to.Ptr("/apis/5931a75ae4bbd512a88c680b/operations/-"),
		// 			Method: to.Ptr("GET"),
		// 			APIID: to.Ptr("/apis/5931a75ae4bbd512a88c680b"),
		// 			APIRegion: to.Ptr("East Asia"),
		// 			APITime: to.Ptr[float64](221.1544),
		// 			Cache: to.Ptr("none"),
		// 			IPAddress: to.Ptr("207.xx.155.xx"),
		// 			ProductID: to.Ptr("/products/-"),
		// 			RequestID: to.Ptr("63e7119c-26aa-433c-96d7-f6f3267ff52f"),
		// 			RequestSize: to.Ptr[int32](0),
		// 			ResponseCode: to.Ptr[int32](404),
		// 			ResponseSize: to.Ptr[int32](405),
		// 			ServiceTime: to.Ptr[float64](0),
		// 			SubscriptionID: to.Ptr("/subscriptions/5600b59475ff190048070002"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-03T00:17:00.1649134Z"); return t}()),
		// 			URL: to.Ptr("https://apimService1.azure-api.net/echo/resource?param1=sample"),
		// 			UserID: to.Ptr("/users/1"),
		// 		},
		// 		{
		// 			OperationID: to.Ptr("/apis/5931a75ae4bbd512a88c680b/operations/-"),
		// 			Method: to.Ptr("POST"),
		// 			APIID: to.Ptr("/apis/5931a75ae4bbd512a88c680b"),
		// 			APIRegion: to.Ptr("East Asia"),
		// 			APITime: to.Ptr[float64](6.675400000000001),
		// 			Cache: to.Ptr("none"),
		// 			IPAddress: to.Ptr("207.xx.155.xx"),
		// 			ProductID: to.Ptr("/products/-"),
		// 			RequestID: to.Ptr("e581b7f7-c9ec-4fc6-8ab9-3855d9b00b04"),
		// 			RequestSize: to.Ptr[int32](0),
		// 			ResponseCode: to.Ptr[int32](404),
		// 			ResponseSize: to.Ptr[int32](403),
		// 			ServiceTime: to.Ptr[float64](0),
		// 			SubscriptionID: to.Ptr("/subscriptions/5600b59475ff190048070002"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-03T00:17:20.5255131Z"); return t}()),
		// 			URL: to.Ptr("https://apimService1.azure-api.net/echo/resource"),
		// 			UserID: to.Ptr("/users/1"),
		// 	}},
		// }
	}
}
