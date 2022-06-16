import com.azure.core.util.Context;

/** Samples for BackupResourceStorageConfigsNonCrr Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/Common/BackupStorageConfig_Get.json
     */
    /**
     * Sample code: Get Vault Storage Configuration.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getVaultStorageConfiguration(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupResourceStorageConfigsNonCrrs()
            .getWithResponse("PySDKBackupTestRsVault", "PythonSDKBackupTestRg", Context.NONE);
    }
}
