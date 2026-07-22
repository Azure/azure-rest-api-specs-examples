
/**
 * Samples for GremlinResources ListGremlinDatabases.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBGremlinDatabaseList.json
     */
    /**
     * Sample code: CosmosDBGremlinDatabaseList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGremlinDatabaseList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().listGremlinDatabases("rgName", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
