import com.azure.core.util.Context;

/** Samples for NetworkWatchers ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkWatcherList.json
     */
    /**
     * Sample code: List network watchers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listNetworkWatchers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkWatchers().listByResourceGroup("rg1", Context.NONE);
    }
}
