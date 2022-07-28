import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.ElasticPoolUpdate;

/** Samples for ElasticPools Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/ElasticPoolUpdateMin.json
     */
    /**
     * Sample code: Update an elastic pool with minimum parameters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAnElasticPoolWithMinimumParameters(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getElasticPools()
            .update("sqlcrudtest-2369", "sqlcrudtest-8069", "sqlcrudtest-8102", new ElasticPoolUpdate(), Context.NONE);
    }
}
