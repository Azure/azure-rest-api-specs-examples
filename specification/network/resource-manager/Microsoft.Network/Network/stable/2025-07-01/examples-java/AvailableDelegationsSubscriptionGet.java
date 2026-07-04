
/**
 * Samples for AvailableDelegations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AvailableDelegationsSubscriptionGet.json
     */
    /**
     * Sample code: Get available delegations.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAvailableDelegations(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAvailableDelegations().list("westcentralus", com.azure.core.util.Context.NONE);
    }
}
