
/**
 * Samples for CassandraClusters Deallocate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBManagedCassandraClusterDeallocate.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraClusterDeallocate.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBManagedCassandraClusterDeallocate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraClusters().deallocate("cassandra-prod-rg", "cassandra-prod", null,
            com.azure.core.util.Context.NONE);
    }
}
