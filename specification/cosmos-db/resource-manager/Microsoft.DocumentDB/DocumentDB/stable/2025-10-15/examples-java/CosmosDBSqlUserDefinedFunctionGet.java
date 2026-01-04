
/**
 * Samples for SqlResources GetSqlUserDefinedFunction.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlUserDefinedFunctionGet.json
     */
    /**
     * Sample code: CosmosDBSqlUserDefinedFunctionGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlUserDefinedFunctionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().getSqlUserDefinedFunctionWithResponse(
            "rgName", "ddb1", "databaseName", "containerName", "userDefinedFunctionName",
            com.azure.core.util.Context.NONE);
    }
}
