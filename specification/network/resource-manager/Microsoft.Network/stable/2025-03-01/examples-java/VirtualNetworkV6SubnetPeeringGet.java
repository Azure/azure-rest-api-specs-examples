
/**
 * Samples for VirtualNetworkPeerings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VirtualNetworkV6SubnetPeeringGet.json
     */
    /**
     * Sample code: Get V6 subnet peering.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getV6SubnetPeering(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkPeerings().getWithResponse("peerTest", "vnet1",
            "peer", com.azure.core.util.Context.NONE);
    }
}
