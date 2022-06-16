package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationMigrationItems_TestMigrate.json
func ExampleReplicationMigrationItemsClient_BeginTestMigrate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginTestMigrate(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.TestMigrateInput{
			Properties: &armrecoveryservicessiterecovery.TestMigrateInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtTestMigrateInput{
					InstanceType:    to.StringPtr("<instance-type>"),
					NetworkID:       to.StringPtr("<network-id>"),
					RecoveryPointID: to.StringPtr("<recovery-point-id>"),
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
	log.Printf("Response result: %#v\n", res.ReplicationMigrationItemsClientTestMigrateResult)
}
