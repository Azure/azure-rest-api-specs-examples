
/**
 * Samples for BackupResourceVaultConfigs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/Common/BackupResourceVaultConfigs_Get.json
     */
    /**
     * Sample code: Get Vault Security Config.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void
        getVaultSecurityConfig(com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupResourceVaultConfigs().getWithResponse("SwaggerTest", "SwaggerTestRg",
            com.azure.core.util.Context.NONE);
    }
}
