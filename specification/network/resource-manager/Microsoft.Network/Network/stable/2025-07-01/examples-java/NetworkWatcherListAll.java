
/**
 * Samples for NetworkWatchers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkWatcherListAll.json
     */
    /**
     * Sample code: List all network watchers.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllNetworkWatchers(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkWatchers().list(com.azure.core.util.Context.NONE);
    }
}
