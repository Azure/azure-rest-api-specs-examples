/** Samples for CloudServicesNetworks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2023-07-01/examples/CloudServicesNetworks_ListBySubscription.json
     */
    /**
     * Sample code: List cloud services networks for subscription.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listCloudServicesNetworksForSubscription(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.cloudServicesNetworks().list(com.azure.core.util.Context.NONE);
    }
}
