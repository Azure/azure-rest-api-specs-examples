
/**
 * Samples for GraphResources GetGraph.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBGraphResourceGet.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlDatabaseGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGraphResources().getGraphWithResponse("rg1", "ddb1", "graphName",
            com.azure.core.util.Context.NONE);
    }
}
