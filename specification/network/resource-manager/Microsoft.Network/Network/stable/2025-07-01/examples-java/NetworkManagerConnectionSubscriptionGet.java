
/**
 * Samples for SubscriptionNetworkManagerConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerConnectionSubscriptionGet.json
     */
    /**
     * Sample code: Get Subscription Network Manager Connection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getSubscriptionNetworkManagerConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSubscriptionNetworkManagerConnections().getWithResponse("TestNMConnection",
            com.azure.core.util.Context.NONE);
    }
}
