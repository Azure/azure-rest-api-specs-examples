
/**
 * Samples for GremlinResources GetGremlinDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBGremlinDatabaseGet.json
     */
    /**
     * Sample code: CosmosDBGremlinDatabaseGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGremlinDatabaseGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().getGremlinDatabaseWithResponse("rg1", "ddb1", "databaseName",
            com.azure.core.util.Context.NONE);
    }
}
