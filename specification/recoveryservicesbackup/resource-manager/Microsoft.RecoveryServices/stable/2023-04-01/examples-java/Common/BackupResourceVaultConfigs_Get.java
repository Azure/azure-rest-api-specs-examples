/** Samples for BackupResourceVaultConfigs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/Common/BackupResourceVaultConfigs_Get.json
     */
    /**
     * Sample code: Get Vault Security Config.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getVaultSecurityConfig(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupResourceVaultConfigs()
            .getWithResponse("SwaggerTest", "SwaggerTestRg", com.azure.core.util.Context.NONE);
    }
}
