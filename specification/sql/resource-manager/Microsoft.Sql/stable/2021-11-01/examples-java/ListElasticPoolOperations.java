
/**
 * Samples for ElasticPoolOperations ListByElasticPool.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ListElasticPoolOperations.json
     */
    /**
     * Sample code: List the elastic pool management operations.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listTheElasticPoolManagementOperations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getElasticPoolOperations().listByElasticPool("sqlcrudtestgroup",
            "sqlcrudtestserver", "testpool", com.azure.core.util.Context.NONE);
    }
}
