
/**
 * Samples for GremlinResources ListGremlinDatabases.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBGremlinDatabaseList.json
     */
    /**
     * Sample code: CosmosDBGremlinDatabaseList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBGremlinDatabaseList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getGremlinResources().listGremlinDatabases("rgName", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
