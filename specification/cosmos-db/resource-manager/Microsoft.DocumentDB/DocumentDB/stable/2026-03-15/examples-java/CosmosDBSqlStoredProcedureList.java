
/**
 * Samples for SqlResources ListSqlStoredProcedures.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlStoredProcedureList.json
     */
    /**
     * Sample code: CosmosDBSqlStoredProcedureList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlStoredProcedureList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().listSqlStoredProcedures("rgName", "ddb1", "databaseName",
            "containerName", com.azure.core.util.Context.NONE);
    }
}
