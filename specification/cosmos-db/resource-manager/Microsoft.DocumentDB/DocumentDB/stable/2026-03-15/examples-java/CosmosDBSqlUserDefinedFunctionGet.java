
/**
 * Samples for SqlResources GetSqlUserDefinedFunction.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlUserDefinedFunctionGet.json
     */
    /**
     * Sample code: CosmosDBSqlUserDefinedFunctionGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlUserDefinedFunctionGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().getSqlUserDefinedFunctionWithResponse("rgName", "ddb1",
            "databaseName", "containerName", "userDefinedFunctionName", com.azure.core.util.Context.NONE);
    }
}
