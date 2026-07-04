
/**
 * Samples for VirtualRouterPeerings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualRouterPeeringGet.json
     */
    /**
     * Sample code: Get Virtual Router Peering.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getVirtualRouterPeering(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualRouterPeerings().getWithResponse("rg1", "virtualRouter", "peering1",
            com.azure.core.util.Context.NONE);
    }
}
