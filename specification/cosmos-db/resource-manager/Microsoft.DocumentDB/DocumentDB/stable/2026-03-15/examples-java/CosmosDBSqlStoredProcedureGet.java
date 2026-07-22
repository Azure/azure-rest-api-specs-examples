
/**
 * Samples for SqlResources GetSqlStoredProcedure.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlStoredProcedureGet.json
     */
    /**
     * Sample code: CosmosDBSqlStoredProcedureGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlStoredProcedureGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().getSqlStoredProcedureWithResponse("rgName", "ddb1", "databaseName",
            "containerName", "storedProcedureName", com.azure.core.util.Context.NONE);
    }
}
