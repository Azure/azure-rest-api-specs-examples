
/**
 * Samples for BackupResourceEncryptionConfigs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/BackupResourceEncryptionConfig_Get.json
     */
    /**
     * Sample code: Get Vault Encryption Configuration.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getVaultEncryptionConfiguration(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupResourceEncryptionConfigs().getWithResponse("rishTestVault", "rishgrp",
            com.azure.core.util.Context.NONE);
    }
}
