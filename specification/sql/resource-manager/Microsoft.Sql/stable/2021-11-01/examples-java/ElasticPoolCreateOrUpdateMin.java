
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ElasticPoolInner;

/** Samples for ElasticPools CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ElasticPoolCreateOrUpdateMin.json
     */
    /**
     * Sample code: Create or update elastic pool with minimum parameters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateElasticPoolWithMinimumParameters(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getElasticPools().createOrUpdate("sqlcrudtest-2369",
            "sqlcrudtest-8069", "sqlcrudtest-8102", new ElasticPoolInner().withLocation("Japan East"), Context.NONE);
    }
}
