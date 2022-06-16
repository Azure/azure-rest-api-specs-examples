package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/preview/2017-05-01-preview/examples/createOrUpdateDiagnosticSetting.json
func ExampleDiagnosticSettingsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewDiagnosticSettingsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6",
		"mysetting",
		armmonitor.DiagnosticSettingsResource{
			Properties: &armmonitor.DiagnosticSettings{
				EventHubAuthorizationRuleID: to.Ptr("/subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourceGroups/montest/providers/microsoft.eventhub/namespaces/mynamespace/eventhubs/myeventhub/authorizationrules/myrule"),
				EventHubName:                to.Ptr("myeventhub"),
				LogAnalyticsDestinationType: to.Ptr("Dedicated"),
				Logs: []*armmonitor.LogSettings{
					{
						Category: to.Ptr("WorkflowRuntime"),
						Enabled:  to.Ptr(true),
						RetentionPolicy: &armmonitor.RetentionPolicy{
							Days:    to.Ptr[int32](0),
							Enabled: to.Ptr(false),
						},
					}},
				Metrics: []*armmonitor.MetricSettings{
					{
						Category: to.Ptr("WorkflowMetrics"),
						Enabled:  to.Ptr(true),
						RetentionPolicy: &armmonitor.RetentionPolicy{
							Days:    to.Ptr[int32](0),
							Enabled: to.Ptr(false),
						},
					}},
				StorageAccountID: to.Ptr("/subscriptions/df602c9c-7aa0-407d-a6fb-eb20c8bd1192/resourceGroups/apptest/providers/Microsoft.Storage/storageAccounts/appteststorage1"),
				WorkspaceID:      to.Ptr(""),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
