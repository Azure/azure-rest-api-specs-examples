package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/stable/2022-01-01/examples/Alerts/UpdateAlertResourceGroupLocation_activate_example.json
func ExampleAlertsClient_UpdateResourceGroupLevelStateToActivate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAlertsClient().UpdateResourceGroupLevelStateToActivate(ctx, "myRg2", "westeurope", "2518765996949954086_2325cf9e-42a2-4f72-ae7f-9b863cba2d22", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
