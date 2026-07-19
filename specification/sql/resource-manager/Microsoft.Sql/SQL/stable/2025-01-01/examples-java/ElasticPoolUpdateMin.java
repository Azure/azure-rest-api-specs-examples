
import com.azure.resourcemanager.sql.models.ElasticPoolUpdate;

/**
 * Samples for ElasticPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ElasticPoolUpdateMin.json
     */
    /**
     * Sample code: Update an elastic pool with minimum parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        updateAnElasticPoolWithMinimumParameters(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPools().update("sqlcrudtest-2369", "sqlcrudtest-8069", "sqlcrudtest-8102",
            new ElasticPoolUpdate(), com.azure.core.util.Context.NONE);
    }
}
