
/**
 * Samples for GremlinResources ListGremlinGraphs.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBGremlinGraphList.json
     */
    /**
     * Sample code: CosmosDBGremlinGraphList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGremlinGraphList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().listGremlinGraphs("rgName", "ddb1", "databaseName",
            com.azure.core.util.Context.NONE);
    }
}
