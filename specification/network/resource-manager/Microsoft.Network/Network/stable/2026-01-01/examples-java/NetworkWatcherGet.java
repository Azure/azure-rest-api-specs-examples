
/**
 * Samples for NetworkWatchers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NetworkWatcherGet.json
     */
    /**
     * Sample code: Get network watcher.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getNetworkWatcher(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkWatchers().getByResourceGroupWithResponse("rg1", "nw1",
            com.azure.core.util.Context.NONE);
    }
}
