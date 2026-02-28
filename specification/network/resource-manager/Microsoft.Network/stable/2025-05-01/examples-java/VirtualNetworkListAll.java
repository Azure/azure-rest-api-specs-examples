
/**
 * Samples for VirtualNetworks List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkListAll.json
     */
    /**
     * Sample code: List all virtual networks.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllVirtualNetworks(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworks().list(com.azure.core.util.Context.NONE);
    }
}
