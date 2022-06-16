package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Migrationconfigurations/SBMigrationconfigurationCompleteMigration.json
func ExampleMigrationConfigsClient_CompleteMigration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewMigrationConfigsClient("SubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.CompleteMigration(ctx,
		"ResourceGroup",
		"sdk-Namespace-41",
		armservicebus.MigrationConfigurationNameDefault,
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
