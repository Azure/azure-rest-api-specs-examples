
/**
 * Samples for SqlResources DeleteSqlUserDefinedFunction.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlUserDefinedFunctionDelete.json
     */
    /**
     * Sample code: CosmosDBSqlUserDefinedFunctionDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlUserDefinedFunctionDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().deleteSqlUserDefinedFunction("rg1", "ddb1", "databaseName",
            "containerName", "userDefinedFunctionName", com.azure.core.util.Context.NONE);
    }
}
