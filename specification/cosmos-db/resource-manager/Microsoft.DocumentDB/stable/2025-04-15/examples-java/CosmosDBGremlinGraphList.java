
/**
 * Samples for GremlinResources ListGremlinGraphs.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBGremlinGraphList
     * .json
     */
    /**
     * Sample code: CosmosDBGremlinGraphList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBGremlinGraphList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getGremlinResources().listGremlinGraphs("rgName", "ddb1",
            "databaseName", com.azure.core.util.Context.NONE);
    }
}
