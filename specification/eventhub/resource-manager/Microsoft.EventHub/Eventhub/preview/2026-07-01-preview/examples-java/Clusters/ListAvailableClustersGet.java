
/**
 * Samples for Clusters ListAvailableClusterRegion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Clusters/ListAvailableClustersGet.json
     */
    /**
     * Sample code: ListAvailableClusters.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void listAvailableClusters(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getClusters().listAvailableClusterRegionWithResponse(com.azure.core.util.Context.NONE);
    }
}
