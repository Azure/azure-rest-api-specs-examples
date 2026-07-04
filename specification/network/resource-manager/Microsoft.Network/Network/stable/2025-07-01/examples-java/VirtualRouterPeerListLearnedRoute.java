
/**
 * Samples for VirtualHubBgpConnections ListLearnedRoutes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualRouterPeerListLearnedRoute.json
     */
    /**
     * Sample code: VirtualRouterPeerListLearnedRoutes.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualRouterPeerListLearnedRoutes(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubBgpConnections().listLearnedRoutes("rg1", "virtualRouter1", "peer1",
            com.azure.core.util.Context.NONE);
    }
}
