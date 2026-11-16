
/**
 * Samples for Clusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Clusters/ClusterDelete.json
     */
    /**
     * Sample code: ClusterDelete.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void clusterDelete(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getClusters().delete("myResourceGroup", "testCluster",
            com.azure.core.util.Context.NONE);
    }
}
