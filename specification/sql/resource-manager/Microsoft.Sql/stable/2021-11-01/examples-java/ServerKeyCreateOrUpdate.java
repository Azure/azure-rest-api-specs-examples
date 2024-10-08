
import com.azure.resourcemanager.sql.fluent.models.ServerKeyInner;
import com.azure.resourcemanager.sql.models.ServerKeyType;

/**
 * Samples for ServerKeys CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerKeyCreateOrUpdate.json
     */
    /**
     * Sample code: Creates or updates a server key.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsOrUpdatesAServerKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerKeys().createOrUpdate("sqlcrudtest-7398",
            "sqlcrudtest-4645", "someVault_someKey_01234567890123456789012345678901",
            new ServerKeyInner().withServerKeyType(ServerKeyType.AZURE_KEY_VAULT)
                .withUri("https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901"),
            com.azure.core.util.Context.NONE);
    }
}
