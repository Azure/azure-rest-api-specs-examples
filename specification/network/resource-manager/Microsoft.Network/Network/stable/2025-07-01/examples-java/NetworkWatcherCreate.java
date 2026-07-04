
import com.azure.resourcemanager.network.fluent.models.NetworkWatcherInner;

/**
 * Samples for NetworkWatchers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkWatcherCreate.json
     */
    /**
     * Sample code: Create network watcher.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createNetworkWatcher(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkWatchers().createOrUpdateWithResponse("rg1", "nw1",
            new NetworkWatcherInner().withLocation("eastus"), com.azure.core.util.Context.NONE);
    }
}
