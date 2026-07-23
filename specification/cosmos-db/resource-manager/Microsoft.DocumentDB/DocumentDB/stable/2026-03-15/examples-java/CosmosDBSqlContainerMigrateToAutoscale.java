
/**
 * Samples for SqlResources MigrateSqlContainerToAutoscale.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlContainerMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBSqlContainerMigrateToAutoscale.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlContainerMigrateToAutoscale(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().migrateSqlContainerToAutoscale("rg1", "ddb1", "databaseName",
            "containerName", com.azure.core.util.Context.NONE);
    }
}
