import com.azure.core.util.Context;

/** Samples for ServerKeys Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2020-01-01/examples/ServerKeyDelete.json
     */
    /**
     * Sample code: Delete the MySQL Server key.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void deleteTheMySQLServerKey(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .serverKeys()
            .delete("testserver", "someVault_someKey_01234567890123456789012345678901", "testrg", Context.NONE);
    }
}
