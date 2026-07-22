
/**
 * Samples for GremlinResources GetGremlinGraph.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBGremlinGraphGet.json
     */
    /**
     * Sample code: CosmosDBGremlinGraphGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGremlinGraphGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().getGremlinGraphWithResponse("rgName", "ddb1", "databaseName",
            "graphName", com.azure.core.util.Context.NONE);
    }
}
