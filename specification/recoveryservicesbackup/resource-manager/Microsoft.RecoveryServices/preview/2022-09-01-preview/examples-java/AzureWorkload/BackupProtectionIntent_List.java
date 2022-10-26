import com.azure.core.util.Context;

/** Samples for BackupProtectionIntent List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/AzureWorkload/BackupProtectionIntent_List.json
     */
    /**
     * Sample code: List protection intent with backupManagementType filter.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listProtectionIntentWithBackupManagementTypeFilter(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupProtectionIntents().list("myVault", "myRG", null, null, Context.NONE);
    }
}
