
/**
 * Samples for GremlinResources DeleteGremlinDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBGremlinDatabaseDelete.json
     */
    /**
     * Sample code: CosmosDBGremlinDatabaseDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGremlinDatabaseDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().deleteGremlinDatabase("rg1", "ddb1", "databaseName",
            com.azure.core.util.Context.NONE);
    }
}
