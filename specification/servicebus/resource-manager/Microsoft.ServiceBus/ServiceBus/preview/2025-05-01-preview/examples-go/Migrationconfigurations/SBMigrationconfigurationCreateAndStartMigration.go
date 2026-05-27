package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: 2025-05-01-preview/Migrationconfigurations/SBMigrationconfigurationCreateAndStartMigration.json
func ExampleMigrationConfigsClient_BeginCreateAndStartMigration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMigrationConfigsClient().BeginCreateAndStartMigration(ctx, "ResourceGroup", "sdk-Namespace-41", armservicebus.MigrationConfigurationNameDefault, armservicebus.MigrationConfigProperties{
		Properties: &armservicebus.MigrationConfigPropertiesProperties{
			PostMigrationName: to.Ptr("sdk-PostMigration-5919"),
			TargetNamespace:   to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-4028"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicebus.MigrationConfigsClientCreateAndStartMigrationResponse{
	// 	MigrationConfigProperties: armservicebus.MigrationConfigProperties{
	// 		Name: to.Ptr("sdk-Namespace-41"),
	// 		Type: to.Ptr("Microsoft.ServiceBus/Namespaces/disasterrecoveryconfigs"),
	// 		ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-41/migrationConfigs/$default"),
	// 		Properties: &armservicebus.MigrationConfigPropertiesProperties{
	// 			MigrationState: to.Ptr("Initiating"),
	// 			PostMigrationName: to.Ptr("sdk-PostMigration-5919"),
	// 			ProvisioningState: to.Ptr("Accepted"),
	// 			TargetNamespace: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-4028"),
	// 		},
	// 	},
	// }
}
