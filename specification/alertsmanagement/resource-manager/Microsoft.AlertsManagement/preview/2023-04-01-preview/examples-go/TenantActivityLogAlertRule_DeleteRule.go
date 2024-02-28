package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6d2438481021a94793b07b226df06d5f3c61d51d/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2023-04-01-preview/examples/TenantActivityLogAlertRule_DeleteRule.json
func ExampleTenantActivityLogAlertsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTenantActivityLogAlertsClient().Delete(ctx, "72f988bf-86f1-41af-91ab-2d7cd011db47", "SampleActivityLogAlertSHRuleOnTenantLevel", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
