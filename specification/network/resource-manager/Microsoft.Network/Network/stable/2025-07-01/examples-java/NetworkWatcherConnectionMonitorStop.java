
/**
 * Samples for ConnectionMonitors Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkWatcherConnectionMonitorStop.json
     */
    /**
     * Sample code: Stop connection monitor.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void stopConnectionMonitor(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getConnectionMonitors().stop("rg1", "nw1", "cm1", com.azure.core.util.Context.NONE);
    }
}
