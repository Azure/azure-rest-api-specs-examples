package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2021-11-01/examples/Alerts/UpdateAlertResourceGroupLocation_example.json
func ExampleAlertsClient_UpdateResourceGroupLevelStateToActivate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAlertsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.UpdateResourceGroupLevelStateToActivate(ctx,
		"westeurope",
		"2518765996949954086_2325cf9e-42a2-4f72-ae7f-9b863cba2d22",
		"myRg2",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
