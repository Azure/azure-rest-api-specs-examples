
/**
 * Samples for SubscriptionNetworkManagerConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerConnectionSubscriptionDelete.json
     */
    /**
     * Sample code: Delete Subscription Network Manager Connection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        deleteSubscriptionNetworkManagerConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSubscriptionNetworkManagerConnections().deleteWithResponse("TestNMConnection",
            com.azure.core.util.Context.NONE);
    }
}
