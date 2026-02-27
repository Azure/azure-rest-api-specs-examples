
/**
 * Samples for VirtualNetworkTaps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkTapDelete.json
     */
    /**
     * Sample code: Delete Virtual Network Tap resource.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteVirtualNetworkTapResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkTaps().delete("rg1", "test-vtap",
            com.azure.core.util.Context.NONE);
    }
}
