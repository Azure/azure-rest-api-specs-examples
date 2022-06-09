```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.BackupResourceEncryptionConfig;
import com.azure.resourcemanager.recoveryservicesbackup.models.BackupResourceEncryptionConfigResource;
import com.azure.resourcemanager.recoveryservicesbackup.models.EncryptionAtRestType;
import com.azure.resourcemanager.recoveryservicesbackup.models.InfrastructureEncryptionState;

/** Samples for BackupResourceEncryptionConfigs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/BackupResourceEncryptionConfig_Put.json
     */
    /**
     * Sample code: Update Vault Encryption Configuration.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void updateVaultEncryptionConfiguration(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupResourceEncryptionConfigs()
            .updateWithResponse(
                "source-rsv",
                "test-rg",
                new BackupResourceEncryptionConfigResource()
                    .withProperties(
                        new BackupResourceEncryptionConfig()
                            .withEncryptionAtRestType(EncryptionAtRestType.CUSTOMER_MANAGED)
                            .withKeyUri("https://gktestkv1.vault.azure.net/keys/Test1/ed2e8cdc7f86477ebf0c6462b504a9ed")
                            .withSubscriptionId("1a2311d9-66f5-47d3-a9fb-7a37da63934b")
                            .withInfrastructureEncryptionState(InfrastructureEncryptionState.fromString("true"))),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.4/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.
