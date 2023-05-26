/** Samples for VirtualNetworkPeerings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/VirtualNetworkPeeringGetWithRemoteVirtualNetworkEncryption.json
     */
    /**
     * Sample code: Get peering with remote virtual network encryption.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPeeringWithRemoteVirtualNetworkEncryption(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkPeerings()
            .getWithResponse("peerTest", "vnet1", "peer", com.azure.core.util.Context.NONE);
    }
}
