
/**
 * Samples for MetricsConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * ClusterMetricsConfigurations_Get.json
     */
    /**
     * Sample code: Get metrics configuration of cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        getMetricsConfigurationOfCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.metricsConfigurations().getWithResponse("resourceGroupName", "clusterName", "default",
            com.azure.core.util.Context.NONE);
    }
}
