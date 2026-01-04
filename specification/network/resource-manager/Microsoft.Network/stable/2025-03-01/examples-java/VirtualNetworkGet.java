
/**
 * Samples for VirtualNetworks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkGet.json
     */
    /**
     * Sample code: Get virtual network.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualNetwork(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworks().getByResourceGroupWithResponse("rg1",
            "test-vnet", null, com.azure.core.util.Context.NONE);
    }
}
