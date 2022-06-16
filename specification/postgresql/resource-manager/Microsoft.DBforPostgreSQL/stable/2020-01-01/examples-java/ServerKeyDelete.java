import com.azure.core.util.Context;

/** Samples for ServerKeys Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2020-01-01/examples/ServerKeyDelete.json
     */
    /**
     * Sample code: Delete the PostgreSQL Server key.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void deleteThePostgreSQLServerKey(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager
            .serverKeys()
            .delete("testserver", "someVault_someKey_01234567890123456789012345678901", "testrg", Context.NONE);
    }
}
