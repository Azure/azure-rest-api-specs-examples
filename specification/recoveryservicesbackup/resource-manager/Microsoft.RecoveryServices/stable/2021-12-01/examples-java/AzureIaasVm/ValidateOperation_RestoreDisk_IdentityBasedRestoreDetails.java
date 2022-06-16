import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.EncryptionDetails;
import com.azure.resourcemanager.recoveryservicesbackup.models.IaasVMRestoreRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.IdentityBasedRestoreDetails;
import com.azure.resourcemanager.recoveryservicesbackup.models.IdentityInfo;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryType;
import com.azure.resourcemanager.recoveryservicesbackup.models.ValidateIaasVMRestoreOperationRequest;

/** Samples for OperationOperation Validate. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureIaasVm/ValidateOperation_RestoreDisk_IdentityBasedRestoreDetails.json
     */
    /**
     * Sample code: Validate Operation with identityBasedRestoreDetails.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void validateOperationWithIdentityBasedRestoreDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .operationOperations()
            .validateWithResponse(
                "testVault",
                "testRG",
                new ValidateIaasVMRestoreOperationRequest()
                    .withRestoreRequest(
                        new IaasVMRestoreRequest()
                            .withRecoveryPointId("348916168024334")
                            .withRecoveryType(RecoveryType.RESTORE_DISKS)
                            .withSourceResourceId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1")
                            .withRegion("southeastasia")
                            .withCreateNewCloudService(true)
                            .withOriginalStorageAccountOption(false)
                            .withEncryptionDetails(new EncryptionDetails().withEncryptionEnabled(false))
                            .withIdentityInfo(
                                new IdentityInfo()
                                    .withIsSystemAssignedIdentity(false)
                                    .withManagedIdentityResourceId(
                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/asmaskarRG1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/asmaskartestmsi"))
                            .withIdentityBasedRestoreDetails(
                                new IdentityBasedRestoreDetails()
                                    .withTargetStorageAccountId(
                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testingRg/providers/Microsoft.Storage/storageAccounts/testAccount"))),
                Context.NONE);
    }
}
