
/**
 * Samples for BackupProtectionIntent List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/AzureWorkload/BackupProtectionIntent_List.json
     */
    /**
     * Sample code: List protection intent with backupManagementType filter.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listProtectionIntentWithBackupManagementTypeFilter(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupProtectionIntents().list("myVault", "myRG", null, null, com.azure.core.util.Context.NONE);
    }
}
