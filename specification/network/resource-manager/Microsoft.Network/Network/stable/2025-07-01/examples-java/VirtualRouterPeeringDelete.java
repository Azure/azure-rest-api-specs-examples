
/**
 * Samples for VirtualRouterPeerings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualRouterPeeringDelete.json
     */
    /**
     * Sample code: Delete VirtualRouterPeering.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteVirtualRouterPeering(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualRouterPeerings().delete("rg1", "virtualRouter", "peering1",
            com.azure.core.util.Context.NONE);
    }
}
