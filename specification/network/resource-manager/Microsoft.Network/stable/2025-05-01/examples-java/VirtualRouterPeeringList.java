
/**
 * Samples for VirtualRouterPeerings List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualRouterPeeringList.json
     */
    /**
     * Sample code: List all Virtual Router Peerings for a given Virtual Router.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAllVirtualRouterPeeringsForAGivenVirtualRouter(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualRouterPeerings().list("rg1", "virtualRouter",
            com.azure.core.util.Context.NONE);
    }
}
