import com.azure.core.util.Context;

/** Samples for BackupPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/AzureWorkload/BackupPolicies_List.json
     */
    /**
     * Sample code: List protection policies with backupManagementType filter as AzureWorkload.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listProtectionPoliciesWithBackupManagementTypeFilterAsAzureWorkload(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupPolicies()
            .list("NetSDKTestRsVault", "SwaggerTestRg", "backupManagementType eq 'AzureWorkload'", Context.NONE);
    }
}
