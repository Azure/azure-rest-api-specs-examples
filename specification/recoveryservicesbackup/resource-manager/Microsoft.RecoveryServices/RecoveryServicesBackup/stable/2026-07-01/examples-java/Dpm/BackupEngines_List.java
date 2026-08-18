
/**
 * Samples for BackupEngines List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Dpm/BackupEngines_List.json
     */
    /**
     * Sample code: List Dpm/AzureBackupServer/Lajolla Backup Engines.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listDpmAzureBackupServerLajollaBackupEngines(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupEngines().list("testVault", "testRG", null, null, com.azure.core.util.Context.NONE);
    }
}
