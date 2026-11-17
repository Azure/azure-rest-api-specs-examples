
/**
 * Samples for Clusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Clusters/ClustersListByResourceGroup.json
     */
    /**
     * Sample code: ClustersListByResourceGroup.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void clustersListByResourceGroup(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getClusters().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
