import com.azure.core.util.Context;

/** Samples for Servers Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ServerStop.json
     */
    /**
     * Sample code: Stop a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void stopAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().stop("TestGroup", "testserver", Context.NONE);
    }
}
