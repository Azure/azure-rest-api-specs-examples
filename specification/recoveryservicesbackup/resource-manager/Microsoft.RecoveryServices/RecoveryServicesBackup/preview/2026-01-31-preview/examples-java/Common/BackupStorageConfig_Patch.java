
import com.azure.resourcemanager.recoveryservicesbackup.fluent.models.BackupResourceConfigResourceInner;
import com.azure.resourcemanager.recoveryservicesbackup.models.BackupResourceConfig;
import com.azure.resourcemanager.recoveryservicesbackup.models.StorageType;
import com.azure.resourcemanager.recoveryservicesbackup.models.StorageTypeState;

/**
 * Samples for BackupResourceStorageConfigsNonCrr Patch.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/Common/BackupStorageConfig_Patch.json
     */
    /**
     * Sample code: Update Vault Storage Configuration.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void updateVaultStorageConfiguration(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupResourceStorageConfigsNonCrrs().patchWithResponse("PySDKBackupTestRsVault",
            "PythonSDKBackupTestRg",
            new BackupResourceConfigResourceInner().withProperties(new BackupResourceConfig()
                .withStorageType(StorageType.LOCALLY_REDUNDANT).withStorageTypeState(StorageTypeState.UNLOCKED)),
            com.azure.core.util.Context.NONE);
    }
}
