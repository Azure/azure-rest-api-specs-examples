
/**
 * Samples for CassandraClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-08-15/examples/
     * CosmosDBManagedCassandraClusterListBySubscription.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraClusterListBySubscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        cosmosDBManagedCassandraClusterListBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCassandraClusters()
            .list(com.azure.core.util.Context.NONE);
    }
}
