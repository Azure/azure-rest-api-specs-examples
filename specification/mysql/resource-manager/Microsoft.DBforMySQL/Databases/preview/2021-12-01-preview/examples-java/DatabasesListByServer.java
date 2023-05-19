/** Samples for Databases ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/Databases/preview/2021-12-01-preview/examples/DatabasesListByServer.json
     */
    /**
     * Sample code: List databases in a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void listDatabasesInAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.databases().listByServer("TestGroup", "testserver", com.azure.core.util.Context.NONE);
    }
}
