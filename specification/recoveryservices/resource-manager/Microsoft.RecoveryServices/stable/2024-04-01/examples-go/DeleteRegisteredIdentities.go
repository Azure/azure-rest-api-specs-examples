package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/DeleteRegisteredIdentities.json
func ExampleRegisteredIdentitiesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewRegisteredIdentitiesClient().Delete(ctx, "BCDRIbzRG", "BCDRIbzVault", "dpmcontainer01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
