
/**
 * Samples for SqlResources DeleteSqlTrigger.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlTriggerDelete.json
     */
    /**
     * Sample code: CosmosDBSqlTriggerDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlTriggerDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().deleteSqlTrigger("rg1", "ddb1", "databaseName", "containerName",
            "triggerName", com.azure.core.util.Context.NONE);
    }
}
