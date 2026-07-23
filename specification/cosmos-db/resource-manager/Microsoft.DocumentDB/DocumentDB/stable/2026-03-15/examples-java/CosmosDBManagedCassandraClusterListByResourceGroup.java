
/**
 * Samples for CassandraClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBManagedCassandraClusterListByResourceGroup.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraClusterListByResourceGroup.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBManagedCassandraClusterListByResourceGroup(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraClusters().listByResourceGroup("cassandra-prod-rg",
            com.azure.core.util.Context.NONE);
    }
}
