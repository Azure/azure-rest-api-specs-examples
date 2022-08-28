import com.azure.core.util.Context;

/** Samples for VirtualHubBgpConnections ListAdvertisedRoutes. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualRouterPeerListAdvertisedRoute.json
     */
    /**
     * Sample code: VirtualRouterPeerListAdvertisedRoutes.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualRouterPeerListAdvertisedRoutes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualHubBgpConnections()
            .listAdvertisedRoutes("rg1", "virtualRouter1", "peer1", Context.NONE);
    }
}
