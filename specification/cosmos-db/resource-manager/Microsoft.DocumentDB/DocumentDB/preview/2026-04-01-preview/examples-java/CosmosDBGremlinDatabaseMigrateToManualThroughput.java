
/**
 * Samples for GremlinResources MigrateGremlinDatabaseToManualThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBGremlinDatabaseMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBGremlinDatabaseMigrateToManualThroughput.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBGremlinDatabaseMigrateToManualThroughput(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().migrateGremlinDatabaseToManualThroughput("rg1", "ddb1",
            "databaseName", com.azure.core.util.Context.NONE);
    }
}
