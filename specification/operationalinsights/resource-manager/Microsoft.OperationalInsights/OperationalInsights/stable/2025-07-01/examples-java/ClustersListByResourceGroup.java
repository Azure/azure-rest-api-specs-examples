
/**
 * Samples for Clusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ClustersListByResourceGroup.json
     */
    /**
     * Sample code: ClustersGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void clustersGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.clusters().listByResourceGroup("oiautorest6685", com.azure.core.util.Context.NONE);
    }
}
