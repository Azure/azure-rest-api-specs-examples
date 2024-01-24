
/**
 * Samples for BackupProtectionIntent List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/
     * AzureWorkload/BackupProtectionIntent_List.json
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
