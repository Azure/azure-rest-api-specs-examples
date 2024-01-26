
import com.azure.core.util.Context;

/** Samples for ElasticPools Failover. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/FailoverElasticPool.json
     */
    /**
     * Sample code: Failover an elastic pool.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void failoverAnElasticPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getElasticPools().failover("group1", "testServer",
            "testElasticPool", Context.NONE);
    }
}
