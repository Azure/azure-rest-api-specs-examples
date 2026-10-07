
/**
 * Samples for SqlResources ListSqlTriggers.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSqlTriggerList.json
     */
    /**
     * Sample code: CosmosDBSqlTriggerList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlTriggerList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().listSqlTriggers("rgName", "ddb1", "databaseName", "containerName",
            com.azure.core.util.Context.NONE);
    }
}
