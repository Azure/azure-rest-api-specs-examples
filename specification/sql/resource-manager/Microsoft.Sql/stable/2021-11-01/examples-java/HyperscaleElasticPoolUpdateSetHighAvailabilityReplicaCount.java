
import com.azure.resourcemanager.sql.models.ElasticPoolUpdate;

/**
 * Samples for ElasticPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * HyperscaleElasticPoolUpdateSetHighAvailabilityReplicaCount.json
     */
    /**
     * Sample code: Update high availability replica count of a Hyperscale elastic pool.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateHighAvailabilityReplicaCountOfAHyperscaleElasticPool(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getElasticPools().update("sqlcrudtest-2369", "sqlcrudtest-8069",
            "sqlcrudtest-8102", new ElasticPoolUpdate().withHighAvailabilityReplicaCount(2),
            com.azure.core.util.Context.NONE);
    }
}
