import com.azure.core.util.Context;

/** Samples for GremlinResources GetGremlinDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBGremlinDatabaseGet.json
     */
    /**
     * Sample code: CosmosDBGremlinDatabaseGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBGremlinDatabaseGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getGremlinResources()
            .getGremlinDatabaseWithResponse("rg1", "ddb1", "databaseName", Context.NONE);
    }
}
