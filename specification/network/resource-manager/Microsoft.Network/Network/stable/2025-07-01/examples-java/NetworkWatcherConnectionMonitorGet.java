
/**
 * Samples for ConnectionMonitors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkWatcherConnectionMonitorGet.json
     */
    /**
     * Sample code: Get connection monitor.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getConnectionMonitor(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getConnectionMonitors().getWithResponse("rg1", "nw1", "cm1",
            com.azure.core.util.Context.NONE);
    }
}
