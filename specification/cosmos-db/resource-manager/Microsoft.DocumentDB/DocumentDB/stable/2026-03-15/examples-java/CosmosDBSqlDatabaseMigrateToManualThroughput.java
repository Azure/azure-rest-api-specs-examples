
/**
 * Samples for SqlResources MigrateSqlDatabaseToManualThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlDatabaseMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseMigrateToManualThroughput.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBSqlDatabaseMigrateToManualThroughput(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().migrateSqlDatabaseToManualThroughput("rg1", "ddb1", "databaseName",
            com.azure.core.util.Context.NONE);
    }
}
