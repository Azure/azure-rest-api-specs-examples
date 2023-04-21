/** Samples for GremlinResources MigrateGremlinGraphToAutoscale. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBGremlinGraphMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBGremlinGraphMigrateToAutoscale.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBGremlinGraphMigrateToAutoscale(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getGremlinResources()
            .migrateGremlinGraphToAutoscale(
                "rg1", "ddb1", "databaseName", "graphName", com.azure.core.util.Context.NONE);
    }
}
