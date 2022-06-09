```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.fluent.models.BackupResourceVaultConfigResourceInner;
import com.azure.resourcemanager.recoveryservicesbackup.models.BackupResourceVaultConfig;
import com.azure.resourcemanager.recoveryservicesbackup.models.EnhancedSecurityState;
import com.azure.resourcemanager.recoveryservicesbackup.models.SoftDeleteFeatureState;

/** Samples for BackupResourceVaultConfigs Put. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/Common/BackupResourceVaultConfigs_Put.json
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
            .putWithResponse(
                "SwaggerTest",
                "SwaggerTestRg",
                new BackupResourceVaultConfigResourceInner()
                    .withProperties(
                        new BackupResourceVaultConfig()
                            .withEnhancedSecurityState(EnhancedSecurityState.ENABLED)
                            .withSoftDeleteFeatureState(SoftDeleteFeatureState.ENABLED)),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.2/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.
