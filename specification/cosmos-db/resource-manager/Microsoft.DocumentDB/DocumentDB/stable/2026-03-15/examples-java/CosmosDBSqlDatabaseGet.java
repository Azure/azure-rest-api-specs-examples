
/**
 * Samples for SqlResources GetSqlDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlDatabaseGet.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlDatabaseGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().getSqlDatabaseWithResponse("rg1", "ddb1", "databaseName",
            com.azure.core.util.Context.NONE);
    }
}
