
/**
 * Samples for Replicas ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/
     * ReplicasListByServer.json
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
