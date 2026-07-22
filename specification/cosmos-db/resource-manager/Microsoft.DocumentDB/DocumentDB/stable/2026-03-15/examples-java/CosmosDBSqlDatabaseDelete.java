
/**
 * Samples for SqlResources DeleteSqlDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlDatabaseDelete.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlDatabaseDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().deleteSqlDatabase("rg1", "ddb1", "databaseName",
            com.azure.core.util.Context.NONE);
    }
}
