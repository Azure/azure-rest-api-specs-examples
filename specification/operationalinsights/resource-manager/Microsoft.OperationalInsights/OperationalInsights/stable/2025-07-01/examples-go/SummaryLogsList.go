package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f0a5127d4e8b1ea6007b0bf9570904d9df860b97/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/SummaryLogsList.json
func ExampleSummaryLogsClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSummaryLogsClient().NewListByWorkspacePager("oiautorest6685", "oiautorest6685", nil)
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
		// page.SummaryLogsListResult = armoperationalinsights.SummaryLogsListResult{
		// 	Value: []*armoperationalinsights.SummaryLogs{
		// 		{
		// 			Name: to.Ptr("summarylogs1"),
		// 			ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/oiautorest6685/providers/Microsoft.OperationalInsights/workspaces/oiautorest6685/summaryLogs/summarylogs1"),
		// 			SystemData: &armoperationalinsights.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-03T04:05:06.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("22222222-2222-2222-2222-222222222222"),
		// 			},
		// 			Properties: &armoperationalinsights.SummaryLogsProperties{
		// 				IsActive: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armoperationalinsights.ProvisioningStateEnumSucceeded),
		// 				RuleDefinition: &armoperationalinsights.RuleDefinition{
		// 					BinDelay: to.Ptr[int32](10),
		// 					BinSize: to.Ptr[int32](180),
		// 					BinStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T04:05:06.000Z"); return t}()),
		// 					DestinationTable: to.Ptr("MyDestinationTable_CL"),
		// 					Query: to.Ptr("MyTable_CL"),
		// 				},
		// 				RuleType: to.Ptr(armoperationalinsights.RuleTypeEnumUser),
		// 				StatusCode: to.Ptr(armoperationalinsights.StatusCodeEnumUserAction),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("summarylogs2"),
		// 			ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/oiautorest6685/providers/Microsoft.OperationalInsights/workspaces/oiautorest6685/summaryLogs/summarylogs2"),
		// 			SystemData: &armoperationalinsights.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-03T04:05:06.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("22222222-2222-2222-2222-222222222222"),
		// 			},
		// 			Properties: &armoperationalinsights.SummaryLogsProperties{
		// 				IsActive: to.Ptr(true),
		// 				ProvisioningState: to.Ptr(armoperationalinsights.ProvisioningStateEnumSucceeded),
		// 				RuleDefinition: &armoperationalinsights.RuleDefinition{
		// 					BinDelay: to.Ptr[int32](60),
		// 					BinSize: to.Ptr[int32](720),
		// 					DestinationTable: to.Ptr("MyDestinationTable2_CL"),
		// 					Query: to.Ptr("MyTable2_CL"),
		// 				},
		// 				RuleType: to.Ptr(armoperationalinsights.RuleTypeEnumUser),
		// 			},
		// 	}},
		// }
	}
}
