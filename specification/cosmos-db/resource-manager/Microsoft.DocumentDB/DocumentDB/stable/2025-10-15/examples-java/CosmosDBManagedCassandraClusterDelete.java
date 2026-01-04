
/**
 * Samples for CassandraClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBManagedCassandraClusterDelete.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraClusterDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBManagedCassandraClusterDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCassandraClusters().delete("cassandra-prod-rg",
            "cassandra-prod", com.azure.core.util.Context.NONE);
    }
}
