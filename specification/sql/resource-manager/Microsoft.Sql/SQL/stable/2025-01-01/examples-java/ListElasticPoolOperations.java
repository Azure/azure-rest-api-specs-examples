
/**
 * Samples for ElasticPoolOperations ListByElasticPool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListElasticPoolOperations.json
     */
    /**
     * Sample code: List the elastic pool management operations.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listTheElasticPoolManagementOperations(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPoolOperations().listByElasticPool("sqlcrudtestgroup", "sqlcrudtestserver",
            "testpool", com.azure.core.util.Context.NONE);
    }
}
