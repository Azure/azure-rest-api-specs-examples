
/**
 * Samples for CassandraResources ListCassandraTables.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBCassandraTableList.json
     */
    /**
     * Sample code: CosmosDBCassandraTableList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBCassandraTableList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCassandraResources().listCassandraTables("rgName", "ddb1",
            "keyspaceName", com.azure.core.util.Context.NONE);
    }
}
