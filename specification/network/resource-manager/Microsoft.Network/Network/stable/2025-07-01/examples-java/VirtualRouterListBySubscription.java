
/**
 * Samples for VirtualRouters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualRouterListBySubscription.json
     */
    /**
     * Sample code: List all Virtual Routers for a given subscription.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listAllVirtualRoutersForAGivenSubscription(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualRouters().list(com.azure.core.util.Context.NONE);
    }
}
