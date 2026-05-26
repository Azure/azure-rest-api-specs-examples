package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2022-01-01/Alerts/UpdateAlertResourceGroupLocation_inProgress_example.json
func ExampleAlertsClient_UpdateResourceGroupLevelStateToInProgress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAlertsClient().UpdateResourceGroupLevelStateToInProgress(ctx, "myRg2", "westeurope", "2518765996949954086_2325cf9e-42a2-4f72-ae7f-9b863cba2d22", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
