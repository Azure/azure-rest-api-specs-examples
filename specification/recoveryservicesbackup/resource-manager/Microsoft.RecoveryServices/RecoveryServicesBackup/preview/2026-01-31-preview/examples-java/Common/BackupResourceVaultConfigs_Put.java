
import com.azure.resourcemanager.recoveryservicesbackup.fluent.models.BackupResourceVaultConfigResourceInner;
import com.azure.resourcemanager.recoveryservicesbackup.models.BackupResourceVaultConfig;
import com.azure.resourcemanager.recoveryservicesbackup.models.EnhancedSecurityState;
import com.azure.resourcemanager.recoveryservicesbackup.models.SoftDeleteFeatureState;

/**
 * Samples for BackupResourceVaultConfigs Put.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/Common/BackupResourceVaultConfigs_Put.json
     */
    /**
     * Sample code: Update Vault Security Config.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void updateVaultSecurityConfig(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupResourceVaultConfigs().putWithResponse("SwaggerTest", "SwaggerTestRg",
            new BackupResourceVaultConfigResourceInner()
                .withProperties(new BackupResourceVaultConfig().withEnhancedSecurityState(EnhancedSecurityState.ENABLED)
                    .withSoftDeleteFeatureState(SoftDeleteFeatureState.ENABLED)),
            com.azure.core.util.Context.NONE);
    }
}
