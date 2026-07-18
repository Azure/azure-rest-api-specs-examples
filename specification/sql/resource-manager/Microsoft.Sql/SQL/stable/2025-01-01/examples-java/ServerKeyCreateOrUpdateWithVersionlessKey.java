
import com.azure.resourcemanager.sql.fluent.models.ServerKeyInner;
import com.azure.resourcemanager.sql.models.ServerKeyType;

/**
 * Samples for ServerKeys CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerKeyCreateOrUpdateWithVersionlessKey.json
     */
    /**
     * Sample code: Creates or updates a server key with versionless key.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createsOrUpdatesAServerKeyWithVersionlessKey(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerKeys().createOrUpdate("sqlcrudtest-7398", "sqlcrudtest-4645",
            "someVault_someKey", new ServerKeyInner().withServerKeyType(ServerKeyType.AZURE_KEY_VAULT)
                .withUri("https://someVault.vault.azure.net/keys/someKey"),
            com.azure.core.util.Context.NONE);
    }
}
