
/**
 * Samples for ElasticPools Failover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/FailoverElasticPool.json
     */
    /**
     * Sample code: Failover an elastic pool.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void failoverAnElasticPool(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPools().failover("group1", "testServer", "testElasticPool",
            com.azure.core.util.Context.NONE);
    }
}
