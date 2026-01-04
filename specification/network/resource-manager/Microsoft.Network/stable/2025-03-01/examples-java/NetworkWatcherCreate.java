
import com.azure.resourcemanager.network.fluent.models.NetworkWatcherInner;

/**
 * Samples for NetworkWatchers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkWatcherCreate.json
     */
    /**
     * Sample code: Create network watcher.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createNetworkWatcher(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkWatchers().createOrUpdateWithResponse("rg1", "nw1",
            new NetworkWatcherInner().withLocation("eastus"), com.azure.core.util.Context.NONE);
    }
}
