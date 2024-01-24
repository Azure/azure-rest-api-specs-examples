package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d402f685809d6d08be9c0b45065cadd7d78ab870/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/AzureIaasVm/Compute_ProtectedItem_Get.json
func ExampleProtectedItemsClient_Get_getProtectedVirtualMachineDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectedItemsClient().Get(ctx, "PySDKBackupTestRsVault", "PythonSDKBackupTestRg", "Azure", "iaasvmcontainer;iaasvmcontainerv2;iaasvm-rg;iaasvm-1", "vm;iaasvmcontainerv2;iaasvm-rg;iaasvm-1", &armrecoveryservicesbackup.ProtectedItemsClientGetOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectedItemResource = armrecoveryservicesbackup.ProtectedItemResource{
	// 	Name: to.Ptr("VM;iaasvmcontainerv2;iaasvm-rg;iaasvm-1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/protectedItems"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/PythonSDKBackupTestRg/providers/Microsoft.RecoveryServices/vaults/PySDKBackupTestRsVault/backupFabrics/Azure/protectionContainers/IaasVMContainer;iaasvmcontainerv2;iaasvm-rg;iaasvm-1/protectedItems/VM;iaasvmcontainerv2;iaasvm-rg;iaasvm-1"),
	// 	Properties: &armrecoveryservicesbackup.AzureIaaSComputeVMProtectedItem{
	// 		BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureIaasVM),
	// 		ContainerName: to.Ptr("iaasvmcontainerv2;iaasvm-rg;iaasvm-1"),
	// 		LastRecoveryPoint: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-22T12:25:32.048Z"); return t}()),
	// 		PolicyID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/PythonSDKBackupTestRg/providers/Microsoft.RecoveryServices/vaults/PySDKBackupTestRsVault/backupPolicies/testPolicy1"),
	// 		ProtectedItemType: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 		SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/providers/Microsoft.Compute/virtualMachines/iaasvm-1"),
	// 		WorkloadType: to.Ptr(armrecoveryservicesbackup.DataSourceTypeVM),
	// 		FriendlyName: to.Ptr("iaasvm-1"),
	// 		HealthStatus: to.Ptr(armrecoveryservicesbackup.HealthStatusPassed),
	// 		LastBackupStatus: to.Ptr("Completed"),
	// 		LastBackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-22T12:25:32.048Z"); return t}()),
	// 		ProtectedItemDataID: to.Ptr("636482643132986882"),
	// 		ProtectionState: to.Ptr(armrecoveryservicesbackup.ProtectionStateProtected),
	// 		ProtectionStatus: to.Ptr("Healthy"),
	// 		VirtualMachineID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/providers/Microsoft.Compute/virtualMachines/iaasvm-1"),
	// 	},
	// }
}
