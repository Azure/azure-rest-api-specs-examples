
import com.azure.resourcemanager.sql.fluent.models.ManagedInstanceEncryptionProtectorInner;
import com.azure.resourcemanager.sql.models.EncryptionProtectorName;
import com.azure.resourcemanager.sql.models.ServerKeyType;

/**
 * Samples for ManagedInstanceEncryptionProtectors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceEncryptionProtectorCreateOrUpdateKeyVault.json
     */
    /**
     * Sample code: Update the encryption protector to key vault.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateTheEncryptionProtectorToKeyVault(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceEncryptionProtectors().createOrUpdate("sqlcrudtest-7398",
            "sqlcrudtest-4645", EncryptionProtectorName.CURRENT,
            new ManagedInstanceEncryptionProtectorInner().withServerKeyName("fakeTokenPlaceholder")
                .withServerKeyType(ServerKeyType.AZURE_KEY_VAULT).withAutoRotationEnabled(false),
            com.azure.core.util.Context.NONE);
    }
}
