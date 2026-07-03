
/**
 * Samples for NetworkVirtualAppliances List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkVirtualApplianceListBySubscription.json
     */
    /**
     * Sample code: List all Network Virtual Appliances for a given subscription.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listAllNetworkVirtualAppliancesForAGivenSubscription(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualAppliances().list(com.azure.core.util.Context.NONE);
    }
}
