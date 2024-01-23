
/**
 * Samples for Clusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2020-03-01-preview/examples/
     * Cluster_Get.json
     */
    /**
     * Sample code: Get a cluster.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getACluster(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.clusters().getByResourceGroupWithResponse("sjrg", "testcluster", com.azure.core.util.Context.NONE);
    }
}
