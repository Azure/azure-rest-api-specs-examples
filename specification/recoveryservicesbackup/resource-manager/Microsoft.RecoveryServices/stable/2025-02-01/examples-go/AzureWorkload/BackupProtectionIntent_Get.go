package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureWorkload/BackupProtectionIntent_Get.json
func ExampleProtectionIntentClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectionIntentClient().Get(ctx, "myVault", "myRG", "Azure", "249D9B07-D2EF-4202-AA64-65F35418564E", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectionIntentResource = armrecoveryservicesbackup.ProtectionIntentResource{
	// 	Name: to.Ptr("249D9B07-D2EF-4202-AA64-65F35418564E"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupProtectionIntent"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.RecoveryServices/vaults/myVault/backupFabrics/Azure/backupProtectionIntent/249D9B07-D2EF-4202-AA64-65F35418564E"),
	// 	Properties: &armrecoveryservicesbackup.AzureWorkloadContainerAutoProtectionIntent{
	// 		BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureWorkload),
	// 		ItemID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.RecoveryServices/vaults/myVault/backupProtectionContainer/VMAppContainer;Compute;testVmName/backupProtectableItems/SQLInstance;MSSQLSERVER"),
	// 		PolicyID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.RecoveryServices/vaults/myVault/backupPolicies/myPolicy"),
	// 		ProtectionIntentItemType: to.Ptr(armrecoveryservicesbackup.ProtectionIntentItemTypeAzureWorkloadContainerAutoProtectionIntent),
	// 	},
	// }
}
