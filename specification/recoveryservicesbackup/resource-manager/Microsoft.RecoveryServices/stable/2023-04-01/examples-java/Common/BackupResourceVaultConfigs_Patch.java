import com.azure.resourcemanager.recoveryservicesbackup.fluent.models.BackupResourceVaultConfigResourceInner;
import com.azure.resourcemanager.recoveryservicesbackup.models.BackupResourceVaultConfig;
import com.azure.resourcemanager.recoveryservicesbackup.models.EnhancedSecurityState;

/** Samples for BackupResourceVaultConfigs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/Common/BackupResourceVaultConfigs_Patch.json
     */
    /**
     * Sample code: Update Vault Security Config.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void updateVaultSecurityConfig(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupResourceVaultConfigs()
            .updateWithResponse(
                "SwaggerTest",
                "SwaggerTestRg",
                new BackupResourceVaultConfigResourceInner()
                    .withProperties(
                        new BackupResourceVaultConfig().withEnhancedSecurityState(EnhancedSecurityState.ENABLED)),
                com.azure.core.util.Context.NONE);
    }
}
