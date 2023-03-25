package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/preview/2021-05-01-preview/examples/getDiagnosticSetting.json
func ExampleDiagnosticSettingsClient_Get_getsTheDiagnosticSetting() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiagnosticSettingsClient().Get(ctx, "subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6", "mysetting", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiagnosticSettingsResource = armmonitor.DiagnosticSettingsResource{
	// 	Name: to.Ptr("mysetting"),
	// 	Type: to.Ptr("Microsoft.Insights/diagnosticSettings"),
	// 	ID: to.Ptr("/subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6/providers/microsoft.insights/diagnosticSettings/mysetting"),
	// 	Properties: &armmonitor.DiagnosticSettings{
	// 		EventHubAuthorizationRuleID: to.Ptr("/subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourceGroups/montest/providers/microsoft.eventhub/namespaces/mynamespace/authorizationrules/myrule"),
	// 		Logs: []*armmonitor.LogSettings{
	// 			{
	// 				CategoryGroup: to.Ptr("allLogs"),
	// 				Enabled: to.Ptr(true),
	// 				RetentionPolicy: &armmonitor.RetentionPolicy{
	// 					Days: to.Ptr[int32](0),
	// 					Enabled: to.Ptr(false),
	// 				},
	// 		}},
	// 		MarketplacePartnerID: to.Ptr("/subscriptions/abcdeabc-1234-1234-ab12-123a1234567a/resourceGroups/test-rg/providers/Microsoft.Datadog/monitors/dd1"),
	// 		Metrics: []*armmonitor.MetricSettings{
	// 			{
	// 				Category: to.Ptr("WorkflowMetrics"),
	// 				Enabled: to.Ptr(true),
	// 				RetentionPolicy: &armmonitor.RetentionPolicy{
	// 					Days: to.Ptr[int32](0),
	// 					Enabled: to.Ptr(false),
	// 				},
	// 		}},
	// 		StorageAccountID: to.Ptr("/subscriptions/df602c9c-7aa0-407d-a6fb-eb20c8bd1192/resourceGroups/apptest/providers/Microsoft.Storage/storageAccounts/appteststorage1"),
	// 		WorkspaceID: to.Ptr(""),
	// 	},
	// }
}
