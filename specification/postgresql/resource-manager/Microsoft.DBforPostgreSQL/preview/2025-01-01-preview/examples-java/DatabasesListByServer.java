
/**
 * Samples for Databases ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * DatabasesListByServer.json
     */
    /**
     * Sample code: List databases in a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        listDatabasesInAServer(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.databases().listByServer("TestGroup", "testserver", com.azure.core.util.Context.NONE);
    }
}
