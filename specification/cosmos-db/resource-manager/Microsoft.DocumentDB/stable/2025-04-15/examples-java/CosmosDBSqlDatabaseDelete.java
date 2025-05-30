
/**
 * Samples for SqlResources DeleteSqlDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/
     * CosmosDBSqlDatabaseDelete.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlDatabaseDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().deleteSqlDatabase("rg1", "ddb1",
            "databaseName", com.azure.core.util.Context.NONE);
    }
}
