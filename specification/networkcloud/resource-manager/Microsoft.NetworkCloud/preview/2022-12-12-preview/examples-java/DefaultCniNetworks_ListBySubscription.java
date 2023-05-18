/** Samples for DefaultCniNetworks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/DefaultCniNetworks_ListBySubscription.json
     */
    /**
     * Sample code: List default CNI networks for subscription.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listDefaultCNINetworksForSubscription(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.defaultCniNetworks().list(com.azure.core.util.Context.NONE);
    }
}
