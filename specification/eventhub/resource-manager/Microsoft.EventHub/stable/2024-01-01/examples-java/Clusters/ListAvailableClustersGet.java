
/**
 * Samples for Clusters ListAvailableClusterRegion.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/Clusters/
     * ListAvailableClustersGet.json
     */
    /**
     * Sample code: ListAvailableClusters.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAvailableClusters(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getClusters()
            .listAvailableClusterRegionWithResponse(com.azure.core.util.Context.NONE);
    }
}
