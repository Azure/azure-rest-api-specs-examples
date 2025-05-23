
import com.azure.resourcemanager.kusto.models.ClusterMigrateRequest;

/**
 * Samples for Clusters Migrate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoClusterMigrate.json
     */
    /**
     * Sample code: KustoClusterMigrate.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterMigrate(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().migrate("kustorptest", "kustoCluster1", new ClusterMigrateRequest().withClusterResourceId(
            "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/clusters/kustoCluster2"),
            com.azure.core.util.Context.NONE);
    }
}
