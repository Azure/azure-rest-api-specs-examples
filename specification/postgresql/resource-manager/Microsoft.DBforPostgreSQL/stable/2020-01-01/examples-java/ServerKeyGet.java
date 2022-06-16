import com.azure.core.util.Context;

/** Samples for ServerKeys Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2020-01-01/examples/ServerKeyGet.json
     */
    /**
     * Sample code: Get the PostgreSQL Server key.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getThePostgreSQLServerKey(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager
            .serverKeys()
            .getWithResponse(
                "testrg", "testserver", "someVault_someKey_01234567890123456789012345678901", Context.NONE);
    }
}
