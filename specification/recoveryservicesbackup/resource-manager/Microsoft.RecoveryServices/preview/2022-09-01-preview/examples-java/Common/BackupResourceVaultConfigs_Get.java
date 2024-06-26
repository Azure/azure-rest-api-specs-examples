import com.azure.core.util.Context;

/** Samples for BackupResourceVaultConfigs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/Common/BackupResourceVaultConfigs_Get.json
     */
    /**
     * Sample code: Get Vault Security Config.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getVaultSecurityConfig(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupResourceVaultConfigs().getWithResponse("SwaggerTest", "SwaggerTestRg", Context.NONE);
    }
}
