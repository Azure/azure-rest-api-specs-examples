package armattestation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/attestation/armattestation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Delete_AttestationProvider.json
func ExampleProvidersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armattestation.NewProvidersClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"sample-resource-group",
		"myattestationprovider",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
