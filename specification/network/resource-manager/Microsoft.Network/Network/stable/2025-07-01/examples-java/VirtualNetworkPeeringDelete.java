
/**
 * Samples for VirtualNetworkPeerings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkPeeringDelete.json
     */
    /**
     * Sample code: Delete peering.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deletePeering(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkPeerings().delete("peerTest", "vnet1", "peer",
            com.azure.core.util.Context.NONE);
    }
}
