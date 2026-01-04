
/**
 * Samples for GremlinResources DeleteGremlinGraph.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBGremlinGraphDelete.json
     */
    /**
     * Sample code: CosmosDBGremlinGraphDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBGremlinGraphDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getGremlinResources().deleteGremlinGraph("rg1", "ddb1",
            "databaseName", "graphName", com.azure.core.util.Context.NONE);
    }
}
