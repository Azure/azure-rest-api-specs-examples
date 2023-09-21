package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/MigrationRecoveryPoints_Get.json
func ExampleMigrationRecoveryPointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMigrationRecoveryPointsClient().Get(ctx, "migrationvault", "resourcegroup1", "vmwarefabric1", "vmwareContainer1", "virtualmachine1", "b22134ea-620c-474b-9fa5-3c1cb47708e3", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MigrationRecoveryPoint = armrecoveryservicessiterecovery.MigrationRecoveryPoint{
	// 	Name: to.Ptr("b22134ea-620c-474b-9fa5-3c1cb47708e3"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationMigrationItems/migrationRecoveryPoints"),
	// 	ID: to.Ptr("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.RecoveryServices/vaults/migrationvault/replicationFabrics/vmwarefabric1/replicationProtectionContainers/vmwareContainer1/replicationMigrationItems/virtualmachine1/migrationRecoveryPoints/b22134ea-620c-474b-9fa5-3c1cb47708e3"),
	// 	Properties: &armrecoveryservicessiterecovery.MigrationRecoveryPointProperties{
	// 		RecoveryPointTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-26T06:37:50.8082715Z"); return t}()),
	// 		RecoveryPointType: to.Ptr(armrecoveryservicessiterecovery.MigrationRecoveryPointTypeCrashConsistent),
	// 	},
	// }
}
