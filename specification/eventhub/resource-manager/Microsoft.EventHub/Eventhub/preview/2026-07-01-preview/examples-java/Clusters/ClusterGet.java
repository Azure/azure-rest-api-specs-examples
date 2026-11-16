
/**
 * Samples for Clusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Clusters/ClusterGet.json
     */
    /**
     * Sample code: ClusterGet.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void clusterGet(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getClusters().getByResourceGroupWithResponse("myResourceGroup", "testCluster",
            com.azure.core.util.Context.NONE);
    }
}
