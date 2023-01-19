package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3751704f5318f1175875c94b66af769db917f2d3/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-01-01/examples/AzureIaasVm/ProtectionIntent_CreateOrUpdate.json
func ExampleProtectionIntentClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewProtectionIntentClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myVault", "myRG", "Azure", "vm;iaasvmcontainerv2;chamsrgtest;chamscandel", armrecoveryservicesbackup.ProtectionIntentResource{
		Properties: &armrecoveryservicesbackup.AzureResourceProtectionIntent{
			PolicyID:                 to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.RecoveryServices/vaults/myVault/backupPolicies/myPolicy"),
			ProtectionIntentItemType: to.Ptr(armrecoveryservicesbackup.ProtectionIntentItemTypeAzureResourceItem),
			SourceResourceID:         to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/chamsrgtest/providers/Microsoft.Compute/virtualMachines/chamscandel"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectionIntentResource = armrecoveryservicesbackup.ProtectionIntentResource{
	// 	Name: to.Ptr("vm;iaasvmcontainerv2;chamsrgtest;chamscandel"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupProtectionIntent"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.RecoveryServices/vaults/myVault/backupFabrics/Azure/backupProtectionIntent/vm;iaasvmcontainerv2;chamsrgtest;chamscandel"),
	// 	Properties: &armrecoveryservicesbackup.AzureResourceProtectionIntent{
	// 		BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureIaasVM),
	// 		PolicyID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.RecoveryServices/vaults/myVault/backupPolicies/myPolicy"),
	// 		ProtectionIntentItemType: to.Ptr(armrecoveryservicesbackup.ProtectionIntentItemTypeAzureResourceItem),
	// 		ProtectionState: to.Ptr(armrecoveryservicesbackup.ProtectionStatusProtected),
	// 	},
	// }
}
