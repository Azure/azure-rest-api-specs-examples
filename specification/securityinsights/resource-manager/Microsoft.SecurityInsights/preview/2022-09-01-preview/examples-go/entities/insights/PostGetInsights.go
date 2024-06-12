package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/insights/PostGetInsights.json
func ExampleEntitiesClient_GetInsights() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().GetInsights(ctx, "myRg", "myWorkspace", "e1d3d618-e11f-478b-98e3-bb381539a8e1", armsecurityinsights.EntityGetInsightsParameters{
		AddDefaultExtendedTimeRange: to.Ptr(false),
		EndTime:                     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-01T00:00:00.000Z"); return t }()),
		InsightQueryIDs: []*string{
			to.Ptr("cae8d0aa-aa45-4d53-8d88-17dd64ffd4e4")},
		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T00:00:00.000Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EntityGetInsightsResponse = armsecurityinsights.EntityGetInsightsResponse{
	// 	MetaData: &armsecurityinsights.GetInsightsResultsMetadata{
	// 		Errors: []*armsecurityinsights.GetInsightsErrorKind{
	// 			{
	// 				ErrorMessage: to.Ptr("Internal server error"),
	// 				Kind: to.Ptr(armsecurityinsights.GetInsightsErrorInsight),
	// 				QueryID: to.Ptr("4a70a63d-25c4-6312-b73e-4f302a90c06a"),
	// 		}},
	// 		TotalCount: to.Ptr[int32](7),
	// 	},
	// 	Value: []*armsecurityinsights.EntityInsightItem{
	// 		{
	// 			ChartQueryResults: []*armsecurityinsights.InsightsTableResult{
	// 				{
	// 					Columns: []*armsecurityinsights.InsightsTableResultColumnsItem{
	// 						{
	// 							Name: to.Ptr("TimeGenerated"),
	// 							Type: to.Ptr("datetime"),
	// 						},
	// 						{
	// 							Name: to.Ptr("Count"),
	// 							Type: to.Ptr("long"),
	// 						},
	// 						{
	// 							Name: to.Ptr("Legend"),
	// 							Type: to.Ptr("string"),
	// 					}},
	// 					Rows: [][]*string{
	// 						[]*string{
	// 							to.Ptr("2021-09-01T00:00:00.000Z"),
	// 							to.Ptr("55"),
	// 							to.Ptr("SomeLegend")}},
	// 					}},
	// 					QueryID: to.Ptr("e29ee1ef-7445-455e-85f1-269f2d536d61"),
	// 					QueryTimeInterval: &armsecurityinsights.EntityInsightItemQueryTimeInterval{
	// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T23:35:20.000Z"); return t}()),
	// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T23:35:20.000Z"); return t}()),
	// 					},
	// 					TableQueryResults: &armsecurityinsights.InsightsTableResult{
	// 						Columns: []*armsecurityinsights.InsightsTableResultColumnsItem{
	// 							{
	// 								Name: to.Ptr("Title"),
	// 								Type: to.Ptr("string"),
	// 							},
	// 							{
	// 								Name: to.Ptr("NameCount"),
	// 								Type: to.Ptr("long"),
	// 							},
	// 							{
	// 								Name: to.Ptr("SIDCount"),
	// 								Type: to.Ptr("long"),
	// 							},
	// 							{
	// 								Name: to.Ptr("InternalOrder"),
	// 								Type: to.Ptr("long"),
	// 							},
	// 							{
	// 								Name: to.Ptr("Index"),
	// 								Type: to.Ptr("long"),
	// 						}},
	// 						Rows: [][]*string{
	// 							[]*string{
	// 								to.Ptr("MyTitle"),
	// 								to.Ptr("15"),
	// 								to.Ptr("SID"),
	// 								to.Ptr("1"),
	// 								to.Ptr("1")}},
	// 							},
	// 					}},
	// 				}
}
