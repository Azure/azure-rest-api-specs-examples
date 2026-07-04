
/**
 * Samples for ConnectionMonitors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkWatcherConnectionMonitorDelete.json
     */
    /**
     * Sample code: Delete connection monitor.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteConnectionMonitor(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getConnectionMonitors().delete("rg1", "nw1", "cm1", com.azure.core.util.Context.NONE);
    }
}
