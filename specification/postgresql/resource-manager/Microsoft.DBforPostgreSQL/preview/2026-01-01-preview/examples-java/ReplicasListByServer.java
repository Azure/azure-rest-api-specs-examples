
/**
 * Samples for Replicas ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ReplicasListByServer.json
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
