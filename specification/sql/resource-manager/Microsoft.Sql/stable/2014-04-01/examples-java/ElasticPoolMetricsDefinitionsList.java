import com.azure.core.util.Context;

/** Samples for ElasticPools ListMetricDefinitions. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ElasticPoolMetricsDefinitionsList.json
     */
    /**
     * Sample code: List database usage metrics.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDatabaseUsageMetrics(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getElasticPools()
            .listMetricDefinitions("sqlcrudtest-6730", "sqlcrudtest-9007", "3481", Context.NONE);
    }
}
