
/**
 * Samples for TrunkedNetworks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/TrunkedNetworks_ListBySubscription.json
     */
    /**
     * Sample code: List trunked networks for subscription.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listTrunkedNetworksForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.trunkedNetworks().list(null, null, com.azure.core.util.Context.NONE);
    }
}
