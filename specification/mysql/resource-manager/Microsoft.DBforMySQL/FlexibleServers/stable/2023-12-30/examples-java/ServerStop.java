
/**
 * Samples for Servers Stop.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/ServerStop.
     * json
     */
    /**
     * Sample code: Stop a server.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void stopAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().stop("TestGroup", "testserver", com.azure.core.util.Context.NONE);
    }
}
