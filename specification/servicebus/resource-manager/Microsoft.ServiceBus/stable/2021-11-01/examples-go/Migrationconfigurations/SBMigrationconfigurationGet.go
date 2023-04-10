package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Migrationconfigurations/SBMigrationconfigurationGet.json
func ExampleMigrationConfigsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMigrationConfigsClient().Get(ctx, "ResourceGroup", "sdk-Namespace-41", armservicebus.MigrationConfigurationNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MigrationConfigProperties = armservicebus.MigrationConfigProperties{
	// 	Name: to.Ptr("sdk-Namespace-41"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/disasterrecoveryconfigs"),
	// 	ID: to.Ptr("/subscriptions/SubscriptionId/resourceGroups/ResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-41/migrationConfigs/$default"),
	// 	Properties: &armservicebus.MigrationConfigPropertiesProperties{
	// 		MigrationState: to.Ptr("Active"),
	// 		PendingReplicationOperationsCount: to.Ptr[int64](0),
	// 		PostMigrationName: to.Ptr("sdk-PostMigration-5919"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		TargetNamespace: to.Ptr("/subscriptions/SubscriptionId/resourceGroups/ResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-4028"),
	// 	},
	// }
}
