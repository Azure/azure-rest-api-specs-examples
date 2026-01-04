
/**
 * Samples for SqlResources DeleteSqlUserDefinedFunction.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlUserDefinedFunctionDelete.json
     */
    /**
     * Sample code: CosmosDBSqlUserDefinedFunctionDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlUserDefinedFunctionDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().deleteSqlUserDefinedFunction("rg1", "ddb1",
            "databaseName", "containerName", "userDefinedFunctionName", com.azure.core.util.Context.NONE);
    }
}
