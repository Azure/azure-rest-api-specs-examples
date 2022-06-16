package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationMigrationItems_Migrate.json
func ExampleReplicationMigrationItemsClient_BeginMigrate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginMigrate(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.MigrateInput{
			Properties: &armrecoveryservicessiterecovery.MigrateInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtMigrateInput{
					MigrateProviderSpecificInput: armrecoveryservicessiterecovery.MigrateProviderSpecificInput{
						InstanceType: to.StringPtr("<instance-type>"),
					},
					PerformShutdown: to.StringPtr("<perform-shutdown>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("MigrationItem.ID: %s\n", *res.ID)
}
