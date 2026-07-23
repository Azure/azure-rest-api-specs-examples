
/**
 * Samples for GremlinResources MigrateGremlinDatabaseToAutoscale.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBGremlinDatabaseMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBGremlinDatabaseMigrateToAutoscale.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBGremlinDatabaseMigrateToAutoscale(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().migrateGremlinDatabaseToAutoscale("rg1", "ddb1", "databaseName",
            com.azure.core.util.Context.NONE);
    }
}
