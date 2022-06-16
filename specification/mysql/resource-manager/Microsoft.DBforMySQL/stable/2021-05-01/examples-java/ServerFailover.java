import com.azure.core.util.Context;

/** Samples for Servers Failover. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ServerFailover.json
     */
    /**
     * Sample code: Restart a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void restartAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().failover("TestGroup", "testserver", Context.NONE);
    }
}
