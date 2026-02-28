
/**
 * Samples for ConnectionMonitors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkWatcherConnectionMonitorGet.json
     */
    /**
     * Sample code: Get connection monitor.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getConnectionMonitor(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getConnectionMonitors().getWithResponse("rg1", "nw1", "cm1",
            com.azure.core.util.Context.NONE);
    }
}
