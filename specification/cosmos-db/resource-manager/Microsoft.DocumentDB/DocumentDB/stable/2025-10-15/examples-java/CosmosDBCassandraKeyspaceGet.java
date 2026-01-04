
/**
 * Samples for CassandraResources GetCassandraKeyspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBCassandraKeyspaceGet.json
     */
    /**
     * Sample code: CosmosDBCassandraKeyspaceGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBCassandraKeyspaceGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCassandraResources()
            .getCassandraKeyspaceWithResponse("rg1", "ddb1", "keyspaceName", com.azure.core.util.Context.NONE);
    }
}
