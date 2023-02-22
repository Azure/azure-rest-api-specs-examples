/** Samples for Replicas ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/preview/2021-12-01-preview/examples/ReplicasListByServer.json
     */
    /**
     * Sample code: List replicas for a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void listReplicasForAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.replicas().listByServer("TestGroup", "mysqltestserver", com.azure.core.util.Context.NONE);
    }
}
