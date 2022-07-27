import com.azure.core.util.Context;

/** Samples for RecommendedElasticPools ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/RecommendedElasticPoolListDecoupled.json
     */
    /**
     * Sample code: List recommended elastic pools.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRecommendedElasticPools(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getRecommendedElasticPools()
            .listByServer("sqlcrudtest-6852", "sqlcrudtest-2080", Context.NONE);
    }
}
