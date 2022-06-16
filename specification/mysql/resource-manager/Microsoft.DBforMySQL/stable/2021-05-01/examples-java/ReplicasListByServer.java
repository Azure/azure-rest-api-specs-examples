import com.azure.core.util.Context;

/** Samples for Replicas ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ReplicasListByServer.json
     */
    /**
     * Sample code: List replicas for a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void listReplicasForAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.replicas().listByServer("TestGroup", "mysqltestserver", Context.NONE);
    }
}
