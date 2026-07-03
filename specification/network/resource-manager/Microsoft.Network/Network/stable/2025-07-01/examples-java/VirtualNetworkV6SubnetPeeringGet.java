
/**
 * Samples for VirtualNetworkPeerings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkV6SubnetPeeringGet.json
     */
    /**
     * Sample code: Get V6 subnet peering.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getV6SubnetPeering(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkPeerings().getWithResponse("peerTest", "vnet1", "peer",
            com.azure.core.util.Context.NONE);
    }
}
