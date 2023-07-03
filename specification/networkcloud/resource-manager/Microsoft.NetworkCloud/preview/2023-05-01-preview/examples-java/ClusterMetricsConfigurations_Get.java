/** Samples for MetricsConfigurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/ClusterMetricsConfigurations_Get.json
     */
    /**
     * Sample code: Get metrics configuration of cluster.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getMetricsConfigurationOfCluster(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .metricsConfigurations()
            .getWithResponse("resourceGroupName", "clusterName", "default", com.azure.core.util.Context.NONE);
    }
}
