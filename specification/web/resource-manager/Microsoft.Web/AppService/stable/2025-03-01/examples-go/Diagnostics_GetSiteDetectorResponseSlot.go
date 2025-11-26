package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/Diagnostics_GetSiteDetectorResponseSlot.json
func ExampleDiagnosticsClient_GetSiteDetectorResponseSlot_getAppSlotDetectorResponse() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiagnosticsClient().GetSiteDetectorResponseSlot(ctx, "Sample-WestUSResourceGroup", "SampleApp", "runtimeavailability", "staging", &armappservice.DiagnosticsClientGetSiteDetectorResponseSlotOptions{StartTime: nil,
		EndTime:   nil,
		TimeGrain: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DetectorResponse = armappservice.DetectorResponse{
	// 	Name: to.Ptr("runtimeavailability"),
	// 	ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/slots/staging/detectors/runtimeavailability"),
	// 	Properties: &armappservice.DetectorResponseProperties{
	// 		Dataset: []*armappservice.DiagnosticData{
	// 			{
	// 				RenderingProperties: &armappservice.Rendering{
	// 					Description: to.Ptr("This detector breaks down the number of requests that your app received for each status code."),
	// 					Title: to.Ptr("Requests by Status Code"),
	// 				},
	// 				Table: &armappservice.DataTableResponseObject{
	// 					Columns: []*armappservice.DataTableResponseColumn{
	// 						{
	// 							ColumnName: to.Ptr("PreciseTimeStamp"),
	// 							ColumnType: to.Ptr("datetime"),
	// 							DataType: to.Ptr("DateTime"),
	// 						},
	// 						{
	// 							ColumnName: to.Ptr("count_Http2xx"),
	// 							ColumnType: to.Ptr("long"),
	// 							DataType: to.Ptr("Int64"),
	// 						},
	// 						{
	// 							ColumnName: to.Ptr("count_Http3xx"),
	// 							ColumnType: to.Ptr("long"),
	// 							DataType: to.Ptr("Int64"),
	// 						},
	// 						{
	// 							ColumnName: to.Ptr("count_Http4xx"),
	// 							ColumnType: to.Ptr("long"),
	// 							DataType: to.Ptr("Int64"),
	// 						},
	// 						{
	// 							ColumnName: to.Ptr("count_Http5xx"),
	// 							ColumnType: to.Ptr("long"),
	// 							DataType: to.Ptr("Int64"),
	// 					}},
	// 					Rows: [][]*string{
	// 						[]*string{
	// 							to.Ptr("2018-03-27T00:25:00Z"),
	// 							to.Ptr("772705"),
	// 							to.Ptr("0"),
	// 							to.Ptr("0"),
	// 							to.Ptr("0")},
	// 							[]*string{
	// 								to.Ptr("2018-03-27T00:30:00Z"),
	// 								to.Ptr("787069"),
	// 								to.Ptr("0"),
	// 								to.Ptr("0"),
	// 								to.Ptr("0")},
	// 								[]*string{
	// 									to.Ptr("2018-03-27T00:35:00Z"),
	// 									to.Ptr("781627"),
	// 									to.Ptr("0"),
	// 									to.Ptr("1"),
	// 									to.Ptr("0")},
	// 									[]*string{
	// 										to.Ptr("2018-03-27T00:40:00Z"),
	// 										to.Ptr("785017"),
	// 										to.Ptr("0"),
	// 										to.Ptr("0"),
	// 										to.Ptr("0")},
	// 										[]*string{
	// 											to.Ptr("2018-03-27T00:45:00Z"),
	// 											to.Ptr("783518"),
	// 											to.Ptr("0"),
	// 											to.Ptr("0"),
	// 											to.Ptr("0")},
	// 											[]*string{
	// 												to.Ptr("2018-03-27T00:50:00Z"),
	// 												to.Ptr("785783"),
	// 												to.Ptr("0"),
	// 												to.Ptr("0"),
	// 												to.Ptr("0")},
	// 												[]*string{
	// 													to.Ptr("2018-03-27T00:55:00Z"),
	// 													to.Ptr("772874"),
	// 													to.Ptr("0"),
	// 													to.Ptr("0"),
	// 													to.Ptr("0")},
	// 													[]*string{
	// 														to.Ptr("2018-03-27T01:00:00Z"),
	// 														to.Ptr("787162"),
	// 														to.Ptr("0"),
	// 														to.Ptr("0"),
	// 														to.Ptr("0")},
	// 														[]*string{
	// 															to.Ptr("2018-03-27T01:05:00Z"),
	// 															to.Ptr("782036"),
	// 															to.Ptr("0"),
	// 															to.Ptr("0"),
	// 															to.Ptr("0")},
	// 															[]*string{
	// 																to.Ptr("2018-03-27T01:10:00Z"),
	// 																to.Ptr("784642"),
	// 																to.Ptr("0"),
	// 																to.Ptr("0"),
	// 																to.Ptr("0")}},
	// 																TableName: to.Ptr("Table_0"),
	// 															},
	// 													}},
	// 													Metadata: &armappservice.DetectorInfo{
	// 														Description: to.Ptr("This detector analyzes the requests to your application."),
	// 														Category: to.Ptr("Availability and Performance"),
	// 													},
	// 												},
	// 											}
}
