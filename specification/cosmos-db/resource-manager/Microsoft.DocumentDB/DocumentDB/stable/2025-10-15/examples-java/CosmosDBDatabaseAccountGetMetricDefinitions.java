
/**
 * Samples for DatabaseAccounts ListMetricDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBDatabaseAccountGetMetricDefinitions.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountGetMetricDefinitions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        cosmosDBDatabaseAccountGetMetricDefinitions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getDatabaseAccounts().listMetricDefinitions("rg1", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
