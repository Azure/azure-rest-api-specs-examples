
/**
 * Samples for SqlResources MigrateSqlDatabaseToAutoscale.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlDatabaseMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseMigrateToAutoscale.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlDatabaseMigrateToAutoscale(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().migrateSqlDatabaseToAutoscale("rg1", "ddb1", "databaseName",
            com.azure.core.util.Context.NONE);
    }
}
