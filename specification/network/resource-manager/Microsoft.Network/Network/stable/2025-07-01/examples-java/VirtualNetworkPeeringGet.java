
/**
 * Samples for VirtualNetworkPeerings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkPeeringGet.json
     */
    /**
     * Sample code: Get peering.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getPeering(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkPeerings().getWithResponse("peerTest", "vnet1", "peer",
            com.azure.core.util.Context.NONE);
    }
}
