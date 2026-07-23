
/**
 * Samples for DatabaseAccounts ListMetricDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountGetMetricDefinitions.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountGetMetricDefinitions.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBDatabaseAccountGetMetricDefinitions(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().listMetricDefinitions("rg1", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
