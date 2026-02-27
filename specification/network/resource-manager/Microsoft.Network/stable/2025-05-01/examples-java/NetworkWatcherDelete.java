
/**
 * Samples for NetworkWatchers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkWatcherDelete.json
     */
    /**
     * Sample code: Delete network watcher.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteNetworkWatcher(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkWatchers().delete("rg1", "nw1",
            com.azure.core.util.Context.NONE);
    }
}
