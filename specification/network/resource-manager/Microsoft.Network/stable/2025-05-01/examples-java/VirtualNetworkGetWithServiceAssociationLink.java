
/**
 * Samples for VirtualNetworks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualNetworkGetWithServiceAssociationLink.json
     */
    /**
     * Sample code: Get virtual network with service association links.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getVirtualNetworkWithServiceAssociationLinks(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworks().getByResourceGroupWithResponse("rg1",
            "test-vnet", null, com.azure.core.util.Context.NONE);
    }
}
