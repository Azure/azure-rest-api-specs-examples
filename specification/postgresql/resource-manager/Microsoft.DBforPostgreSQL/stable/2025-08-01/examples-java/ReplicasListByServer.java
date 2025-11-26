
/**
 * Samples for Replicas ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * ReplicasListByServer.json
     */
    /**
     * Sample code: List all read replicas of a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        listAllReadReplicasOfAServer(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.replicas().listByServer("exampleresourcegroup", "exampleserver", com.azure.core.util.Context.NONE);
    }
}
