import com.azure.core.util.Context;

/** Samples for Replicas ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ReplicasListByServer.json
     */
    /**
     * Sample code: ReplicasListByServer.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void replicasListByServer(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.replicas().listByServer("TestGroup_WestCentralUS", "testserver-master", Context.NONE);
    }
}
