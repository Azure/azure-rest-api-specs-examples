import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.fluent.models.BackupResourceConfigResourceInner;
import com.azure.resourcemanager.recoveryservicesbackup.models.BackupResourceConfig;
import com.azure.resourcemanager.recoveryservicesbackup.models.StorageType;
import com.azure.resourcemanager.recoveryservicesbackup.models.StorageTypeState;

/** Samples for BackupResourceStorageConfigsNonCrr Patch. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/Common/BackupStorageConfig_Patch.json
     */
    /**
     * Sample code: Update Vault Storage Configuration.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void updateVaultStorageConfiguration(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupResourceStorageConfigsNonCrrs()
            .patchWithResponse(
                "PySDKBackupTestRsVault",
                "PythonSDKBackupTestRg",
                new BackupResourceConfigResourceInner()
                    .withProperties(
                        new BackupResourceConfig()
                            .withStorageType(StorageType.LOCALLY_REDUNDANT)
                            .withStorageTypeState(StorageTypeState.UNLOCKED)),
                Context.NONE);
    }
}
