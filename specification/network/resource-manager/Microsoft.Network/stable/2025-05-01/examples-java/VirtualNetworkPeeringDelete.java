
/**
 * Samples for VirtualNetworkPeerings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkPeeringDelete.
     * json
     */
    /**
     * Sample code: Delete peering.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletePeering(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkPeerings().delete("peerTest", "vnet1", "peer",
            com.azure.core.util.Context.NONE);
    }
}
