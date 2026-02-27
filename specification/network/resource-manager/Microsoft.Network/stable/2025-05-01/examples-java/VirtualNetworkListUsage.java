
/**
 * Samples for VirtualNetworks ListUsage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkListUsage.json
     */
    /**
     * Sample code: VnetGetUsage.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vnetGetUsage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworks().listUsage("rg1", "vnetName",
            com.azure.core.util.Context.NONE);
    }
}
