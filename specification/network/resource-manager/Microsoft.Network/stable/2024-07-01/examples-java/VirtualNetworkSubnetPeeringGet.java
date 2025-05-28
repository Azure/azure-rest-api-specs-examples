
/**
 * Samples for VirtualNetworkPeerings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * VirtualNetworkSubnetPeeringGet.json
     */
    /**
     * Sample code: Get subnet peering.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSubnetPeering(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkPeerings().getWithResponse("peerTest", "vnet1",
            "peer", com.azure.core.util.Context.NONE);
    }
}
