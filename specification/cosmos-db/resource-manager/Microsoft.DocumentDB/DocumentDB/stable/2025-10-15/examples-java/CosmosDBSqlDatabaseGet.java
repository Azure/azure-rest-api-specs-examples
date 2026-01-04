
/**
 * Samples for SqlResources GetSqlDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlDatabaseGet.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlDatabaseGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().getSqlDatabaseWithResponse("rg1", "ddb1",
            "databaseName", com.azure.core.util.Context.NONE);
    }
}
