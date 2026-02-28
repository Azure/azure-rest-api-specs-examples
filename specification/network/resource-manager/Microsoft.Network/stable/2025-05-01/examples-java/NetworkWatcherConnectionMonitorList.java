
/**
 * Samples for ConnectionMonitors List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkWatcherConnectionMonitorList.json
     */
    /**
     * Sample code: List connection monitors.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listConnectionMonitors(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getConnectionMonitors().list("rg1", "nw1",
            com.azure.core.util.Context.NONE);
    }
}
