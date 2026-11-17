
/**
 * Samples for Clusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Clusters/ClustersListBySubscription.json
     */
    /**
     * Sample code: ClustersListBySubscription.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void clustersListBySubscription(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getClusters().list(com.azure.core.util.Context.NONE);
    }
}
