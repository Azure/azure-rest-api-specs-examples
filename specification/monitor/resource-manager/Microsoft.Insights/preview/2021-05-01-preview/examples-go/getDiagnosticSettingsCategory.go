package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/preview/2021-05-01-preview/examples/getDiagnosticSettingsCategory.json
func ExampleDiagnosticSettingsCategoryClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiagnosticSettingsCategoryClient().Get(ctx, "subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6", "WorkflowRuntime", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiagnosticSettingsCategoryResource = armmonitor.DiagnosticSettingsCategoryResource{
	// 	Name: to.Ptr("WorkflowRuntime"),
	// 	Type: to.Ptr("microsoft.insights/diagnosticSettingsCategories"),
	// 	ID: to.Ptr("/subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6/providers/microsoft.insights/diagnosticSettingsCategories/WorkflowRuntime"),
	// 	Properties: &armmonitor.DiagnosticSettingsCategory{
	// 		CategoryGroups: []*string{
	// 			to.Ptr("allLogs")},
	// 			CategoryType: to.Ptr(armmonitor.CategoryTypeLogs),
	// 		},
	// 	}
}
