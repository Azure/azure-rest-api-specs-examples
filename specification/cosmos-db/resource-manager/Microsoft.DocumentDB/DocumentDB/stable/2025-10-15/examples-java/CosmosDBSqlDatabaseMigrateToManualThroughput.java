
/**
 * Samples for SqlResources MigrateSqlDatabaseToManualThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlDatabaseMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseMigrateToManualThroughput.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        cosmosDBSqlDatabaseMigrateToManualThroughput(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().migrateSqlDatabaseToManualThroughput("rg1",
            "ddb1", "databaseName", com.azure.core.util.Context.NONE);
    }
}
