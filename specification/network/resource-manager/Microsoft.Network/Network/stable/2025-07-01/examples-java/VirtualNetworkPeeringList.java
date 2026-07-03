
/**
 * Samples for VirtualNetworkPeerings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkPeeringList.json
     */
    /**
     * Sample code: List peerings.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listPeerings(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkPeerings().list("peerTest", "vnet1", com.azure.core.util.Context.NONE);
    }
}
