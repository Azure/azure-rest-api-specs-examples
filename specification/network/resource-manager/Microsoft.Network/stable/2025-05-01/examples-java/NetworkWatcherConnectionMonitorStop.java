
/**
 * Samples for ConnectionMonitors Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkWatcherConnectionMonitorStop.json
     */
    /**
     * Sample code: Stop connection monitor.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void stopConnectionMonitor(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getConnectionMonitors().stop("rg1", "nw1", "cm1",
            com.azure.core.util.Context.NONE);
    }
}
