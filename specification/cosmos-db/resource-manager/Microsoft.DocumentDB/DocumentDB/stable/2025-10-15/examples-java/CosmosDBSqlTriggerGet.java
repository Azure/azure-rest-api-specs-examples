
/**
 * Samples for SqlResources GetSqlTrigger.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlTriggerGet.json
     */
    /**
     * Sample code: CosmosDBSqlTriggerGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlTriggerGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().getSqlTriggerWithResponse("rgName", "ddb1",
            "databaseName", "containerName", "triggerName", com.azure.core.util.Context.NONE);
    }
}
