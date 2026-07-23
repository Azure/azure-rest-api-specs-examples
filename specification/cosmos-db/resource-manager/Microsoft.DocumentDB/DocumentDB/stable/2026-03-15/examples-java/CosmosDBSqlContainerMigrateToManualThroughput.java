
/**
 * Samples for SqlResources MigrateSqlContainerToManualThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlContainerMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBSqlContainerMigrateToManualThroughput.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBSqlContainerMigrateToManualThroughput(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().migrateSqlContainerToManualThroughput("rg1", "ddb1", "databaseName",
            "containerName", com.azure.core.util.Context.NONE);
    }
}
