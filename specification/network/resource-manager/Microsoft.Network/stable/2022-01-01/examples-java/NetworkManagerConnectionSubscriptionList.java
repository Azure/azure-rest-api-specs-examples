import com.azure.core.util.Context;

/** Samples for SubscriptionNetworkManagerConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerConnectionSubscriptionList.json
     */
    /**
     * Sample code: List Subscription Network Manager Connection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSubscriptionNetworkManagerConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getSubscriptionNetworkManagerConnections()
            .list(null, null, Context.NONE);
    }
}
