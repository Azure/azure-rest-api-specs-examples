
/**
 * Samples for CassandraClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBManagedCassandraClusterDelete.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraClusterDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBManagedCassandraClusterDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraClusters().delete("cassandra-prod-rg", "cassandra-prod",
            com.azure.core.util.Context.NONE);
    }
}
