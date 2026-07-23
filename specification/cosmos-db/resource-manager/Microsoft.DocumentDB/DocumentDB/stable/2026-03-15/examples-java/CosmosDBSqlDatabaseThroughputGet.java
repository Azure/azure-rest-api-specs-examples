
/**
 * Samples for SqlResources GetSqlDatabaseThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlDatabaseThroughputGet.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseThroughputGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlDatabaseThroughputGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().getSqlDatabaseThroughputWithResponse("rg1", "ddb1", "databaseName",
            com.azure.core.util.Context.NONE);
    }
}
