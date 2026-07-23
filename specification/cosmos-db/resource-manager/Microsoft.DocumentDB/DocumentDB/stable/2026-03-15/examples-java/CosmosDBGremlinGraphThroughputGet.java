
/**
 * Samples for GremlinResources GetGremlinGraphThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBGremlinGraphThroughputGet.json
     */
    /**
     * Sample code: CosmosDBGremlinGraphThroughputGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGremlinGraphThroughputGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().getGremlinGraphThroughputWithResponse("rg1", "ddb1",
            "databaseName", "graphName", com.azure.core.util.Context.NONE);
    }
}
