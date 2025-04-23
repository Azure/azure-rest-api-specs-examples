
/**
 * Samples for CloudServicesNetworks List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-10-01-preview/examples/
     * CloudServicesNetworks_ListBySubscription.json
     */
    /**
     * Sample code: List cloud services networks for subscription.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listCloudServicesNetworksForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.cloudServicesNetworks().list(com.azure.core.util.Context.NONE);
    }
}
