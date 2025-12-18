
/**
 * Samples for MetricsConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * ClusterMetricsConfigurations_Delete.json
     */
    /**
     * Sample code: Delete metrics configuration of cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        deleteMetricsConfigurationOfCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.metricsConfigurations().delete("resourceGroupName", "clusterName", "default", null, null,
            com.azure.core.util.Context.NONE);
    }
}
