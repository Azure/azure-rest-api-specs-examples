
/**
 * Samples for MetricsConfigurations ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ClusterMetricsConfigurations_ListByCluster.json
     */
    /**
     * Sample code: List metrics configurations of the cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listMetricsConfigurationsOfTheCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.metricsConfigurations().listByCluster("resourceGroupName", "clusterName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
