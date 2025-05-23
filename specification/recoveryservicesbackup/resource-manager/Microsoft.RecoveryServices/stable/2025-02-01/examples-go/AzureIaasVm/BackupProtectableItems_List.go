package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureIaasVm/BackupProtectableItems_List.json
func ExampleBackupProtectableItemsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupProtectableItemsClient().NewListPager("NetSDKTestRsVault", "SwaggerTestRg", &armrecoveryservicesbackup.BackupProtectableItemsClientListOptions{Filter: to.Ptr("backupManagementType eq 'AzureIaasVM'"),
		SkipToken: nil,
	})
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
		// page.WorkloadProtectableItemResourceList = armrecoveryservicesbackup.WorkloadProtectableItemResourceList{
		// 	Value: []*armrecoveryservicesbackup.WorkloadProtectableItemResource{
		// 		{
		// 			Name: to.Ptr("VM;iaasvmcontainer;iaasvm-rg;iaasvm-1"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/protectableItems"),
		// 			ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/NetSDKTestRsVault/protectionContainers/IaasVMContainer;iaasvmcontainer;iaasvm-rg;iaasvm-1/protectableItems/VM;iaasvmcontainer;iaasvm-rg;iaasvm-1"),
		// 			Properties: &armrecoveryservicesbackup.AzureIaaSClassicComputeVMProtectableItem{
		// 				BackupManagementType: to.Ptr("AzureIaasVM"),
		// 				FriendlyName: to.Ptr("iaasvm-1"),
		// 				ProtectableItemType: to.Ptr("Microsoft.ClassicCompute/virtualMachines"),
		// 				ProtectionState: to.Ptr(armrecoveryservicesbackup.ProtectionStatusNotProtected),
		// 				WorkloadType: to.Ptr("VM"),
		// 				VirtualMachineID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/providers/Microsoft.ClassicCompute/virtualMachines/iaasvm-1"),
		// 			},
		// 	}},
		// }
	}
}
