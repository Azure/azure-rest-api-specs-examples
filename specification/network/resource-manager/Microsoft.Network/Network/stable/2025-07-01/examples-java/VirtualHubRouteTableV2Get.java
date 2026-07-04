
/**
 * Samples for VirtualHubRouteTableV2S Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualHubRouteTableV2Get.json
     */
    /**
     * Sample code: VirtualHubVirtualHubRouteTableV2Get.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualHubVirtualHubRouteTableV2Get(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubRouteTableV2S().getWithResponse("rg1", "virtualHub1",
            "virtualHubRouteTable1a", com.azure.core.util.Context.NONE);
    }
}
