
import com.azure.resourcemanager.recoveryservicesbackup.models.DataDiskEncryptionSettings;
import com.azure.resourcemanager.recoveryservicesbackup.models.EncryptionDetails;
import com.azure.resourcemanager.recoveryservicesbackup.models.IaasVMRestoreRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.IdentityBasedRestoreDetails;
import com.azure.resourcemanager.recoveryservicesbackup.models.IdentityInfo;
import com.azure.resourcemanager.recoveryservicesbackup.models.PerDiskEncryptionSetId;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryType;
import com.azure.resourcemanager.recoveryservicesbackup.models.RestoreRequestResource;
import com.azure.resourcemanager.recoveryservicesbackup.models.SecuredVMDetails;
import java.util.Arrays;

/**
 * Samples for Restores Trigger.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/AzureIaasVm/TriggerRestore_ALR_IaasVMRestoreRequest_DataDiskEncryption.json
     */
    /**
     * Sample code: Restore to New Azure IaasVm with IaasVMRestoreRequest with identityBasedRestoreDetails and per disk
     * encryption settings.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void
        restoreToNewAzureIaasVmWithIaasVMRestoreRequestWithIdentityBasedRestoreDetailsAndPerDiskEncryptionSettings(
            com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.restores().trigger("testVault", "netsdktestrg", "Azure",
            "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
            "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "348916168024334",
            new RestoreRequestResource().withProperties(new IaasVMRestoreRequest()
                .withRecoveryPointId("348916168024334").withRecoveryType(RecoveryType.ALTERNATE_LOCATION)
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
                .withRegion("southeastasia").withCreateNewCloudService(false).withOriginalStorageAccountOption(false)
                .withEncryptionDetails(new EncryptionDetails().withEncryptionEnabled(false))
                .withIdentityInfo(new IdentityInfo().withIsSystemAssignedIdentity(true))
                .withIdentityBasedRestoreDetails(new IdentityBasedRestoreDetails().withTargetStorageAccountId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Storage/storageAccounts/testingAccount"))
                .withSecuredVMDetails(new SecuredVMDetails().withSecuredVmosDiskEncryptionSetId(
                    "/subscriptions/5288acd1-ba79-4377-9205-9f220331a44a/resourceGroups/asmaskarrg-cvm-os-cmk-3cmk-1none-data-2606040706/providers/Microsoft.Compute/diskEncryptionSets/des-os")
                    .withDataDiskEncryptionSettings(new DataDiskEncryptionSettings().withPerDiskEncryptionSetIds(
                        Arrays.asList(new PerDiskEncryptionSetId().withLun(0).withDiskEncryptionSetId(
                            "/subscriptions/5288acd1-ba79-4377-9205-9f220331a44a/resourceGroups/asmaskarrg-cvm-os-cmk-3cmk-1none-data-2606040706/providers/Microsoft.Compute/diskEncryptionSets/des-data-lun-0"),
                            new PerDiskEncryptionSetId().withLun(1).withDiskEncryptionSetId(
                                "/subscriptions/5288acd1-ba79-4377-9205-9f220331a44a/resourceGroups/asmaskarrg-cvm-os-cmk-3cmk-1none-data-2606040706/providers/Microsoft.Compute/diskEncryptionSets/des-data-lun-1"),
                            new PerDiskEncryptionSetId().withLun(2).withDiskEncryptionSetId(
                                "/subscriptions/5288acd1-ba79-4377-9205-9f220331a44a/resourceGroups/asmaskarrg-cvm-os-cmk-3cmk-1none-data-2606040706/providers/Microsoft.Compute/diskEncryptionSets/des-data-lun-2")))
                        .withDataDiskEncryptionIdentity(
                            "/subscriptions/5288acd1-ba79-4377-9205-9f220331a44a/resourcegroups/asmaskarrg-cvm-os-cmk-3cmk-1none-data-2606040706/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami-cdde")))),
            com.azure.core.util.Context.NONE);
    }
}
