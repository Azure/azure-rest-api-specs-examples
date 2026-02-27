
/**
 * Samples for VirtualNetworks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkList.json
     */
    /**
     * Sample code: List virtual networks in resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVirtualNetworksInResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworks().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
