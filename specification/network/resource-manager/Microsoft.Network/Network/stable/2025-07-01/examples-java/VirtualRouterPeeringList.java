
/**
 * Samples for VirtualRouterPeerings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualRouterPeeringList.json
     */
    /**
     * Sample code: List all Virtual Router Peerings for a given Virtual Router.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listAllVirtualRouterPeeringsForAGivenVirtualRouter(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualRouterPeerings().list("rg1", "virtualRouter",
            com.azure.core.util.Context.NONE);
    }
}
