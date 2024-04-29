
/**
 * Samples for Servers Failover.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/
     * ServerFailover.json
     */
    /**
     * Sample code: Restart a server.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void restartAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().failover("TestGroup", "testserver", com.azure.core.util.Context.NONE);
    }
}
