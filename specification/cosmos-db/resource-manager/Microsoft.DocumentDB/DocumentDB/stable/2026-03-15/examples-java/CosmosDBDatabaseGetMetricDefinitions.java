
/**
 * Samples for Database ListMetricDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseGetMetricDefinitions.json
     */
    /**
     * Sample code: CosmosDBDatabaseGetMetricDefinitions.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDatabaseGetMetricDefinitions(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabases().listMetricDefinitions("rg1", "ddb1", "databaseRid",
            com.azure.core.util.Context.NONE);
    }
}
