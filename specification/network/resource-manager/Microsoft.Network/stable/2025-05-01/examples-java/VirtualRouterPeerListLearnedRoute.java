
/**
 * Samples for VirtualHubBgpConnections ListLearnedRoutes.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualRouterPeerListLearnedRoute.json
     */
    /**
     * Sample code: VirtualRouterPeerListLearnedRoutes.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualRouterPeerListLearnedRoutes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualHubBgpConnections().listLearnedRoutes("rg1",
            "virtualRouter1", "peer1", com.azure.core.util.Context.NONE);
    }
}
