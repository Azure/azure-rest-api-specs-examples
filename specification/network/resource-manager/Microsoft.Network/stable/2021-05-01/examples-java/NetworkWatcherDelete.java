import com.azure.core.util.Context;

/** Samples for NetworkWatchers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherDelete.json
     */
    /**
     * Sample code: Delete network watcher.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteNetworkWatcher(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkWatchers().delete("rg1", "nw1", Context.NONE);
    }
}
