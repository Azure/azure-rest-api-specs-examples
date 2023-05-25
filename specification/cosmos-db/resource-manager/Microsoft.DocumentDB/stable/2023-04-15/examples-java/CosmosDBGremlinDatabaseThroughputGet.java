/** Samples for GremlinResources GetGremlinDatabaseThroughput. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-04-15/examples/CosmosDBGremlinDatabaseThroughputGet.json
     */
    /**
     * Sample code: CosmosDBGremlinDatabaseThroughputGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBGremlinDatabaseThroughputGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getGremlinResources()
            .getGremlinDatabaseThroughputWithResponse("rg1", "ddb1", "databaseName", com.azure.core.util.Context.NONE);
    }
}
