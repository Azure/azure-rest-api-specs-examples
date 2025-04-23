package armneonpostgres_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/neonpostgres/armneonpostgres"
)

// Generated from example definition: 2025-03-01/Computes_Delete_MaximumSet_Gen.json
func ExampleComputesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("9B8E3300-C5FA-442B-A259-3F6F614D5BD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComputesClient().Delete(ctx, "rgneon", "test-org", "entity-name", "entity-name", "entity-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armneonpostgres.ComputesClientDeleteResponse{
	// }
}
