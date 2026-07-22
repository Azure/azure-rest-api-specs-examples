
/**
 * Samples for SqlResources GetSqlTrigger.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlTriggerGet.json
     */
    /**
     * Sample code: CosmosDBSqlTriggerGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlTriggerGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().getSqlTriggerWithResponse("rgName", "ddb1", "databaseName",
            "containerName", "triggerName", com.azure.core.util.Context.NONE);
    }
}
