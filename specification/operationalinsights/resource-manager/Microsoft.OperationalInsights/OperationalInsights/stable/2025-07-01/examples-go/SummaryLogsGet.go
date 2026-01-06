package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53d56e4ec74156c450d1e51745a971d3f2031dd7/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/SummaryLogsGet.json
func ExampleSummaryLogsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSummaryLogsClient().Get(ctx, "oiautorest6685", "oiautorest6685", "summarylogs1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SummaryLogs = armoperationalinsights.SummaryLogs{
	// 	Name: to.Ptr("summarylogs1"),
	// 	ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/oiautorest6685/providers/Microsoft.OperationalInsights/workspaces/oiautorest6685/summaryLogs/summarylogs1"),
	// 	SystemData: &armoperationalinsights.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T04:05:06.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 	},
	// 	Properties: &armoperationalinsights.SummaryLogsProperties{
	// 		IsActive: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armoperationalinsights.ProvisioningStateEnumSucceeded),
	// 		RuleDefinition: &armoperationalinsights.RuleDefinition{
	// 			BinDelay: to.Ptr[int32](10),
	// 			BinSize: to.Ptr[int32](180),
	// 			BinStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T04:05:06.000Z"); return t}()),
	// 			DestinationTable: to.Ptr("MyDestinationTable_CL"),
	// 			Query: to.Ptr("MyTable_CL"),
	// 		},
	// 		RuleType: to.Ptr(armoperationalinsights.RuleTypeEnumUser),
	// 		StatusCode: to.Ptr(armoperationalinsights.StatusCodeEnumUserAction),
	// 	},
	// }
}
