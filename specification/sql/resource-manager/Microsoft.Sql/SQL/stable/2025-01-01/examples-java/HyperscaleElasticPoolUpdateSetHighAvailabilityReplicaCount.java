
import com.azure.resourcemanager.sql.models.ElasticPoolUpdate;

/**
 * Samples for ElasticPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/HyperscaleElasticPoolUpdateSetHighAvailabilityReplicaCount.json
     */
    /**
     * Sample code: Update high availability replica count of a Hyperscale elastic pool.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateHighAvailabilityReplicaCountOfAHyperscaleElasticPool(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPools().update("sqlcrudtest-2369", "sqlcrudtest-8069", "sqlcrudtest-8102",
            new ElasticPoolUpdate().withHighAvailabilityReplicaCount(2), com.azure.core.util.Context.NONE);
    }
}
