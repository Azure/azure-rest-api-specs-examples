
/**
 * Samples for GremlinResources MigrateGremlinDatabaseToManualThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBGremlinDatabaseMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBGremlinDatabaseMigrateToManualThroughput.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        cosmosDBGremlinDatabaseMigrateToManualThroughput(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getGremlinResources()
            .migrateGremlinDatabaseToManualThroughput("rg1", "ddb1", "databaseName", com.azure.core.util.Context.NONE);
    }
}
