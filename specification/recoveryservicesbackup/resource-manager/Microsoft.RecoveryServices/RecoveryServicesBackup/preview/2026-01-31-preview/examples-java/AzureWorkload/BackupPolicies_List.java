
/**
 * Samples for BackupPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/AzureWorkload/BackupPolicies_List.json
     */
    /**
     * Sample code: List protection policies with backupManagementType filter as AzureWorkload.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listProtectionPoliciesWithBackupManagementTypeFilterAsAzureWorkload(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupPolicies().list("NetSDKTestRsVault", "SwaggerTestRg", "backupManagementType eq 'AzureWorkload'",
            com.azure.core.util.Context.NONE);
    }
}
