/** Samples for GremlinResources MigrateGremlinDatabaseToAutoscale. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBGremlinDatabaseMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBGremlinDatabaseMigrateToAutoscale.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBGremlinDatabaseMigrateToAutoscale(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getGremlinResources()
            .migrateGremlinDatabaseToAutoscale("rg1", "ddb1", "databaseName", com.azure.core.util.Context.NONE);
    }
}
