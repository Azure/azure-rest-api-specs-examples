
/**
 * Samples for GremlinResources MigrateGremlinGraphToManualThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBGremlinGraphMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBGremlinGraphMigrateToManualThroughput.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        cosmosDBGremlinGraphMigrateToManualThroughput(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getGremlinResources().migrateGremlinGraphToManualThroughput(
            "rg1", "ddb1", "databaseName", "graphName", com.azure.core.util.Context.NONE);
    }
}
