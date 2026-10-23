
/**
 * Samples for Clusters Failover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Clusters_Failover.json
     */
    /**
     * Sample code: Trigger failover within a HorizonDB cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        triggerFailoverWithinAHorizonDBCluster(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.clusters().failover("example-resource-group", "example-cluster", com.azure.core.util.Context.NONE);
    }
}
