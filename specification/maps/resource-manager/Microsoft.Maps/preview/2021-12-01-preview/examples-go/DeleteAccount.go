package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/DeleteAccount.json
func ExampleAccountsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmaps.NewAccountsClient("21a9967a-e8a9-4656-a70b-96ff1c4d05a0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"myResourceGroup",
		"myMapsAccount",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
