
/**
 * Samples for GremlinResources DeleteGremlinGraph.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBGremlinGraphDelete.json
     */
    /**
     * Sample code: CosmosDBGremlinGraphDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGremlinGraphDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().deleteGremlinGraph("rg1", "ddb1", "databaseName", "graphName",
            com.azure.core.util.Context.NONE);
    }
}
