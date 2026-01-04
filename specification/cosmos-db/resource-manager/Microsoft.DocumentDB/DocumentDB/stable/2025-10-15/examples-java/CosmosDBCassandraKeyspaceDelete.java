
/**
 * Samples for CassandraResources DeleteCassandraKeyspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBCassandraKeyspaceDelete.json
     */
    /**
     * Sample code: CosmosDBCassandraKeyspaceDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBCassandraKeyspaceDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCassandraResources().deleteCassandraKeyspace("rg1",
            "ddb1", "keyspaceName", com.azure.core.util.Context.NONE);
    }
}
