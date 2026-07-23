
/**
 * Samples for CassandraClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBManagedCassandraClusterGet.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraClusterGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBManagedCassandraClusterGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraClusters().getByResourceGroupWithResponse("cassandra-prod-rg",
            "cassandra-prod", com.azure.core.util.Context.NONE);
    }
}
