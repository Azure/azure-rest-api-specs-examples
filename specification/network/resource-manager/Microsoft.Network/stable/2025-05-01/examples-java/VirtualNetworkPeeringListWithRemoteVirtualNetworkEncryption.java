
/**
 * Samples for VirtualNetworkPeerings List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualNetworkPeeringListWithRemoteVirtualNetworkEncryption.json
     */
    /**
     * Sample code: List peerings with remote virtual network encryption.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listPeeringsWithRemoteVirtualNetworkEncryption(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkPeerings().list("peerTest", "vnet1",
            com.azure.core.util.Context.NONE);
    }
}
