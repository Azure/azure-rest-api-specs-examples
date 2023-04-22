/** Samples for GremlinResources GetGremlinGraph. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBGremlinGraphGet.json
     */
    /**
     * Sample code: CosmosDBGremlinGraphGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBGremlinGraphGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getGremlinResources()
            .getGremlinGraphWithResponse(
                "rgName", "ddb1", "databaseName", "graphName", com.azure.core.util.Context.NONE);
    }
}
