
/**
 * Samples for CassandraClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBManagedCassandraClusterListBySubscription.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraClusterListBySubscription.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBManagedCassandraClusterListBySubscription(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraClusters().list(com.azure.core.util.Context.NONE);
    }
}
