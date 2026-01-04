
/**
 * Samples for CassandraClusters Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBManagedCassandraClusterStart.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraClusterStart.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBManagedCassandraClusterStart(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCassandraClusters().start("cassandra-prod-rg",
            "cassandra-prod", com.azure.core.util.Context.NONE);
    }
}
