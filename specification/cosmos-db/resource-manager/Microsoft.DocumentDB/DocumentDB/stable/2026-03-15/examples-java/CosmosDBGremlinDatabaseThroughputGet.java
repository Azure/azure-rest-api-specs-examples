
/**
 * Samples for GremlinResources GetGremlinDatabaseThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBGremlinDatabaseThroughputGet.json
     */
    /**
     * Sample code: CosmosDBGremlinDatabaseThroughputGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGremlinDatabaseThroughputGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().getGremlinDatabaseThroughputWithResponse("rg1", "ddb1",
            "databaseName", com.azure.core.util.Context.NONE);
    }
}
