
import com.azure.core.util.Context;

/** Samples for ElasticPools Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/HyperscaleElasticPoolGet.json
     */
    /**
     * Sample code: Get a Hyperscale elastic pool.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAHyperscaleElasticPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getElasticPools().getWithResponse("sqlcrudtest-2369",
            "sqlcrudtest-8069", "sqlcrudtest-8102", Context.NONE);
    }
}
