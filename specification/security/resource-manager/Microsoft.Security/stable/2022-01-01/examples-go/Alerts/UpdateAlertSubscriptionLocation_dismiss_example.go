package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2022-01-01/examples/Alerts/UpdateAlertSubscriptionLocation_dismiss_example.json
func ExampleAlertsClient_UpdateSubscriptionLevelStateToDismiss() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAlertsClient().UpdateSubscriptionLevelStateToDismiss(ctx, "westeurope", "2518298467986649999_4d25bfef-2d77-4a08-adc0-3e35715cc92a", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
