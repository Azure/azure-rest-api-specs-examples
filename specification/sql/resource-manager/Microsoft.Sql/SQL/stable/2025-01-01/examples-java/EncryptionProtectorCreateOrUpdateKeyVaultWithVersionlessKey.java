
import com.azure.resourcemanager.sql.fluent.models.EncryptionProtectorInner;
import com.azure.resourcemanager.sql.models.EncryptionProtectorName;
import com.azure.resourcemanager.sql.models.ServerKeyType;

/**
 * Samples for EncryptionProtectors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/EncryptionProtectorCreateOrUpdateKeyVaultWithVersionlessKey.json
     */
    /**
     * Sample code: Update the encryption protector to key vault with versionless key.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateTheEncryptionProtectorToKeyVaultWithVersionlessKey(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getEncryptionProtectors().createOrUpdate("sqlcrudtest-7398", "sqlcrudtest-4645",
            EncryptionProtectorName.CURRENT,
            new EncryptionProtectorInner().withServerKeyName("fakeTokenPlaceholder")
                .withServerKeyType(ServerKeyType.AZURE_KEY_VAULT).withAutoRotationEnabled(false),
            com.azure.core.util.Context.NONE);
    }
}
