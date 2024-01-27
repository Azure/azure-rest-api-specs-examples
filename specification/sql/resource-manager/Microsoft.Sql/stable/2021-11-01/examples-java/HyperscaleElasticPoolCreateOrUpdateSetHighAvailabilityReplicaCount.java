
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ElasticPoolInner;
import com.azure.resourcemanager.sql.models.Sku;

/** Samples for ElasticPools CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * HyperscaleElasticPoolCreateOrUpdateSetHighAvailabilityReplicaCount.json
     */
    /**
     * Sample code: Create or update Hyperscale elastic pool with high availability replica count parameter.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateHyperscaleElasticPoolWithHighAvailabilityReplicaCountParameter(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getElasticPools().createOrUpdate("sqlcrudtest-2369",
            "sqlcrudtest-8069", "sqlcrudtest-8102", new ElasticPoolInner().withLocation("Japan East")
                .withSku(new Sku().withName("HS_Gen5_4")).withHighAvailabilityReplicaCount(2),
            Context.NONE);
    }
}
