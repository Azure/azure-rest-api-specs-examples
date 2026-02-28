
/**
 * Samples for SubscriptionNetworkManagerConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkManagerConnectionSubscriptionDelete.json
     */
    /**
     * Sample code: Delete Subscription Network Manager Connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deleteSubscriptionNetworkManagerConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSubscriptionNetworkManagerConnections()
            .deleteWithResponse("TestNMConnection", com.azure.core.util.Context.NONE);
    }
}
