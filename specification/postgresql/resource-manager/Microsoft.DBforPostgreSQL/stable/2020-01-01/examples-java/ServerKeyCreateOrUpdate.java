
import com.azure.resourcemanager.postgresql.models.ServerKeyType;

/**
 * Samples for ServerKeys CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2020-01-01/examples/
     * ServerKeyCreateOrUpdate.json
     */
    /**
     * Sample code: Creates or updates a PostgreSQL Server key.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        createsOrUpdatesAPostgreSQLServerKey(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.serverKeys().define("someVault_someKey_01234567890123456789012345678901")
            .withExistingServer("testserver", "testrg").withServerKeyType(ServerKeyType.AZURE_KEY_VAULT)
            .withUri("https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901").create();
    }
}
