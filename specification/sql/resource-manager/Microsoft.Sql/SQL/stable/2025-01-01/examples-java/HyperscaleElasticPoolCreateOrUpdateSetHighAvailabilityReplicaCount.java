
import com.azure.resourcemanager.sql.fluent.models.ElasticPoolInner;
import com.azure.resourcemanager.sql.models.Sku;

/**
 * Samples for ElasticPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/HyperscaleElasticPoolCreateOrUpdateSetHighAvailabilityReplicaCount.json
     */
    /**
     * Sample code: Create or update Hyperscale elastic pool with high availability replica count parameter.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createOrUpdateHyperscaleElasticPoolWithHighAvailabilityReplicaCountParameter(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPools().createOrUpdate("sqlcrudtest-2369", "sqlcrudtest-8069",
            "sqlcrudtest-8102", new ElasticPoolInner().withLocation("Japan East")
                .withSku(new Sku().withName("HS_Gen5_4")).withHighAvailabilityReplicaCount(2),
            com.azure.core.util.Context.NONE);
    }
}
