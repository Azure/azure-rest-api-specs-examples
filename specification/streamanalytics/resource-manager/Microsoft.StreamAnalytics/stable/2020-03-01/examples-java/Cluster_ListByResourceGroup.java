import com.azure.core.util.Context;

/** Samples for Clusters ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Cluster_ListByResourceGroup.json
     */
    /**
     * Sample code: List clusters in resource group.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void listClustersInResourceGroup(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.clusters().listByResourceGroup("sjrg", Context.NONE);
    }
}
