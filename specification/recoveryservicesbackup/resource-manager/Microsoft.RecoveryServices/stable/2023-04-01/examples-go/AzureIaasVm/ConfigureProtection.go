package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a4ddec441435d1ef766c4f160eda658a69cc5dc2/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/AzureIaasVm/ConfigureProtection.json
func ExampleProtectedItemsClient_CreateOrUpdate_enableProtectionOnAzureIaasVm() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectedItemsClient().CreateOrUpdate(ctx, "NetSDKTestRsVault", "SwaggerTestRg", "Azure", "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", armrecoveryservicesbackup.ProtectedItemResource{
		Properties: &armrecoveryservicesbackup.AzureIaaSComputeVMProtectedItem{
			PolicyID:          to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/NetSDKTestRsVault/backupPolicies/DefaultPolicy"),
			ProtectedItemType: to.Ptr("Microsoft.Compute/virtualMachines"),
			SourceResourceID:  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectedItemResource = armrecoveryservicesbackup.ProtectedItemResource{
	// 	Name: to.Ptr("VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/protectedItems"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/PythonSDKBackupTestRg/providers/Microsoft.RecoveryServices/vaults/PySDKBackupTestRsVault/backupFabrics/Azure/protectionContainers/IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1/protectedItems/VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
	// 	Properties: &armrecoveryservicesbackup.AzureIaaSComputeVMProtectedItem{
	// 		BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureIaasVM),
	// 		ContainerName: to.Ptr("iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
	// 		PolicyID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/PythonSDKBackupTestRg/providers/Microsoft.RecoveryServices/vaults/PySDKBackupTestRsVault/backupPolicies/testPolicy1"),
	// 		ProtectedItemType: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 		SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
	// 		WorkloadType: to.Ptr(armrecoveryservicesbackup.DataSourceTypeVM),
	// 		FriendlyName: to.Ptr("netvmtestv2vm1"),
	// 		HealthStatus: to.Ptr(armrecoveryservicesbackup.HealthStatusPassed),
	// 		LastBackupStatus: to.Ptr("Completed"),
	// 		LastBackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-22T12:25:32.048Z"); return t}()),
	// 		ProtectedItemDataID: to.Ptr("636482643132986882"),
	// 		ProtectionState: to.Ptr(armrecoveryservicesbackup.ProtectionStateProtected),
	// 		ProtectionStatus: to.Ptr("Healthy"),
	// 		VirtualMachineID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
	// 	},
	// }
}
