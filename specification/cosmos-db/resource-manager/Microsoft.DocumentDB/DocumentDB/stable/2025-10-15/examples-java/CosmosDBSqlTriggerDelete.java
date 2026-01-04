
/**
 * Samples for SqlResources DeleteSqlTrigger.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlTriggerDelete.json
     */
    /**
     * Sample code: CosmosDBSqlTriggerDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlTriggerDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().deleteSqlTrigger("rg1", "ddb1",
            "databaseName", "containerName", "triggerName", com.azure.core.util.Context.NONE);
    }
}
