
/**
 * Samples for CassandraClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBManagedCassandraClusterGet.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraClusterGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBManagedCassandraClusterGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCassandraClusters()
            .getByResourceGroupWithResponse("cassandra-prod-rg", "cassandra-prod", com.azure.core.util.Context.NONE);
    }
}
