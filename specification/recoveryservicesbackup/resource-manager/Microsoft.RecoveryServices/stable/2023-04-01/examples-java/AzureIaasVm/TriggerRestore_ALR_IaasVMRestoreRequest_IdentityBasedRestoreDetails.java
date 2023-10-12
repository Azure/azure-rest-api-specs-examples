import com.azure.resourcemanager.recoveryservicesbackup.models.EncryptionDetails;
import com.azure.resourcemanager.recoveryservicesbackup.models.IaasVMRestoreRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.IdentityBasedRestoreDetails;
import com.azure.resourcemanager.recoveryservicesbackup.models.IdentityInfo;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryType;
import com.azure.resourcemanager.recoveryservicesbackup.models.RestoreRequestResource;

/** Samples for Restores Trigger. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/AzureIaasVm/TriggerRestore_ALR_IaasVMRestoreRequest_IdentityBasedRestoreDetails.json
     */
    /**
     * Sample code: Restore to New Azure IaasVm with IaasVMRestoreRequest with identityBasedRestoreDetails.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void restoreToNewAzureIaasVmWithIaasVMRestoreRequestWithIdentityBasedRestoreDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .restores()
            .trigger(
                "testVault",
                "netsdktestrg",
                "Azure",
                "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
                "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
                "348916168024334",
                new RestoreRequestResource()
                    .withProperties(
                        new IaasVMRestoreRequest()
                            .withRecoveryPointId("348916168024334")
                            .withRecoveryType(RecoveryType.ALTERNATE_LOCATION)
                            .withSourceResourceId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1")
                            .withTargetVirtualMachineId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg2/providers/Microsoft.Compute/virtualmachines/RSMDALRVM981435")
                            .withTargetResourceGroupId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg2")
                            .withVirtualNetworkId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Network/virtualNetworks/testNet")
                            .withSubnetId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Network/virtualNetworks/testNet/subnets/default")
                            .withRegion("southeastasia")
                            .withCreateNewCloudService(false)
                            .withOriginalStorageAccountOption(false)
                            .withEncryptionDetails(new EncryptionDetails().withEncryptionEnabled(false))
                            .withIdentityInfo(new IdentityInfo().withIsSystemAssignedIdentity(true))
                            .withIdentityBasedRestoreDetails(
                                new IdentityBasedRestoreDetails()
                                    .withTargetStorageAccountId(
                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Storage/storageAccounts/testingAccount"))),
                com.azure.core.util.Context.NONE);
    }
}
