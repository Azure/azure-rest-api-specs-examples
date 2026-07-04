
/**
 * Samples for VirtualHubBgpConnections ListAdvertisedRoutes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualRouterPeerListAdvertisedRoute.json
     */
    /**
     * Sample code: VirtualRouterPeerListAdvertisedRoutes.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualRouterPeerListAdvertisedRoutes(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubBgpConnections().listAdvertisedRoutes("rg1", "virtualRouter1", "peer1",
            com.azure.core.util.Context.NONE);
    }
}
