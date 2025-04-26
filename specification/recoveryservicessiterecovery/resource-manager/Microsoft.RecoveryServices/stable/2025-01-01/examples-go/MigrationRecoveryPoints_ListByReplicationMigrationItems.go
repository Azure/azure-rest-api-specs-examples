package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/MigrationRecoveryPoints_ListByReplicationMigrationItems.json
func ExampleMigrationRecoveryPointsClient_NewListByReplicationMigrationItemsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMigrationRecoveryPointsClient().NewListByReplicationMigrationItemsPager("resourcegroup1", "migrationvault", "vmwarefabric1", "vmwareContainer1", "virtualmachine1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.MigrationRecoveryPointCollection = armrecoveryservicessiterecovery.MigrationRecoveryPointCollection{
		// 	Value: []*armrecoveryservicessiterecovery.MigrationRecoveryPoint{
		// 		{
		// 			Name: to.Ptr("648336ef-2d70-4d98-b100-8c299f97cd41"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationMigrationItems/migrationRecoveryPoints"),
		// 			ID: to.Ptr("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.RecoveryServices/vaults/migrationvault/replicationFabrics/vmwarefabric1/replicationProtectionContainers/vmwareContainer1/replicationMigrationItems/virtualmachine1/migrationRecoveryPoints/648336ef-2d70-4d98-b100-8c299f97cd41"),
		// 			Properties: &armrecoveryservicessiterecovery.MigrationRecoveryPointProperties{
		// 				RecoveryPointTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-26T06:37:50.808Z"); return t}()),
		// 				RecoveryPointType: to.Ptr(armrecoveryservicessiterecovery.MigrationRecoveryPointTypeCrashConsistent),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("34cb2d05-e730-4d3f-b96b-a60a5e92acb2"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationMigrationItems/migrationRecoveryPoints"),
		// 			ID: to.Ptr("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.RecoveryServices/vaults/migrationvault/replicationFabrics/vmwarefabric1/replicationProtectionContainers/vmwareContainer1/replicationMigrationItems/virtualmachine1/migrationRecoveryPoints/34cb2d05-e730-4d3f-b96b-a60a5e92acb2"),
		// 			Properties: &armrecoveryservicessiterecovery.MigrationRecoveryPointProperties{
		// 				RecoveryPointTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-26T07:37:30.972Z"); return t}()),
		// 				RecoveryPointType: to.Ptr(armrecoveryservicessiterecovery.MigrationRecoveryPointTypeCrashConsistent),
		// 			},
		// 	}},
		// }
	}
}
