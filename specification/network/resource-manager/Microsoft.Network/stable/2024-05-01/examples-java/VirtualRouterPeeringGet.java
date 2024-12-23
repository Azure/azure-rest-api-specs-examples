
/**
 * Samples for VirtualRouterPeerings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/VirtualRouterPeeringGet.json
     */
    /**
     * Sample code: Get Virtual Router Peering.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualRouterPeering(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualRouterPeerings().getWithResponse("rg1", "virtualRouter",
            "peering1", com.azure.core.util.Context.NONE);
    }
}
