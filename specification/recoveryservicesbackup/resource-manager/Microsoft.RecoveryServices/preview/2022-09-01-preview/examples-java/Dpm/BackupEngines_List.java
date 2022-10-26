import com.azure.core.util.Context;

/** Samples for BackupEngines List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/Dpm/BackupEngines_List.json
     */
    /**
     * Sample code: List Dpm/AzureBackupServer/Lajolla Backup Engines.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listDpmAzureBackupServerLajollaBackupEngines(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupEngines().list("testVault", "testRG", null, null, Context.NONE);
    }
}
