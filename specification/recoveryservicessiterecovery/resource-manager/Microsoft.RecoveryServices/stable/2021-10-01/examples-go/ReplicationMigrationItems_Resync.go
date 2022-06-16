package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationMigrationItems_Resync.json
func ExampleReplicationMigrationItemsClient_BeginResync() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginResync(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.ResyncInput{
			Properties: &armrecoveryservicessiterecovery.ResyncInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtResyncInput{
					ResyncProviderSpecificInput: armrecoveryservicessiterecovery.ResyncProviderSpecificInput{
						InstanceType: to.StringPtr("<instance-type>"),
					},
					SkipCbtReset: to.StringPtr("<skip-cbt-reset>"),
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
