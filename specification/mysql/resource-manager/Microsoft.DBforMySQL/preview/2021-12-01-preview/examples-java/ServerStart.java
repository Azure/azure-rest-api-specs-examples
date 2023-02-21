/** Samples for Servers Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/preview/2021-12-01-preview/examples/ServerStart.json
     */
    /**
     * Sample code: Start a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void startAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().start("TestGroup", "testserver", com.azure.core.util.Context.NONE);
    }
}
