
/**
 * Samples for VirtualRouterPeerings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualRouterPeeringDelete.
     * json
     */
    /**
     * Sample code: Delete VirtualRouterPeering.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteVirtualRouterPeering(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualRouterPeerings().delete("rg1", "virtualRouter", "peering1",
            com.azure.core.util.Context.NONE);
    }
}
