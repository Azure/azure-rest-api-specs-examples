import com.azure.core.util.Context;

/** Samples for PrivateEndpoints ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/PrivateEndpoint_ListByCluster.json
     */
    /**
     * Sample code: Get the private endpoints in a cluster.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getThePrivateEndpointsInACluster(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.privateEndpoints().listByCluster("sjrg", "testcluster", Context.NONE);
    }
}
