
/**
 * Samples for ConnectionMonitors List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NetworkWatcherConnectionMonitorList.json
     */
    /**
     * Sample code: List connection monitors.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listConnectionMonitors(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getConnectionMonitors().list("rg1", "nw1", com.azure.core.util.Context.NONE);
    }
}
