import com.azure.core.util.Context;

/** Samples for VirtualRouters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualRouterListBySubscription.json
     */
    /**
     * Sample code: List all Virtual Routers for a given subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllVirtualRoutersForAGivenSubscription(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualRouters().list(Context.NONE);
    }
}
