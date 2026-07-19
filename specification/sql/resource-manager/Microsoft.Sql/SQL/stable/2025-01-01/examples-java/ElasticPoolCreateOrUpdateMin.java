
import com.azure.resourcemanager.sql.fluent.models.ElasticPoolInner;

/**
 * Samples for ElasticPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ElasticPoolCreateOrUpdateMin.json
     */
    /**
     * Sample code: Create or update elastic pool with minimum parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createOrUpdateElasticPoolWithMinimumParameters(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPools().createOrUpdate("sqlcrudtest-2369", "sqlcrudtest-8069",
            "sqlcrudtest-8102", new ElasticPoolInner().withLocation("Japan East"), com.azure.core.util.Context.NONE);
    }
}
