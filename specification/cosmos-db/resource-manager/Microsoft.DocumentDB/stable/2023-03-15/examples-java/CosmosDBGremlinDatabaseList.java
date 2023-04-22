/** Samples for GremlinResources ListGremlinDatabases. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBGremlinDatabaseList.json
     */
    /**
     * Sample code: CosmosDBGremlinDatabaseList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBGremlinDatabaseList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getGremlinResources()
            .listGremlinDatabases("rgName", "ddb1", com.azure.core.util.Context.NONE);
    }
}
