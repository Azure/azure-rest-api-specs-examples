package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2021-11-01/examples/Alerts/UpdateAlertSubscriptionLocation_resolve_example.json
func ExampleAlertsClient_UpdateSubscriptionLevelStateToResolve() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAlertsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.UpdateSubscriptionLevelStateToResolve(ctx,
		"westeurope",
		"2518298467986649999_4d25bfef-2d77-4a08-adc0-3e35715cc92a",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
