
/**
 * Samples for VirtualNetworkTaps List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkTapListAll.json
     */
    /**
     * Sample code: List all virtual network taps.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllVirtualNetworkTaps(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkTaps().list(com.azure.core.util.Context.NONE);
    }
}
