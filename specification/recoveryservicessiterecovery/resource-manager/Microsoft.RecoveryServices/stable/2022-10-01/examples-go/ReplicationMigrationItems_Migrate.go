package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationMigrationItems_Migrate.json
func ExampleReplicationMigrationItemsClient_BeginMigrate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("migrationvault", "resourcegroup1", "cb53d0c3-bd59-4721-89bc-06916a9147ef", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginMigrate(ctx, "vmwarefabric1", "vmwareContainer1", "virtualmachine1", armrecoveryservicessiterecovery.MigrateInput{
		Properties: &armrecoveryservicessiterecovery.MigrateInputProperties{
			ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtMigrateInput{
				InstanceType:    to.Ptr("VMwareCbt"),
				PerformShutdown: to.Ptr("true"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
