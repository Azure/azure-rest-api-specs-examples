
import com.azure.resourcemanager.recoveryservicesbackup.models.BackupResourceEncryptionConfig;
import com.azure.resourcemanager.recoveryservicesbackup.models.BackupResourceEncryptionConfigResource;
import com.azure.resourcemanager.recoveryservicesbackup.models.EncryptionAtRestType;
import com.azure.resourcemanager.recoveryservicesbackup.models.InfrastructureEncryptionState;

/**
 * Samples for BackupResourceEncryptionConfigs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/
     * BackupResourceEncryptionConfig_Put.json
     */
    /**
     * Sample code: Update Vault Encryption Configuration.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void updateVaultEncryptionConfiguration(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupResourceEncryptionConfigs().updateWithResponse("source-rsv", "test-rg",
            new BackupResourceEncryptionConfigResource().withProperties(
                new BackupResourceEncryptionConfig().withEncryptionAtRestType(EncryptionAtRestType.CUSTOMER_MANAGED)
                    .withKeyUri("fakeTokenPlaceholder").withSubscriptionId("1a2311d9-66f5-47d3-a9fb-7a37da63934b")
                    .withInfrastructureEncryptionState(InfrastructureEncryptionState.fromString("true"))),
            com.azure.core.util.Context.NONE);
    }
}
