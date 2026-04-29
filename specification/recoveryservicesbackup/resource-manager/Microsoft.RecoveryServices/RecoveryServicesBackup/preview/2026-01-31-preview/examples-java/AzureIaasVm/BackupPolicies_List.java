
/**
 * Samples for BackupPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/AzureIaasVm/BackupPolicies_List.json
     */
    /**
     * Sample code: List protection policies with backupManagementType filter as AzureIaasVm.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listProtectionPoliciesWithBackupManagementTypeFilterAsAzureIaasVm(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupPolicies().list("NetSDKTestRsVault", "SwaggerTestRg", "backupManagementType eq 'AzureIaasVM'",
            com.azure.core.util.Context.NONE);
    }
}
