
/**
 * Samples for BackupEngines Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/Dpm/
     * BackupEngines_Get.json
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
