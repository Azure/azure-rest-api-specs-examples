
import com.azure.resourcemanager.streamanalytics.models.Cluster;
import com.azure.resourcemanager.streamanalytics.models.ClusterSku;

/**
 * Samples for Clusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/
     * Cluster_Update.json
     */
    /**
     * Sample code: Update a cluster.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateACluster(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Cluster resource = manager.clusters()
            .getByResourceGroupWithResponse("sjrg", "testcluster", com.azure.core.util.Context.NONE).getValue();
        resource.update().withSku(new ClusterSku().withCapacity(96)).apply();
    }
}
