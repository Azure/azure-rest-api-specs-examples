package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/HyperVCollectors_Delete.json
func ExampleHyperVCollectorsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmigrate.NewHyperVCollectorsClient("8c3c936a-c09b-4de3-830b-3f5f244d72e9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"contosoithyperv",
		"migrateprojectce73project",
		"migrateprojectce73collector",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
