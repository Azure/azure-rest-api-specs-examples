import com.azure.core.util.Context;

/** Samples for BackupProtectionIntent List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureWorkload/BackupProtectionIntent_List.json
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
