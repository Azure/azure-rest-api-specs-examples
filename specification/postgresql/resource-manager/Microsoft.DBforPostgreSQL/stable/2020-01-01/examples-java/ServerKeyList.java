import com.azure.core.util.Context;

/** Samples for ServerKeys List. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2020-01-01/examples/ServerKeyList.json
     */
    /**
     * Sample code: List the keys for a PostgreSQL Server.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listTheKeysForAPostgreSQLServer(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.serverKeys().list("testrg", "testserver", Context.NONE);
    }
}
