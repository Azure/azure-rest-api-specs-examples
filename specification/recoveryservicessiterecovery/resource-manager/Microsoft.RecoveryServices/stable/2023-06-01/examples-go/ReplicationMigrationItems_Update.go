package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationMigrationItems_Update.json
func ExampleReplicationMigrationItemsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationMigrationItemsClient().BeginUpdate(ctx, "migrationvault", "resourcegroup1", "vmwarefabric1", "vmwareContainer1", "virtualmachine1", armrecoveryservicessiterecovery.UpdateMigrationItemInput{
		Properties: &armrecoveryservicessiterecovery.UpdateMigrationItemInputProperties{
			ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtUpdateMigrationItemInput{
				InstanceType: to.Ptr("VMwareCbt"),
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MigrationItem = armrecoveryservicessiterecovery.MigrationItem{
	// 	Name: to.Ptr("virtualmachine1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationMigrationItems"),
	// 	ID: to.Ptr("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.RecoveryServices/vaults/migrationvault/replicationFabrics/vmwarefabric1/replicationProtectionContainers/vmwareContainer1/replicationMigrationItems/virtualmachine1"),
	// 	Properties: &armrecoveryservicessiterecovery.MigrationItemProperties{
	// 		AllowedOperations: []*armrecoveryservicessiterecovery.MigrationItemOperation{
	// 			to.Ptr(armrecoveryservicessiterecovery.MigrationItemOperationMigrate),
	// 			to.Ptr(armrecoveryservicessiterecovery.MigrationItemOperationDisableMigration),
	// 			to.Ptr(armrecoveryservicessiterecovery.MigrationItemOperationTestMigrate),
	// 			to.Ptr(armrecoveryservicessiterecovery.MigrationItemOperationTestMigrateCleanup)},
	// 			CurrentJob: &armrecoveryservicessiterecovery.CurrentJobDetails{
	// 				JobID: to.Ptr("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.RecoveryServices/vaults/migrationvault/replicationJobs/None"),
	// 				JobName: to.Ptr("None"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-26T06:37:50.808Z"); return t}()),
	// 			},
	// 			MachineName: to.Ptr("vm-0520-2"),
	// 			MigrationState: to.Ptr(armrecoveryservicessiterecovery.MigrationStateReplicating),
	// 			MigrationStateDescription: to.Ptr("Ready to migrate"),
	// 			PolicyFriendlyName: to.Ptr("vmwarepolicy1"),
	// 			PolicyID: to.Ptr("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.RecoveryServices/vaults/migrationvault/replicationPolicies/vmwarepolicy1"),
	// 			ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtMigrationDetails{
	// 				InstanceType: to.Ptr("VMwareCbt"),
	// 			},
	// 			TestMigrateState: to.Ptr(armrecoveryservicessiterecovery.TestMigrationStateNone),
	// 			TestMigrateStateDescription: to.Ptr("None"),
	// 		},
	// 	}
}
