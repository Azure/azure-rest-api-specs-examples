import com.azure.core.util.Context;

/** Samples for BackupResourceEncryptionConfigs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/BackupResourceEncryptionConfig_Get.json
     */
    /**
     * Sample code: Get Vault Encryption Configuration.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getVaultEncryptionConfiguration(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupResourceEncryptionConfigs().getWithResponse("rishTestVault", "rishgrp", Context.NONE);
    }
}
