
/**
 * Samples for CassandraClusters Deallocate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-08-15/examples/
     * CosmosDBManagedCassandraClusterDeallocate.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraClusterDeallocate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBManagedCassandraClusterDeallocate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCassandraClusters().deallocate("cassandra-prod-rg",
            "cassandra-prod", com.azure.core.util.Context.NONE);
    }
}
