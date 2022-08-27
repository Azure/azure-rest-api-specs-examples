import com.azure.core.util.Context;

/** Samples for SubscriptionNetworkManagerConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerConnectionSubscriptionGet.json
     */
    /**
     * Sample code: Get Subscription Network Manager Connection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSubscriptionNetworkManagerConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getSubscriptionNetworkManagerConnections()
            .getWithResponse("TestNMConnection", Context.NONE);
    }
}
