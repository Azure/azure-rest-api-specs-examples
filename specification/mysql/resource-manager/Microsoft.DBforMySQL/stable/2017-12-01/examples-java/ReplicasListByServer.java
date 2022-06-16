import com.azure.core.util.Context;

/** Samples for Replicas ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ReplicasListByServer.json
     */
    /**
     * Sample code: ReplicasListByServer.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void replicasListByServer(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.replicas().listByServer("TestGroup", "testmaster", Context.NONE);
    }
}
