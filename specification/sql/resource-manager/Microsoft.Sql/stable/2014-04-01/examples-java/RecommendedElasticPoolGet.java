import com.azure.core.util.Context;

/** Samples for RecommendedElasticPools Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/RecommendedElasticPoolGet.json
     */
    /**
     * Sample code: Get a recommended elastic pool.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getARecommendedElasticPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getRecommendedElasticPools()
            .getWithResponse("sqlcrudtest-6852", "sqlcrudtest-2080", "ElasticPool1", Context.NONE);
    }
}
