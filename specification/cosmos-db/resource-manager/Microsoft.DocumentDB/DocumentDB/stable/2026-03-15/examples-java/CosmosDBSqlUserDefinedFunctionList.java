
/**
 * Samples for SqlResources ListSqlUserDefinedFunctions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlUserDefinedFunctionList.json
     */
    /**
     * Sample code: CosmosDBSqlUserDefinedFunctionList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlUserDefinedFunctionList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().listSqlUserDefinedFunctions("rgName", "ddb1", "databaseName",
            "containerName", com.azure.core.util.Context.NONE);
    }
}
