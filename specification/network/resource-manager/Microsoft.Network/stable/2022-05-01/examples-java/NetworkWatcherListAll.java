import com.azure.core.util.Context;

/** Samples for NetworkWatchers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkWatcherListAll.json
     */
    /**
     * Sample code: List all network watchers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllNetworkWatchers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkWatchers().list(Context.NONE);
    }
}
