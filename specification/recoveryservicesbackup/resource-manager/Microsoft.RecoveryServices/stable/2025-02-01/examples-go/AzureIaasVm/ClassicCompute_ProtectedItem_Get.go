package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureIaasVm/ClassicCompute_ProtectedItem_Get.json
func ExampleProtectedItemsClient_Get_getProtectedClassicVirtualMachineDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectedItemsClient().Get(ctx, "PySDKBackupTestRsVault", "PythonSDKBackupTestRg", "Azure", "iaasvmcontainer;iaasvmcontainer;iaasvm-rg;iaasvm-1", "vm;iaasvmcontainer;iaasvm-rg;iaasvm-1", &armrecoveryservicesbackup.ProtectedItemsClientGetOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectedItemResource = armrecoveryservicesbackup.ProtectedItemResource{
	// 	Name: to.Ptr("VM;iaasvmcontainer;iaasvm-rg;iaasvm-1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/protectedItems"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/PythonSDKBackupTestRg/providers/Microsoft.RecoveryServices/vaults/PySDKBackupTestRsVault/backupFabrics/Azure/protectionContainers/IaasVMContainer;iaasvmcontainer;iaasvm-rg;iaasvm-1/protectedItems/VM;iaasvmcontainer;iaasvm-rg;iaasvm-1"),
	// 	Properties: &armrecoveryservicesbackup.AzureIaaSClassicComputeVMProtectedItem{
	// 		BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureIaasVM),
	// 		ContainerName: to.Ptr("iaasvmcontainer;iaasvm-rg;iaasvm-1"),
	// 		LastRecoveryPoint: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-22T12:25:32.048Z"); return t}()),
	// 		PolicyID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/PythonSDKBackupTestRg/providers/Microsoft.RecoveryServices/vaults/PySDKBackupTestRsVault/backupPolicies/testPolicy1"),
	// 		ProtectedItemType: to.Ptr("Microsoft.ClassicCompute/virtualMachines"),
	// 		SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/providers/Microsoft.ClassicCompute/virtualMachines/iaasvm-1"),
	// 		WorkloadType: to.Ptr(armrecoveryservicesbackup.DataSourceTypeVM),
	// 		FriendlyName: to.Ptr("iaasvm-1"),
	// 		HealthStatus: to.Ptr(armrecoveryservicesbackup.HealthStatusPassed),
	// 		LastBackupStatus: to.Ptr("Completed"),
	// 		LastBackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-22T12:25:32.048Z"); return t}()),
	// 		PolicyType: to.Ptr("V1"),
	// 		ProtectedItemDataID: to.Ptr("636482643132986882"),
	// 		ProtectionState: to.Ptr(armrecoveryservicesbackup.ProtectionStateProtected),
	// 		ProtectionStatus: to.Ptr("Healthy"),
	// 		VirtualMachineID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/providers/Microsoft.ClassicCompute/virtualMachines/iaasvm-1"),
	// 	},
	// }
}
