
/**
 * Samples for Replicas ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/
     * ReplicasListByServer.json
     */
    /**
     * Sample code: ReplicasListByServer.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        replicasListByServer(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.replicas().listByServer("testrg", "sourcepgservername", com.azure.core.util.Context.NONE);
    }
}
