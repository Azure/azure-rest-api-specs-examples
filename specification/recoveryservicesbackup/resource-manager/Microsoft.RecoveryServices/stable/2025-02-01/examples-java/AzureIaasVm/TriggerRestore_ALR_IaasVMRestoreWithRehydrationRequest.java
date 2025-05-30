
import com.azure.resourcemanager.recoveryservicesbackup.models.EncryptionDetails;
import com.azure.resourcemanager.recoveryservicesbackup.models.IaasVMRestoreWithRehydrationRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryPointRehydrationInfo;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryType;
import com.azure.resourcemanager.recoveryservicesbackup.models.RehydrationPriority;
import com.azure.resourcemanager.recoveryservicesbackup.models.RestoreRequestResource;

/**
 * Samples for Restores Trigger.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/
     * AzureIaasVm/TriggerRestore_ALR_IaasVMRestoreWithRehydrationRequest.json
     */
    /**
     * Sample code: Restore to New Azure IaasVm with IaasVMRestoreWithRehydrationRequest.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void restoreToNewAzureIaasVmWithIaasVMRestoreWithRehydrationRequest(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.restores().trigger("testVault", "netsdktestrg", "Azure",
            "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
            "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "348916168024334",
            new RestoreRequestResource().withProperties(new IaasVMRestoreWithRehydrationRequest()
                .withRecoveryPointId("348916168024334").withRecoveryType(RecoveryType.ALTERNATE_LOCATION)
                .withSourceResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1")
                .withTargetVirtualMachineId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg2/providers/Microsoft.Compute/virtualmachines/RSMDALRVM981435")
                .withTargetResourceGroupId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg2")
                .withStorageAccountId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Storage/storageAccounts/testingAccount")
                .withVirtualNetworkId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Network/virtualNetworks/testNet")
                .withSubnetId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Network/virtualNetworks/testNet/subnets/default")
                .withRegion("southeastasia").withCreateNewCloudService(false).withOriginalStorageAccountOption(false)
                .withEncryptionDetails(new EncryptionDetails().withEncryptionEnabled(false))
                .withRecoveryPointRehydrationInfo(new RecoveryPointRehydrationInfo()
                    .withRehydrationRetentionDuration("P7D").withRehydrationPriority(RehydrationPriority.HIGH))),
            com.azure.core.util.Context.NONE);
    }
}
