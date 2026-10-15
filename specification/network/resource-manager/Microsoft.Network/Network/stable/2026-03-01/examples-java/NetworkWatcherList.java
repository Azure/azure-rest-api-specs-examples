
/**
 * Samples for NetworkWatchers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/NetworkWatcherList.json
     */
    /**
     * Sample code: List network watchers.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listNetworkWatchers(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkWatchers().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
