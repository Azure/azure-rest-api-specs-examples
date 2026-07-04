
/**
 * Samples for VirtualNetworkPeerings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkPeeringListWithRemoteVirtualNetworkEncryption.json
     */
    /**
     * Sample code: List peerings with remote virtual network encryption.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listPeeringsWithRemoteVirtualNetworkEncryption(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkPeerings().list("peerTest", "vnet1", com.azure.core.util.Context.NONE);
    }
}
