
/**
 * Samples for BackupEngines Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/Dpm/BackupEngines_Get.json
     */
    /**
     * Sample code: Get Dpm/AzureBackupServer/Lajolla Backup Engine Details.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getDpmAzureBackupServerLajollaBackupEngineDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupEngines().getWithResponse("testVault", "testRG", "testServer", null, null,
            com.azure.core.util.Context.NONE);
    }
}
