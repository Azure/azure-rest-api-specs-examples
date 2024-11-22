
/**
 * Samples for MetricsConfigurations ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-06-01-preview/examples/
     * ClusterMetricsConfigurations_ListByCluster.json
     */
    /**
     * Sample code: List metrics configurations of the cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listMetricsConfigurationsOfTheCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.metricsConfigurations().listByCluster("resourceGroupName", "clusterName",
            com.azure.core.util.Context.NONE);
    }
}
