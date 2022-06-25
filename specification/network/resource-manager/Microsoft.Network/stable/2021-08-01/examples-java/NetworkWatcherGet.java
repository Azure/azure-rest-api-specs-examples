import com.azure.core.util.Context;

/** Samples for NetworkWatchers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkWatcherGet.json
     */
    /**
     * Sample code: Get network watcher.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getNetworkWatcher(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkWatchers()
            .getByResourceGroupWithResponse("rg1", "nw1", Context.NONE);
    }
}
