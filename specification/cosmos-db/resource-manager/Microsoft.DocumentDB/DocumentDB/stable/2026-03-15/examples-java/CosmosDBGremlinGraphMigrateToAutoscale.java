
/**
 * Samples for GremlinResources MigrateGremlinGraphToAutoscale.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBGremlinGraphMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBGremlinGraphMigrateToAutoscale.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGremlinGraphMigrateToAutoscale(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().migrateGremlinGraphToAutoscale("rg1", "ddb1", "databaseName",
            "graphName", com.azure.core.util.Context.NONE);
    }
}
