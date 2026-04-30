
/**
 * Samples for BackupResourceStorageConfigsNonCrr Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/Common/BackupStorageConfig_Get.json
     */
    /**
     * Sample code: Get Vault Storage Configuration.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getVaultStorageConfiguration(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupResourceStorageConfigsNonCrrs().getWithResponse("PySDKBackupTestRsVault", "PythonSDKBackupTestRg",
            com.azure.core.util.Context.NONE);
    }
}
