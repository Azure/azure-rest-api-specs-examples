
/**
 * Samples for NetworkWatchers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkWatcherDelete.json
     */
    /**
     * Sample code: Delete network watcher.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteNetworkWatcher(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkWatchers().delete("rg1", "nw1", com.azure.core.util.Context.NONE);
    }
}
