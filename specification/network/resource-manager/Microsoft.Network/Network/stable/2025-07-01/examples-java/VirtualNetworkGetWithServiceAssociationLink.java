
/**
 * Samples for VirtualNetworks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGetWithServiceAssociationLink.json
     */
    /**
     * Sample code: Get virtual network with service association links.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getVirtualNetworkWithServiceAssociationLinks(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworks().getByResourceGroupWithResponse("rg1", "test-vnet", null,
            com.azure.core.util.Context.NONE);
    }
}
