
/**
 * Samples for CassandraClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-05-15/examples/
     * CosmosDBManagedCassandraClusterListByResourceGroup.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraClusterListByResourceGroup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        cosmosDBManagedCassandraClusterListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCassandraClusters()
            .listByResourceGroup("cassandra-prod-rg", com.azure.core.util.Context.NONE);
    }
}
