import com.azure.resourcemanager.mysql.models.ServerKeyType;

/** Samples for ServerKeys CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2020-01-01/examples/ServerKeyCreateOrUpdate.json
     */
    /**
     * Sample code: Creates or updates a MySQL Server key.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void createsOrUpdatesAMySQLServerKey(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .serverKeys()
            .define("someVault_someKey_01234567890123456789012345678901")
            .withExistingServer("testserver", "testrg")
            .withServerKeyType(ServerKeyType.AZURE_KEY_VAULT)
            .withUri("https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901")
            .create();
    }
}
