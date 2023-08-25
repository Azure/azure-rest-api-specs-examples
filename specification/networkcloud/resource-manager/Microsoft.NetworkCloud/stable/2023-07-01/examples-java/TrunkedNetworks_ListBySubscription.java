/** Samples for TrunkedNetworks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2023-07-01/examples/TrunkedNetworks_ListBySubscription.json
     */
    /**
     * Sample code: List trunked networks for subscription.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listTrunkedNetworksForSubscription(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.trunkedNetworks().list(com.azure.core.util.Context.NONE);
    }
}
