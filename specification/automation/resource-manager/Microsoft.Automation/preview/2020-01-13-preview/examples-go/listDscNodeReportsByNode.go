package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listDscNodeReportsByNode.json
func ExampleNodeReportsClient_NewListByNodePager_listDscReportsByNodeId() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNodeReportsClient().NewListByNodePager("rg", "myAutomationAccount33", "nodeId", &armautomation.NodeReportsClientListByNodeOptions{Filter: nil})
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
		// page.DscNodeReportListResult = armautomation.DscNodeReportListResult{
		// 	Value: []*armautomation.DscNodeReport{
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:16:27.587Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/903a5ead-140c-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:16:29.444Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("903a5ead-140c-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:16:27.587Z"); return t}()),
		// 			Status: to.Ptr("Compliant"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:16:27.015Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/903a5eac-140c-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:16:28.381Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("903a5eac-140c-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:16:26.015Z"); return t}()),
		// 			Status: to.Ptr("Compliant"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:01:26.986Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/77c280c2-140a-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:01:28.216Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("77c280c2-140a-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:01:25.986Z"); return t}()),
		// 			Status: to.Ptr("Compliant"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:46:28.668Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/5f4f5382-1408-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:46:29.043Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("5f4f5382-1408-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:46:27.668Z"); return t}()),
		// 			Status: to.Ptr("Compliant"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:46:26.957Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/5f4f5381-1408-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:46:27.949Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("5f4f5381-1408-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:46:25.957Z"); return t}()),
		// 			Status: to.Ptr("Compliant"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:31:26.941Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/46d97d6a-1406-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:31:27.682Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("46d97d6a-1406-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:31:25.941Z"); return t}()),
		// 			Status: to.Ptr("Compliant"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:17:10.163Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/48c8e301-1404-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:17:11.004Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("48c8e301-1404-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:17:10.163Z"); return t}()),
		// 			Status: to.Ptr("Failed"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:17:09.897Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/2e63fdbc-1404-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:16:27.312Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("2e63fdbc-1404-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:16:25.897Z"); return t}()),
		// 			Status: to.Ptr("Compliant"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:01:27.899Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/15ee63e4-1402-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:01:26.628Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("15ee63e4-1402-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:01:25.899Z"); return t}()),
		// 			Status: to.Ptr("Compliant"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:46:39.511Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/0508f316-1400-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:46:40.577Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("0508f316-1400-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:46:38.511Z"); return t}()),
		// 			Status: to.Ptr("Failed"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:46:37.843Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/fd799a51-13ff-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:46:28.466Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("fd799a51-13ff-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:46:25.843Z"); return t}()),
		// 			Status: to.Ptr("Compliant"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:31:27.818Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/e504ae1b-13fd-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:31:28.150Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("e504ae1b-13fd-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:31:25.818Z"); return t}()),
		// 			Status: to.Ptr("Compliant"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:16:59.538Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/cd3ed224-13fb-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:16:37.176Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("cd3ed224-13fb-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:16:35.538Z"); return t}()),
		// 			Status: to.Ptr("Compliant"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:16:34.956Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/cd3ed223-13fb-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:16:29.237Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("cd3ed223-13fb-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:16:26.956Z"); return t}()),
		// 			Status: to.Ptr("Compliant"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:02:02.916Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/b6915efa-13f9-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:01:31.875Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("b6915efa-13f9-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:01:29.916Z"); return t}()),
		// 			Status: to.Ptr("Compliant"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T20:46:44.626Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/a3560dca-13f7-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T20:46:39.749Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("a3560dca-13f7-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T20:46:38.626Z"); return t}()),
		// 			Status: to.Ptr("Failed"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T20:46:37.676Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/9b9bb016-13f7-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T20:46:27.260Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("9b9bb016-13f7-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T20:46:25.676Z"); return t}()),
		// 			Status: to.Ptr("Compliant"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Consistency"),
		// 			ConfigurationVersion: to.Ptr("2.0.0"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T20:40:24.805Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/833bd89b-13f5-11e7-a943-000d3a6140c9"),
		// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T20:31:27.135Z"); return t}()),
		// 			RebootRequested: to.Ptr("False"),
		// 			RefreshMode: to.Ptr("Pull"),
		// 			ReportFormatVersion: to.Ptr("2.0"),
		// 			ReportID: to.Ptr("833bd89b-13f5-11e7-a943-000d3a6140c9"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T20:31:25.805Z"); return t}()),
		// 			Status: to.Ptr("Compliant"),
		// 	}},
		// }
	}
}
