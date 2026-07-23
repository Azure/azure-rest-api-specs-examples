
/**
 * Samples for GremlinResources MigrateGremlinGraphToManualThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBGremlinGraphMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBGremlinGraphMigrateToManualThroughput.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBGremlinGraphMigrateToManualThroughput(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().migrateGremlinGraphToManualThroughput("rg1", "ddb1",
            "databaseName", "graphName", com.azure.core.util.Context.NONE);
    }
}
